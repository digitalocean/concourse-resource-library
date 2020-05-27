// Code generated for package bootstrap by go-bindata DO NOT EDIT. (@generated)
// sources:
// bootstrap/src/.gitattributes
// bootstrap/src/.gitignore
// bootstrap/src/Dockerfile
// bootstrap/src/Makefile
// bootstrap/src/README.md
// bootstrap/src/cmd.go
// bootstrap/src/resource.go
package bootstrap

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func bindataRead(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}
	if clErr != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

type asset struct {
	bytes []byte
	info  os.FileInfo
}

type bindataFileInfo struct {
	name    string
	size    int64
	mode    os.FileMode
	modTime time.Time
}

// Name return file name
func (fi bindataFileInfo) Name() string {
	return fi.name
}

// Size return file size
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}

// Mode return file mode
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}

// Mode return file modify time
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}

// IsDir return file whether a directory
func (fi bindataFileInfo) IsDir() bool {
	return fi.mode&os.ModeDir != 0
}

// Sys return file is sys mode
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _bootstrapSrcGitattributes = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x6c\x8f\x41\x6e\xc2\x30\x10\x45\xf7\x91\x72\x87\x2f\x65\x87\x54\x6e\xc0\xa2\x82\x54\x54\xa2\xb0\x68\x7a\x00\x63\x4f\xc8\x88\x89\x1d\xd9\xe3\x02\xb7\xaf\x42\xa8\xaa\x36\xdd\xbe\x3f\xcf\xfe\xbf\xaa\xf0\xfc\xd1\x1c\x9e\x36\x75\x53\xaf\x9b\xb2\xa8\x2a\x00\x5b\xe3\x9d\x10\x84\x3d\x81\xbc\x63\x7f\x4a\x30\x59\x43\x6f\x94\xad\x11\xb9\xa1\x0d\x11\x2d\x0b\x25\x38\x52\xb2\x4a\x0e\x26\x3d\x6c\xa5\xab\xc2\x78\x07\x21\xf3\x49\x30\x22\xf3\x53\x1c\xd9\x9b\x78\x43\xf6\x1a\xb2\xed\xc8\x2d\x1f\x72\xd3\x71\xc2\x85\x45\xd0\x4d\x25\x7e\xf4\xfd\xa1\x81\xa3\x96\x3d\x39\x1c\x49\xc2\x65\x59\x16\x8b\xfb\x6f\xab\xb1\x5c\x59\x94\x45\x85\xfa\x3a\x08\x5b\x56\xb9\xc1\x76\x64\xcf\x21\xeb\x43\xbf\x2b\xb8\xb0\x76\xd8\xbd\x4c\xdb\x1c\x09\xf7\xac\x14\xc7\xea\xd8\x92\x0c\x69\x3a\xb0\x31\xa4\x84\x41\x8c\xb6\x21\xf6\x48\x79\x18\x42\xd4\xb2\x58\x2c\xf5\xaa\xdf\x13\x29\xc8\x4a\xda\x11\xf6\x0e\x73\x78\x0a\xff\xc0\x3e\xb8\x39\x4c\xb9\xff\x0b\xcb\x62\xf7\xba\xae\xf7\xef\x35\x66\x8f\xbc\x99\x33\x8d\x8b\x66\xc1\x26\xd8\x33\xc5\x29\xfa\x15\x7c\x05\x00\x00\xff\xff\xe6\xf2\xe1\x7e\xe4\x01\x00\x00")

func bootstrapSrcGitattributesBytes() ([]byte, error) {
	return bindataRead(
		_bootstrapSrcGitattributes,
		"bootstrap/src/.gitattributes",
	)
}

