package docker

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/filters"
	docker "github.com/docker/docker/client"
)

// Client for handling requests to the Docker API
type Client struct {
	client         *docker.Client
	authentication string
}

// Authentication sets credentials to be used by the client
func Authentication(user, password string) func(*Client) error {
	return func(c *Client) error {
		authConfig := types.AuthConfig{
			Username: user,
			Password: password,
		}

		encodedJSON, err := json.Marshal(authConfig)
		if err != nil {
			return err
		}
		c.authentication = base64.URLEncoding.EncodeToString(encodedJSON)

		return nil
	}
}

// NewClient builds the client used to access the Docker API
func NewClient(options ...func(*Client) error) (*Client, error) {
	c := &Client{}

	for _, option := range options {
		if option == nil {
			return nil, fmt.Errorf("option is nil pointer")
		}

		err := option(c)
		if err != nil {
			return nil, err
		}
	}

	cli, err := docker.NewClientWithOpts(docker.FromEnv, docker.WithAPIVersionNegotiation())
	if err != nil {
		return nil, err
	}

	c.client = cli

	return c, nil
}

// Pull a docker image to the host
func (c *Client) Pull(image string) error {
	ctx := context.Background()

	opts := types.ImagePullOptions{}
	if c.authentication != "" {
		opts.RegistryAuth = c.authentication
	}

	_, err := c.client.ImagePull(ctx, image, opts)
	if err != nil {
		return err
	}

	return nil
}

// Image returns metadata about an image on the docker host
func (c *Client) Image(image string) (types.ImageSummary, error) {
	ctx := context.Background()

	f := filters.NewArgs()
	f.Add("reference", image)

	opts := types.ImageListOptions{Filters: f}

	list, err := c.client.ImageList(ctx, opts)
	if err != nil {
		return types.ImageSummary{}, err
	}

	log.Println(list)

	if len(list) < 1 {
		return types.ImageSummary{}, errors.New("image not found")
	}

	return list[0], nil
}

// Save an image by ID as tar to output directory
func (c *Client) Save(output, imageID string) error {
	ctx := context.Background()

	responseBody, err := c.client.ImageSave(ctx, []string{imageID})
	if err != nil {
		return err
	}
	defer responseBody.Close()

	return copyToFile(output, responseBody)
}

func copyToFile(outfile string, r io.Reader) error {
	tmpFile, err := ioutil.TempFile(filepath.Dir(outfile), ".docker_temp_")
	if err != nil {
		return err
	}

	tmpPath := tmpFile.Name()

	_, err = io.Copy(tmpFile, r)
	tmpFile.Close()

	if err != nil {
		os.Remove(tmpPath)
		return err
	}

	if err = os.Rename(tmpPath, outfile); err != nil {
		os.Remove(tmpPath)
		return err
	}

	return nil
}
