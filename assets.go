// Code generated by go-bindata.
// sources:
// templates/default.nix
// templates/deps.nix
// DO NOT EDIT!

package go2nix

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

func (fi bindataFileInfo) Name() string {
	return fi.name
}
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}
func (fi bindataFileInfo) IsDir() bool {
	return false
}
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _templatesDefaultNix = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x6c\x91\x41\x8f\xa2\x40\x10\x85\xef\xfc\x8a\x8a\xee\x51\x20\x31\xd9\x3d\xb8\xd9\x9b\xc9\xde\xd4\x8c\x33\x73\x31\x1c\x8a\xee\xb2\xbb\x03\x34\xa6\xbb\x61\x9c\x21\xfc\xf7\x29\x04\x1d\x4d\xbc\x15\x55\xdf\xab\xf7\xe8\x9a\xc3\xab\x36\x1e\x8e\xa6\x24\xf8\x40\x0f\x8a\x2c\x39\x0c\x24\x21\xff\x04\x1d\xc2\xc9\xaf\xd2\x54\x99\xa0\x9b\x3c\x11\x75\x95\x16\x58\x99\x52\xe8\x2a\x55\xf5\xd2\x9a\x33\xb4\x87\x03\x24\xef\xe4\xbc\xa9\x2d\x64\x59\xd4\x81\x0f\x92\x6c\xbb\x80\xbc\x31\xa5\xfc\x5f\xef\x50\x14\xa8\x68\x01\x47\x0a\x42\xf3\xa6\xa9\xd2\x6a\x2a\xf2\x2f\x37\x55\xbe\xb5\xd0\xaf\xa2\xe8\x51\x09\x8e\x04\x74\x11\x80\xc5\x8a\xe0\x1f\xcc\x06\xc7\x5d\xa1\x92\x8d\x39\x6f\x86\x56\x96\xc5\x8d\xf5\x01\xf3\x92\xe2\x5f\x5d\x3b\x66\xe9\x67\x7f\x59\x32\x7d\xdc\xab\x6e\x7b\xef\x52\x33\xcb\xb0\xba\x4e\x76\x18\xf4\x83\xd1\xe8\x32\x52\xde\x09\x9e\x5d\x7f\xe6\x12\x0c\x38\x62\x7b\x2f\xd8\xd7\x8d\x13\x94\xbc\x50\x6b\x7e\x1c\x06\xae\x71\xe5\x13\xee\x8d\xbb\x37\xc4\x6b\x5c\xfe\xfe\xf3\x84\xda\x8f\x83\x09\xec\xa7\xc8\x6b\x3a\x79\x86\x93\x54\x72\x91\xf0\x45\x2e\x7d\xbe\xea\x76\xbd\x5d\x01\x4a\x09\x15\x05\x94\x18\xf0\x76\x4d\x86\x6a\x9f\xd4\x4e\x0d\xd5\xa9\x50\x3e\xad\xd0\x36\x58\xa6\x73\x4f\x22\xe6\x87\xb4\x12\x9d\x8c\x07\x5d\x8c\x21\x38\x93\x37\x81\x3c\x6f\x1d\x3a\xec\xd5\x8d\xf6\x7d\xf4\x1d\x00\x00\xff\xff\xc1\x1b\x2d\xbf\x3c\x02\x00\x00")

func templatesDefaultNixBytes() ([]byte, error) {
	return bindataRead(
		_templatesDefaultNix,
		"templates/default.nix",
	)
}

func templatesDefaultNix() (*asset, error) {
	bytes, err := templatesDefaultNixBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/default.nix", size: 572, mode: os.FileMode(420), modTime: time.Unix(1514990199, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesDepsNix = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x5c\xcf\x41\x4b\xc3\x30\x1c\x05\xf0\x7b\x3f\xc5\x63\x9e\x6d\x60\xa0\x07\xc5\x9b\x67\x19\x6e\x7a\x09\x3d\xfc\x97\xfd\x97\x84\x35\x49\x49\xd2\xea\x90\x7d\x77\x9b\x55\xcb\xe8\x21\x04\x5e\x7e\x0f\xf2\xee\xb0\x33\x36\xe1\x68\x5b\xc6\x17\x25\x68\xf6\x1c\x29\xf3\x01\xfb\x33\x4c\xce\x5d\x7a\x12\x42\xdb\x6c\xfa\x7d\xad\x82\x13\x27\x72\xb6\x55\xc6\x09\x1d\xd6\xde\x7e\x63\x90\x12\xf5\x27\xc7\x64\x83\x47\xd3\x54\xb2\x92\xf2\x1e\x91\xbc\x66\xd4\xaf\xdc\xa5\x12\x02\x3f\xe3\x01\x74\xd8\x90\x3a\x91\xe6\x0d\x65\x83\x17\xac\x4a\xf9\x8d\x1c\x8f\x68\xf5\x7c\x25\x47\xce\xaa\x3c\x4d\x05\x20\x9f\x3b\xfe\x97\xdb\xd0\x47\xc5\xf5\xae\x44\x73\x01\xe8\x63\xbb\x10\x1f\x63\x72\x03\x22\x0f\x0b\xf0\xce\x83\xfd\xfb\xf1\xac\x92\xa1\xf5\xc3\xe3\x02\x6e\xa7\x70\x66\x97\x72\x5d\xae\x1b\xd9\x1f\xca\xb6\xa6\xfa\x0d\x00\x00\xff\xff\xe9\x00\xd5\xc8\x43\x01\x00\x00")

func templatesDepsNixBytes() ([]byte, error) {
	return bindataRead(
		_templatesDepsNix,
		"templates/deps.nix",
	)
}

func templatesDepsNix() (*asset, error) {
	bytes, err := templatesDepsNixBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/deps.nix", size: 323, mode: os.FileMode(420), modTime: time.Unix(1514911321, 0)}
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
	"templates/default.nix": templatesDefaultNix,
	"templates/deps.nix": templatesDepsNix,
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
	"templates": &bintree{nil, map[string]*bintree{
		"default.nix": &bintree{templatesDefaultNix, map[string]*bintree{}},
		"deps.nix": &bintree{templatesDepsNix, map[string]*bintree{}},
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

