package artifactory

import (
	"archive/tar"
	"compress/gzip"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/concourse/go-archive/tarfs"
	"github.com/fatih/color"
	"github.com/google/go-containerregistry/pkg/authn"
	"github.com/google/go-containerregistry/pkg/name"
	v1 "github.com/google/go-containerregistry/pkg/v1"
	"github.com/google/go-containerregistry/pkg/v1/empty"
	"github.com/google/go-containerregistry/pkg/v1/remote"
	"github.com/google/go-containerregistry/pkg/v1/tarball"
	"github.com/vbauerster/mpb"
	"github.com/vbauerster/mpb/decor"
)

// NOTICE
// Significant portions of this code have been adapted from Alex Suraci's work
// in this project: https://github.com/concourse/registry-image-resource

// PullImage downloads an image in etiher OCI or RootFS format
func (c *Client) PullImage(dest, format, repository, identifier, digest string) error {
	repo, err := name.NewRepository(repository, name.WeakValidation)
	if err != nil {
		return err
	}

	var image v1.Image

	dig := repo.Digest(digest)
	image, err = get(c.BasicCredentials(), dig)
	if err != nil {
		return err
	}

	tag := repo.Tag(identifier)
	err = save(dest, tag, image, format)
	if err != nil {
		return err
	}

	return nil
}

func get(creds BasicCredentials, digest name.Digest) (v1.Image, error) {
	auth := &authn.Basic{
		Username: creds.Username,
		Password: creds.Password,
	}

	imageOpts := []remote.Option{}

	if auth.Username != "" && auth.Password != "" {
		imageOpts = append(imageOpts, remote.WithAuth(auth))
	}

	image, err := remote.Image(digest, imageOpts...)
	if err != nil {
		return nil, fmt.Errorf("locate remote image: %w", err)
	}
	if image == empty.Image {
		return nil, fmt.Errorf("download image")
	}

	return image, err
}

func save(dest string, tag name.Tag, image v1.Image, format string) error {
	switch format {
	case "oci":
		err := ociFormat(dest, tag, image)
		if err != nil {
			return fmt.Errorf("write oci image: %w", err)
		}
	default:
		err := rootfsFormat(dest, image)
		if err != nil {
			return fmt.Errorf("write rootfs: %w", err)
		}
	}

	err := ioutil.WriteFile(filepath.Join(dest, "tag"), []byte(tag.TagStr()), 0644)
	if err != nil {
		return fmt.Errorf("save image tag: %w", err)
	}

	err = saveDigest(dest, image)
	if err != nil {
		return fmt.Errorf("save image digest: %w", err)
	}

	return err
}

func saveDigest(dest string, image v1.Image) error {
	digest, err := image.Digest()
	if err != nil {
		return fmt.Errorf("get image digest: %w", err)
	}

	digestDest := filepath.Join(dest, "digest")
	return ioutil.WriteFile(digestDest, []byte(digest.String()), 0644)
}

func ociFormat(dest string, tag name.Tag, image v1.Image) error {
	err := tarball.WriteToFile(filepath.Join(dest, "image.tar"), tag, image)
	if err != nil {
		return fmt.Errorf("write OCI image: %s", err)
	}

	return nil
}

// ImageMetadata is used for rootfs format
type ImageMetadata struct {
	Env  []string `json:"env"`
	User string   `json:"user"`
}

func rootfsFormat(dest string, image v1.Image) error {
	err := unpackImage(filepath.Join(dest, "rootfs"), image)
	if err != nil {
		return fmt.Errorf("extract image: %w", err)
	}

	cfg, err := image.ConfigFile()
	if err != nil {
		return fmt.Errorf("inspect image config: %w", err)
	}

	meta, err := os.Create(filepath.Join(dest, "metadata.json"))
	if err != nil {
		return fmt.Errorf("create image metadata: %w", err)
	}

	err = json.NewEncoder(meta).Encode(ImageMetadata{Env: cfg.Config.Env, User: cfg.Config.User})
	if err != nil {
		return fmt.Errorf("write image metadata: %w", err)
	}

	err = meta.Close()
	if err != nil {
		return fmt.Errorf("close image metadata file: %w", err)
	}

	return nil
}

const whiteoutPrefix = ".wh."