func bootstrapSrcGitattributes() (*asset, error) {
	bytes, err := bootstrapSrcGitattributesBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "bootstrap/src/.gitattributes", size: 484, mode: os.FileMode(438), modTime: time.Unix(1590538435, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _bootstrapSrcGitignore = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x2c\xcf\x41\x6a\xf3\x30\x10\x05\xe0\xbd\xc1\x77\x78\x90\xcd\xff\x87\x34\xbd\x43\x49\x17\x85\x42\x37\x3d\x40\x6c\x69\x2c\x0f\x8c\x35\x42\x1a\xd9\xf5\xa6\x67\x2f\x72\xb2\x7a\x12\xc3\xfb\x86\x39\xe1\x8d\xe3\x90\x99\x0a\x26\xcd\x48\x59\x43\x1e\x96\x82\x21\x7a\x24\xa9\x81\x63\xe9\xbb\xf3\x95\x7e\xe8\x19\xbf\x2d\xbd\x48\x8b\xa2\xc7\x67\x17\x1e\xfb\xae\xef\x4e\xf8\xa6\x62\x18\x1b\xb8\x5f\x30\x56\x16\xc3\xc6\x36\xe3\x1e\x14\xd6\x66\x2f\xee\xde\x2a\xed\xfd\x68\x7c\x55\x4b\xd5\xa0\x13\x6c\x26\x04\x85\xd3\x95\xf2\x10\x08\xa6\x2a\x17\x94\x44\x8e\x27\x76\x83\xc8\x8e\x6d\xa6\x88\x5a\xc8\x3f\xd4\x4f\x36\xfa\xb8\xbd\x37\x50\xeb\xd3\xbb\x51\xa2\xe8\x29\xba\x1d\x9e\x33\x39\xd3\xe3\xb6\x7f\x99\x16\x5d\xe9\x58\xe2\x74\x59\x28\x1a\x46\x12\xdd\x60\x0a\x8e\x4e\xaa\x27\xb0\xfd\x6f\xc4\x4a\xd1\x6b\x7e\x6d\xde\xf9\x2a\x1a\xfa\xee\x2f\x00\x00\xff\xff\x4f\xce\x88\xc4\x25\x01\x00\x00")

func bootstrapSrcGitignoreBytes() ([]byte, error) {
	return bindataRead(
		_bootstrapSrcGitignore,
		"bootstrap/src/.gitignore",
	)
}

func bootstrapSrcGitignore() (*asset, error) {
	bytes, err := bootstrapSrcGitignoreBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "bootstrap/src/.gitignore", size: 293, mode: os.FileMode(438), modTime: time.Unix(1590538435, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _bootstrapSrcDockerfile = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x01\x00\x00\xff\xff\x00\x00\x00\x00\x00\x00\x00\x00")

func bootstrapSrcDockerfileBytes() ([]byte, error) {
	return bindataRead(
		_bootstrapSrcDockerfile,
		"bootstrap/src/Dockerfile",
	)
}

func bootstrapSrcDockerfile() (*asset, error) {
	bytes, err := bootstrapSrcDockerfileBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "bootstrap/src/Dockerfile", size: 0, mode: os.FileMode(438), modTime: time.Unix(1590538435, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _bootstrapSrcMakefile = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xca\xcd\x4e\xcb\xcc\x49\x55\xb0\xb2\x55\x50\xd1\x48\x4c\x2a\x2e\x48\x2c\xc9\x50\x50\xd1\xc8\x49\x2c\x2e\x29\xcf\x2f\x4a\x51\x50\xd1\xf0\x75\xf4\x76\x75\xf3\xf4\x71\x8d\xf7\xf1\x0c\x0e\xd1\xd4\xd4\xe4\x4a\xc9\x2c\x82\x28\x07\x31\x54\x34\x20\x06\x68\x6a\x72\x71\xa5\x56\x14\xe4\x17\x95\x28\xf8\xf8\xbb\xc7\x87\x04\x85\xfa\x39\x3b\x86\xb8\xda\x96\x14\x95\xa6\x22\x4b\xb8\x78\x06\xb9\x3a\x87\xf8\x07\x45\xda\x82\xf5\x6b\x72\x71\xe9\x05\x78\xf8\xfb\x45\x5a\x29\x94\xa4\x16\x97\x70\x81\x08\x2b\x2e\x4e\x87\xf4\x7c\x30\x5f\x41\x57\x37\x39\xbf\x2c\xb5\x48\xa1\xba\x5a\x41\xcf\x37\x3f\xa5\x34\x27\x55\xa1\xb6\x56\x5f\x4f\x4f\x8f\x0b\x10\x00\x00\xff\xff\xbb\xfd\xa6\x16\xb8\x00\x00\x00")

func bootstrapSrcMakefileBytes() ([]byte, error) {
	return bindataRead(
		_bootstrapSrcMakefile,
		"bootstrap/src/Makefile",
	)
}

func bootstrapSrcMakefile() (*asset, error) {
	bytes, err := bootstrapSrcMakefileBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "bootstrap/src/Makefile", size: 184, mode: os.FileMode(438), modTime: time.Unix(1590538435, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _bootstrapSrcReadmeMd = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x52\x56\xa8\xae\x56\xd0\xf3\xcd\x4f\x29\xcd\x49\x55\xa8\xad\xe5\xe2\x72\xce\xcf\x4b\xce\x2f\x2d\x2a\x4e\x55\x28\x4a\x2d\xce\x2f\x2d\x4a\x4e\x55\x48\xcb\x2f\x52\xd0\xd3\xd3\xe3\xe2\x52\x56\x56\x70\xce\xcf\x4b\xcb\x4c\x07\x33\xdd\x53\x4b\xc0\x74\x40\x29\x84\x76\xad\x48\xcc\x2d\xc8\x49\xe5\xe2\x4a\x48\x48\xa8\x4c\xcc\xcd\x01\xd1\x5c\x80\x00\x00\x00\xff\xff\x92\xae\x22\x55\x60\x00\x00\x00")

func bootstrapSrcReadmeMdBytes() ([]byte, error) {
	return bindataRead(
		_bootstrapSrcReadmeMd,
		"bootstrap/src/README.md",
	)
}

func bootstrapSrcReadmeMd() (*asset, error) {
	bytes, err := bootstrapSrcReadmeMdBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "bootstrap/src/README.md", size: 96, mode: os.FileMode(438), modTime: time.Unix(1590538435, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _bootstrapSrcCmdGo = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x94\x90\xb1\x8a\xe3\x30\x10\x86\x6b\xcd\x53\xcc\x09\x0e\x64\xb8\x38\x7d\x20\x55\xe0\xba\x83\x23\x57\x5c\xad\x48\x63\x47\x44\xd6\x78\xc7\xf2\x2e\x4b\xf0\xbb\x2f\x92\x37\xd9\x22\xdb\xa4\xd4\xfc\x7c\xff\x37\xa3\xd1\xba\x8b\xed\x09\x07\x1b\x12\x40\x18\x46\x96\x8c\x06\x94\x8e\xdc\x6b\x00\x25\x91\x7b\xd4\x7d\xc8\xe7\xf9\xd4\x3a\x1e\xb6\x3e\xf4\x21\xdb\xc8\x8e\x6c\xda\x3a\x4e\x8e\x67\x99\x68\x23\x34\xf1\x2c\x8e\x36\x31\x9c\xc4\xca\xfb\xb6\xf2\xea\x36\x46\x7d\xbd\x62\xfb\x87\xfd\x1c\x09\x97\x45\x43\x03\xd0\xcd\xc9\x55\xaf\x69\xf0\x0a\x2a\xa4\x71\xce\xb8\xdb\x63\x51\xb6\xff\x25\x64\xfa\x97\x7d\x49\x41\x79\xea\x48\xd6\xe0\x10\x79\x22\xd3\x00\xa8\x57\x2b\x28\xf4\x32\xd3\x94\xf1\xe6\x69\x8b\xe6\x30\x78\x5c\x96\xe3\x1a\x81\x22\x91\x5a\xbb\xbe\xdb\x23\x59\x6f\xaa\xac\x01\x15\x3a\x2c\xf1\x8f\x3d\xa6\x10\xcb\x16\xaa\x38\x7e\xdb\x6c\x63\x67\x74\x67\x43\x24\x8f\x99\x51\xc8\xfa\xbb\xac\xc2\x3b\xfc\x39\xe9\x5f\x85\x6e\x40\x2d\x50\x4f\x1d\x39\x4d\x54\x67\xab\xf0\x61\x27\xf3\x59\xf1\x84\x79\x24\xe9\x58\x06\x74\x67\x72\x97\x07\x69\xa9\xa8\xa6\xaa\x5e\x7f\xcd\x3c\xd1\xfe\x56\x80\x3b\x5f\x26\x53\xf6\xfc\xcd\x75\x05\xff\x2b\x21\xe5\x98\x8c\xfe\x3a\x08\x1d\x0f\x63\xa4\x4c\xba\x81\x05\x3e\x02\x00\x00\xff\xff\xd7\xaf\x4b\xe1\x4d\x02\x00\x00")

func bootstrapSrcCmdGoBytes() ([]byte, error) {
	return bindataRead(
		_bootstrapSrcCmdGo,
		"bootstrap/src/cmd.go",
	)
}

func bootstrapSrcCmdGo() (*asset, error) {
	bytes, err := bootstrapSrcCmdGoBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "bootstrap/src/cmd.go", size: 589, mode: os.FileMode(438), modTime: time.Unix(1590543960, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _bootstrapSrcResourceGo = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xac\x94\x31\x6f\xdb\x3e\x10\xc5\x67\xf1\x53\x1c\x34\x49\x7f\x24\xf2\xfe\x5f\x83\x22\x4b\x5b\x04\x09\xda\x0e\x41\x80\xd0\xd4\xd9\x66\x23\x91\xea\xf1\x68\xc3\x08\xfa\xdd\x0b\x52\xa4\x2c\xcb\xee\x50\x20\x93\x09\xfa\xf8\xde\xdd\x8f\x7a\x1c\xa4\x7a\x93\x5b\x04\x42\x67\x3d\x29\x14\x42\xf7\x83\x25\x86\x4a\x14\x25\x1a\x65\x5b\x6d\xb6\xab\x9f\xce\x9a\x52\x14\xa5\x75\xa5\x10\x45\x0f\xe5\x56\xf3\xce\xaf\x1b\x65\xfb\x55\xab\xb7\x9a\x65\x67\x15\x4a\xb3\x52\xd6\x28\xeb\xc9\xe1\x6d\x16\xbc\xed\xf4\x9a\x24\x1d\x57\x3d\xb2\x6c\x25\xcb\x52\xd4\x42\xac\x56\xf0\x14\xff\x06\xc2\x81\xd0\xa1\x61\x07\xbc\x43\x50\xd6\x6c\xf4\xd6\x93\x64\x6d\x0d\x6c\x2c\xc5\xdd\xa9\x3b\x3e\x0e\x98\x4f\x3a\x26\xaf\x18\xde\xc5\xef\xa8\xf7\x5d\x76\xba\x95\x8c\x80\xc6\x79\xc2\x20\x27\x39\x9e\x1e\xcf\x2e\xa4\xb5\x83\x7d\x38\x21\x36\xde\x28\xa8\x5c\x52\xad\x27\x9d\xaa\x06\x24\xb2\x04\xef\xa2\x20\x64\x4f\x06\x8c\xee\xb2\x19\x92\x0b\x2a\xca\x1a\x96\xda\x8c\xbd\xef\xd3\x66\x98\x12\xee\x32\x09\xf0\x2e\x34\x63\xa1\x45\x46\xea\xb5\x41\xd0\x1b\x90\xb0\xf6\xba\x6b\xc1\xed\xac\xef\x5a\x20\x6f\xc6\xd9\xb2\xf0\x62\xb8\xbb\x1d\xaa\xb7\x47\xfc\xe5\xd1\x71\x68\x3d\xd8\x45\x9b\x54\x47\xa8\x50\xef\xb1\x85\x0d\xd9\x7e\xf4\x26\xef\x10\xd6\xc7\x33\x7e\xa0\x82\x0e\xd8\x01\x47\x0a\xa3\xe7\x99\xf8\x64\x5c\x24\xcc\x19\x37\xbc\x86\xaf\xe0\xff\x72\x54\x2a\x5f\x45\x91\x7b\xcd\xbf\xa9\x20\x61\x28\x5f\x53\xef\x8f\x28\x5b\x38\xe8\xae\x03\x0a\xab\xd0\x4f\x28\x0c\x4d\x0d\xd6\x38\x9c\xf5\x1c\x79\xed\x75\x18\xab\xd5\x26\x5d\x0d\xc1\x7f\xf3\x0e\xeb\x28\x58\x69\x33\x78\x86\xe7\x97\xf5\x91\xf1\xf2\xa6\x82\x41\xf3\xcd\xf4\x92\xdc\x4e\x76\x63\xf1\x0d\x50\x7d\x8e\x33\xf9\x5f\xe5\x19\x64\xb0\x0d\xf7\x76\xea\xec\x5f\x68\x26\xed\xe7\x97\x44\x27\xfa\x7e\x46\x93\xa4\x47\x4b\xe3\xfb\x35\x12\xd8\x4d\xfe\x76\x1c\x68\x93\x3d\xa2\xc0\xc4\xe0\x4c\xb6\x0e\x4a\x55\x0d\xda\xf0\x6c\xe8\x0e\x4d\x35\x8d\xf8\x83\x34\xe3\x88\xfd\x10\x97\x97\xdc\xd9\x06\xce\xd6\x73\xcc\xd9\x69\x4c\xb6\x30\x48\xfa\xbb\x77\x94\xbe\x12\x8f\x08\xfd\x2b\x1e\x3e\x85\x67\x03\xa9\xb2\xae\x79\x8a\xfa\x75\x33\x6e\x9d\xba\xbb\x47\x7e\x90\x24\xfb\x10\x09\x97\x2f\xe0\x32\xfb\xf2\xc4\xda\x31\x0e\x23\xe0\xf3\xb3\x8b\xa0\xdc\x23\x7f\x40\x4c\xb6\xc8\xcb\x6b\x9d\x09\x5f\xb3\x4c\x48\x9b\xa6\x99\x95\xa7\xcd\x53\xa4\x72\x52\xa6\xc8\x00\x5c\xa6\xa6\xf8\x92\x5e\x4a\xe8\x9b\x69\x99\xaa\xf2\x23\x7a\x63\x7b\xcd\xd8\x0f\x7c\x9c\x62\xf6\xe0\xe7\x58\xae\xbf\x9c\xe7\x35\x8b\x39\x1e\xfc\x47\xa0\x0b\xa9\x5c\xa0\x9b\x09\xcf\x2c\xff\x04\x00\x00\xff\xff\xaa\x2b\x91\x92\x7a\x06\x00\x00")

func bootstrapSrcResourceGoBytes() ([]byte, error) {
	return bindataRead(
		_bootstrapSrcResourceGo,
		"bootstrap/src/resource.go",
	)
}

func bootstrapSrcResourceGo() (*asset, error) {
	bytes, err := bootstrapSrcResourceGoBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "bootstrap/src/resource.go", size: 1658, mode: os.FileMode(438), modTime: time.Unix(1590545631, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// MustAsset is like Asset but panics when Asset would return an error.
// It simplifies safe initialization of global variables.
func MustAsset(name string) []byte {
	a, err := Asset(name)
	if err != nil {
		panic("asset: Asset(" + name + "): " + err.Error())
	}

	return a
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetNames returns the names of the assets.
func AssetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

// _bindata is a table, holding each asset generator, mapped to its name.
var _bindata = map[string]func() (*asset, error){
	"bootstrap/src/.gitattributes": bootstrapSrcGitattributes,
	"bootstrap/src/.gitignore":     bootstrapSrcGitignore,
	"bootstrap/src/Dockerfile":     bootstrapSrcDockerfile,
	"bootstrap/src/Makefile":       bootstrapSrcMakefile,
	"bootstrap/src/README.md":      bootstrapSrcReadmeMd,
	"bootstrap/src/cmd.go":         bootstrapSrcCmdGo,
	"bootstrap/src/resource.go":    bootstrapSrcResourceGo,
}

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//     data/
//       foo.txt
//       img/
//         a.png
//         b.png
// then AssetDir("data") would return []string{"foo.txt", "img"}
// AssetDir("data/img") would return []string{"a.png", "b.png"}
// AssetDir("foo.txt") and AssetDir("notexist") would return an error
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		cannonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(cannonicalName, "/")
		for _, p := range pathList {
			node = node.Children[p]
			if node == nil {
				return nil, fmt.Errorf("Asset %s not found", name)
			}
		}
	}
	if node.Func != nil {
		return nil, fmt.Errorf("Asset %s not found", name)
	}
	rv := make([]string, 0, len(node.Children))
	for childName := range node.Children {
		rv = append(rv, childName)
	}
	return rv, nil
}

type bintree struct {
	Func     func() (*asset, error)
	Children map[string]*bintree
}

var _bintree = &bintree{nil, map[string]*bintree{
	"bootstrap": &bintree{nil, map[string]*bintree{
		"src": &bintree{nil, map[string]*bintree{
			".gitattributes": &bintree{bootstrapSrcGitattributes, map[string]*bintree{}},
			".gitignore":     &bintree{bootstrapSrcGitignore, map[string]*bintree{}},
			"Dockerfile":     &bintree{bootstrapSrcDockerfile, map[string]*bintree{}},
			"Makefile":       &bintree{bootstrapSrcMakefile, map[string]*bintree{}},
			"README.md":      &bintree{bootstrapSrcReadmeMd, map[string]*bintree{}},
			"cmd.go":         &bintree{bootstrapSrcCmdGo, map[string]*bintree{}},
			"resource.go":    &bintree{bootstrapSrcResourceGo, map[string]*bintree{}},
		}},
	}},
}}

// RestoreAsset restores an asset under the given directory
func RestoreAsset(dir, name string) error {
	data, err := Asset(name)
	if err != nil {
		return err
	}
	info, err := AssetInfo(name)
	if err != nil {
		return err
	}
	err = os.MkdirAll(_filePath(dir, filepath.Dir(name)), os.FileMode(0755))
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(_filePath(dir, name), data, info.Mode())
	if err != nil {
		return err
	}
	err = os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
	if err != nil {
		return err
	}
	return nil
}

// RestoreAssets restores an asset under the given directory recursively
func RestoreAssets(dir, name string) error {
	children, err := AssetDir(name)
	// File
	if err != nil {
		return RestoreAsset(dir, name)
	}
	// Dir
	for _, child := range children {
		err = RestoreAssets(dir, filepath.Join(name, child))
		if err != nil {
			return err
		}
	}
	return nil
}

func _filePath(dir, name string) string {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(cannonicalName, "/")...)...)
}
