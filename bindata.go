// Code generated by go-bindata.
// sources:
// templates/index.html
// DO NOT EDIT!

package main

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

var _templatesIndexHtml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x8c\x55\xdd\x8e\x9b\x3c\x10\xbd\xcf\x53\xcc\xe7\x4f\x95\x12\x75\x03\x21\xfb\xa3\x8a\x00\x17\xad\xaa\xaa\xd2\x5e\xed\x6a\x1f\xc0\xb1\x87\x60\xad\xb1\x91\x6d\xf2\xd3\x88\x77\xaf\x08\x6c\x02\x4e\x5a\xd5\xb9\x19\x7b\xce\xcc\x99\x19\x8e\x9d\xa4\x70\xa5\xcc\x26\x93\xa4\x40\xca\xb3\x09\x00\x40\x52\xa2\xa3\xa0\x68\x89\x29\xd9\x0a\xdc\x55\xda\x38\x02\x4c\x2b\x87\xca\xa5\x64\x27\xb8\x2b\x52\x8e\x5b\xc1\x70\x7e\xda\xdc\x81\x50\xc2\x09\x2a\xe7\x96\x51\x89\x69\x44\xfa\x44\xd6\x1d\x24\x76\x76\xbb\xd6\x9a\x1f\xe0\x78\xde\xb6\x2b\xd7\xca\xc5\x10\x3d\x54\xfb\x30\x0a\x96\x8f\x58\x82\xa5\xca\xce\x2d\x1a\x91\xaf\x46\xc8\x92\x9a\x8d\x50\x31\x3c\x2c\xaa\x3d\xd0\xda\x69\xdf\xbd\xef\x8a\x89\xe1\xe9\x71\x51\xed\xc7\x5e\x29\x14\xce\x0b\x14\x9b\xa2\x65\x0b\x9e\x56\x57\x45\xcc\xad\xf8\x85\x31\x44\x5f\xfc\x50\xa6\xa5\x36\x31\xfc\x1f\xad\xdb\xdf\xd8\x57\x51\xce\x85\xda\xc4\xb0\x80\x68\x51\xed\xcf\xbe\x66\x72\x36\x8b\xe8\xee\x62\x2f\x07\xf6\xbd\x37\x09\xaf\xc4\xe5\xad\x64\xd4\x8b\x71\xb8\x77\x73\x8e\x4c\x1b\xea\x84\x56\x31\x28\xad\xf0\x56\xa0\x50\x55\xed\x2e\xe4\xeb\xda\x39\xad\xbc\x64\xfd\xf4\xa2\xc5\xe2\xd3\xb8\xcb\xb5\x36\x1c\x4d\x0c\x91\x3f\x9a\x73\xfb\x0f\xbe\x67\x30\xd1\xfb\xc7\xa1\xb3\xe9\x84\x11\xf6\xca\x48\xc2\x4e\x76\x93\xa4\xd5\x46\xaf\x9a\xf6\x08\xcd\x45\x36\x09\x17\x5b\x60\x92\x5a\x9b\x12\xa1\x9c\xd1\x24\x1b\x91\x25\x45\x94\xbd\x16\xda\x38\x54\xf0\xf6\xf2\x9c\x84\x45\x34\x06\x1c\x8f\x20\x72\x08\x6c\x87\x41\x0e\x4d\xe3\x25\x58\x66\x09\x85\xc2\x60\x9e\x92\xe3\x11\x82\x42\x5b\x07\x4d\x13\xb6\xf6\x30\x8a\x64\xfe\x49\x12\xd2\x2c\x09\x8b\xe5\x15\x21\x4a\x8b\x27\x56\x34\x46\x9b\x5b\x8c\x6d\xaa\x0f\xe7\x9f\x53\xf8\x81\x6b\x73\x8d\x53\xdc\x47\x55\xde\x84\x42\xff\xe0\x24\x08\x10\x3c\x25\xb5\x91\x7d\x43\x04\x2a\x49\x19\x16\x5a\x72\x34\x29\xc1\x3d\x2d\x2b\x89\x01\xd3\x25\x01\xad\xde\xf1\x50\x19\x6c\xbf\x81\x41\x57\x1b\x05\xa6\x56\xaf\xcc\x88\xca\x4d\x71\x8b\xca\xcd\x08\x84\xd9\xdf\x6b\xbd\x71\xd0\x09\xd1\x1d\x2a\x4c\x49\xb7\x69\xb9\x98\x14\xec\x3d\x25\x7d\x5d\x6f\x2f\xcf\xd3\x19\xc9\x7e\xe8\xff\x92\xb0\xc3\x0c\xb4\x11\x72\xb1\xcd\x26\x37\xb5\xc2\x24\x52\x43\xb2\x21\xa4\xd3\xdb\x87\xb8\x12\x7b\xaa\xff\x92\x8d\x6b\x56\x97\xa8\x5c\xb0\x41\xf7\x5d\x62\x6b\x7e\x3d\xfc\xe4\xd3\xe1\x90\x66\x41\xae\x59\x6d\xa7\xb3\xd5\x85\x35\xaf\x15\x6b\x2f\xe0\x70\x26\x33\xef\x7a\x89\x1c\xa6\x18\xbc\xe3\xe1\x9b\xe6\x08\x69\x0a\xd1\xbd\x0f\x69\xd7\xb0\xe7\xd5\x95\xb7\x1f\x7d\x4e\xa5\xc5\xb1\xb7\xb9\x75\xed\xcf\x75\x0d\xb3\x7a\xa4\x5b\x6a\xc0\xad\x21\xfd\xb7\xe6\xfd\xb7\x51\x59\x2d\x31\x90\x7a\x33\x75\xeb\x60\x4b\x65\x8d\x1e\x64\x27\x14\xd7\xbb\x40\x6a\x76\x7a\xa2\x20\x85\xd1\x0d\x23\xf0\x19\x3e\x22\xc7\x81\xc3\x56\xaf\x9f\x8f\xfe\xc3\x4d\x92\xb0\x7b\x38\xda\x97\xe4\xf4\x47\xf6\x3b\x00\x00\xff\xff\x99\xf8\xbe\x7c\xd0\x06\x00\x00")

func templatesIndexHtmlBytes() ([]byte, error) {
	return bindataRead(
		_templatesIndexHtml,
		"templates/index.html",
	)
}

func templatesIndexHtml() (*asset, error) {
	bytes, err := templatesIndexHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/index.html", size: 1744, mode: os.FileMode(436), modTime: time.Unix(1485635817, 0)}
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
	"templates/index.html": templatesIndexHtml,
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
		"index.html": &bintree{templatesIndexHtml, map[string]*bintree{}},
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