func unpackImage(dest string, img v1.Image) error {
	layers, err := img.Layers()
	if err != nil {
		return err
	}

	chown := os.Getuid() == 0

	var out io.Writer
	out = os.Stderr

	progress := mpb.New(mpb.WithOutput(out))
	bars := make([]*mpb.Bar, len(layers))

	for i, layer := range layers {
		size, err := layer.Size()
		if err != nil {
			return err
		}

		digest, err := layer.Digest()
		if err != nil {
			return err
		}

		bars[i] = progress.AddBar(
			size,
			mpb.PrependDecorators(decor.Name(color.HiBlackString(digest.Hex[0:12]))),
			mpb.AppendDecorators(decor.CountersKibiByte("%.1f/%.1f")),
		)
	}

	// iterate over layers in reverse order; no need to write things files that
	// are modified by later layers anyway
	for i, layer := range layers {
		log.Printf("extracting layer %d of %d", i+1, len(layers))

		err := extractLayer(dest, layer, bars[i], chown)
		if err != nil {
			return err
		}
	}

	progress.Wait()

	return nil
}

func extractLayer(dest string, layer v1.Layer, bar *mpb.Bar, chown bool) error {
	r, err := layer.Compressed()
	if err != nil {
		return err
	}

	gr, err := gzip.NewReader(bar.ProxyReader(r))
	if err != nil {
		return err
	}

	tr := tar.NewReader(gr)

	for {
		hdr, err := tr.Next()
		if err == io.EOF {
			break
		}

		if err != nil {
			return err
		}

		path := filepath.Join(dest, filepath.Clean(hdr.Name))
		base := filepath.Base(path)
		dir := filepath.Dir(path)

		log.Println("unpacking")

		if strings.HasPrefix(base, whiteoutPrefix) {
			// layer has marked a file as deleted
			name := strings.TrimPrefix(base, whiteoutPrefix)
			removedPath := filepath.Join(dir, name)

			log.Printf("removing %s", removedPath)

			err := os.RemoveAll(removedPath)
			if err != nil {
				return nil
			}

			continue
		}

		if hdr.Typeflag == tar.TypeBlock || hdr.Typeflag == tar.TypeChar {
			// devices can't be created in a user namespace
			log.Printf("skipping device %s", hdr.Name)
			continue
		}

		if hdr.Typeflag == tar.TypeSymlink {
			log.Printf("symlinking to %s", hdr.Linkname)
		}

		if hdr.Typeflag == tar.TypeLink {
			log.Printf("hardlinking to %s", hdr.Linkname)
		}

		if fi, err := os.Lstat(path); err == nil {
			if fi.IsDir() && hdr.Name == "." {
				continue
			}

			if !(fi.IsDir() && hdr.Typeflag == tar.TypeDir) {
				log.Printf("removing existing path")
				if err := os.RemoveAll(path); err != nil {
					return err
				}
			}
		}

		if err := tarfs.ExtractEntry(hdr, dest, tr, chown); err != nil {
			log.Printf("extracting")
			return err
		}
	}

	err = gr.Close()
	if err != nil {
		return err
	}

	err = r.Close()
	if err != nil {
		return err
	}

	return nil
}

// PushImage uploads an image in OCI format
func (c *Client) PushImage(src, repository, image string, tags []string) error {
	if len(tags) < 1 {
		return errors.New("tags not specified")
	}

	ref, err := name.ParseReference(repository+tags[0], name.WeakValidation)
	if err != nil {
		return err
	}

	var extraRefs []name.Reference
	for _, tag := range tags[1:] {
		n := fmt.Sprintf("%s:%s", repository, tag)

		extraRef, err := name.ParseReference(n, name.WeakValidation)
		if err != nil {
			return err
		}

		extraRefs = append(extraRefs, extraRef)
	}

	imagePath := filepath.Join(src, image)
	matches, err := filepath.Glob(imagePath)
	switch {
	case err != nil:
		return err
	case len(matches) == 0:
		return errors.New("no images found")
	case len(matches) > 1:
		return errors.New("more than 1 image found")
	}

	img, err := tarball.ImageFromPath(matches[0], nil)
	if err != nil {
		return err
	}

	digest, err := img.Digest()
	if err != nil {
		return err
	}

	log.Printf("pushing %s to %s", digest, ref.Name())

	return put(c.BasicCredentials(), img, ref, extraRefs)
}

func put(creds BasicCredentials, img v1.Image, ref name.Reference, extraRefs []name.Reference) error {
	auth := &authn.Basic{
		Username: creds.Username,
		Password: creds.Password,
	}

	err := remote.Write(ref, img, remote.WithAuth(auth))
	if err != nil {
		return fmt.Errorf("upload image: %w", err)
	}

	log.Println("pushed")

	for _, extraRef := range extraRefs {
		log.Printf("pushing as tag %s", extraRef.Identifier())

		err = remote.Write(extraRef, img, remote.WithAuth(auth), remote.WithTransport(http.DefaultTransport))
		if err != nil {
			return fmt.Errorf("tag image: %w", err)
		}

		log.Println("tagged")
	}

	return nil
}
