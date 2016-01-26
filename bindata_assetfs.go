// Code generated by go-bindata.
// sources:
// public/index.html
// DO NOT EDIT!

package main

import (
	"github.com/elazarl/go-bindata-assetfs"
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

var _publicIndexHtml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x94\x57\x6d\x6f\xdb\xb6\x13\x7f\xed\x7c\x0a\x56\xff\x17\x96\x51\x5b\x4a\xd2\xfe\xb1\xa5\x95\xbd\x17\x69\xbb\x07\x74\x6d\x31\x67\x43\x87\x22\x18\x18\xe9\x14\x31\xa1\x48\x8d\xa4\xe3\x7a\xab\xbf\xfb\x8e\x0f\x92\xa5\xd8\x4d\xd0\x00\x89\xc8\x7b\xf8\xdd\x91\xf7\xbb\x93\x92\x3d\x79\xf5\xfe\xfc\xe2\xcf\x0f\xaf\x49\x65\x6a\xbe\x38\xca\xec\x83\x70\x2a\xae\xe7\x11\x88\x68\x71\x34\xca\x2a\xa0\x05\x3e\x47\x59\x0d\x86\x92\xbc\xa2\x4a\x83\x99\x47\xbf\x5f\xbc\x99\x7d\x1f\x39\x85\x61\x86\xc3\xe2\x0f\xa6\x21\x4b\xfd\xda\x4a\x9f\xcc\x66\xe4\x5c\xd6\x0d\xe3\x50\x10\x2a\x0a\x52\x33\xc1\x4a\x86\x9b\xf3\xe5\x92\xcc\x66\xce\x88\x33\x71\x4b\x14\xf0\x79\xa4\xcd\x86\x83\xae\x00\x4c\x44\x2a\x05\xe5\x3c\xaa\x8c\x69\xf4\x8b\x34\xcd\x0b\x71\xa3\x93\x9c\xcb\x55\x51\x72\xaa\x20\xc9\x65\x9d\xd2\x1b\xfa\x39\xe5\xec\x4a\xa7\x35\x35\xa0\x18\xe5\xec\x1f\x48\x8f\x93\xb3\xef\x92\xff\xa7\xb9\x1e\x88\x13\x8c\x9c\xa0\xcc\x67\xab\x73\xc5\x1a\x43\xb4\xca\x7b\x21\x64\x01\xc9\xcd\xdf\x2b\x50\x1b\x87\xee\x97\xb3\xd3\xe4\x34\x39\x76\xde\x37\xe8\x9c\xa5\xde\xf5\xeb\x28\xdf\x9c\xe8\xcd\x7e\x9e\xf7\x22\x65\x69\xb8\xff\xec\x4a\x16\x1b\x17\xba\x60\x77\x24\xe7\x54\xeb\x79\x94\x4b\x61\x28\x13\xa0\xdc\xd1\xb0\x56\x27\xa1\x0c\xb8\x70\x82\x52\xaa\x9a\xb0\x62\x1e\xd1\xdc\x30\x29\xde\xe0\x36\x22\x58\xc8\x4a\xa2\xec\xc3\xfb\xe5\x45\x44\x40\xe4\x66\xd3\xc0\x3c\xaa\x57\xdc\xb0\x86\x2a\x93\x5a\xaf\x59\x41\x0d\xf5\xb0\x83\x90\x25\xd6\x73\x86\x65\xe4\x05\x61\xa2\x59\x19\xbf\x0e\x86\x03\xcb\x2b\x23\x5a\x31\xde\x57\x43\xc5\xe2\x0d\xfa\xe2\xd9\xec\xb2\x95\x3b\x0c\xe2\x13\xb0\xd0\x11\x11\xb4\x6e\xd7\x8b\xec\x4a\xb9\xdf\x00\x9e\x22\xfa\x81\x40\x2e\xa5\x86\x9a\x6a\xb6\x56\xb4\x69\xda\xdb\xd8\xc1\xdf\xb7\x8b\x42\x40\x03\x9f\x91\x6d\x0d\xa7\x39\x54\x92\x17\xa0\x90\xd6\x0d\x97\xb4\x20\x3e\xfc\x5e\xd8\xfe\xb2\x09\x4f\x1f\xc2\xde\x71\x41\x37\xcf\x5a\x68\x45\x0b\x26\xdb\xc3\xa0\x42\x47\xe4\x8e\xf2\x15\x6e\x9e\xb5\x97\xca\xe9\x15\x70\x82\x77\x1d\x3c\x17\xcf\x88\x35\xcc\x52\xa7\x68\x03\x36\x0f\x85\x3b\x7b\x3c\xdc\xd9\xe1\x70\x28\x3e\xfb\xe6\x70\x27\xa7\x8f\xc7\x43\x9b\x83\x01\xad\xfc\xe4\xf4\xa1\x90\x6d\x9d\xb3\x4a\x0d\x82\x87\xea\xad\xe9\x1d\xe8\x19\x94\x25\xe4\x86\xf8\x0d\x67\xd7\x95\x21\x96\x67\x21\x2d\xbd\xba\xaa\x99\xe9\x72\x59\xfa\xad\x6f\x05\xc7\x6a\xbf\xec\x71\x47\xc9\xf5\x01\x92\xe7\x92\x13\xdd\x9d\x04\x33\x3a\x5d\x78\x62\xe0\xf4\xb2\x24\xc6\x23\xa0\xa8\xc7\x44\x7b\x41\x76\x96\xb9\xe6\x3d\x40\x97\x6e\xb5\x5b\xb4\x2d\x1e\x97\x2b\xe1\x9a\x33\x9e\x90\x7f\x8f\x46\x6b\x26\x0a\xb9\x4e\xa4\x70\x3c\x9c\x93\xa1\x76\xd4\x6e\x49\x4d\x6f\xe1\x63\xa5\x62\xdf\xcc\x53\xb2\x52\x7c\x4a\x1a\xba\xb1\x6e\x53\x52\x0a\x6f\x3e\xba\xa3\x8a\x7c\xae\x14\x02\x09\x58\x93\x8f\xbf\xbe\xfd\x09\x87\xd5\x6f\x80\xf3\x4d\x9b\x78\xf2\xd2\x9a\xa0\x3a\x91\x0d\x88\x3e\x54\x4f\x23\x14\x0e\xa0\x8d\x36\x38\xa7\x70\xf8\x8b\x6b\xd8\x4f\x6a\x34\x62\x25\x89\xad\xb5\xb3\x5d\x5a\x5b\x32\x9f\x93\xe7\xad\x7a\x54\x8a\xf8\x97\xe5\xfb\x77\x49\x63\xdf\x1e\xc1\x52\x37\x52\x68\xb8\xc0\x1e\x9c\xf8\x70\xa3\xad\xfd\xbb\x75\xeb\x70\x12\xf2\x83\xcd\x3e\xd1\x20\x8a\x38\x88\x26\xe4\xc5\x4e\xe6\x1c\xb7\xfd\x7b\xb1\x9d\xeb\x6b\x15\xdb\x65\xef\x1a\xca\x22\xdc\x82\x9d\x82\xaf\x70\xba\x85\xf3\x97\x45\x62\xa7\x06\xa2\x8d\xad\xc7\x78\xea\x30\xf6\x74\x96\xb9\xa8\x2b\x64\xbe\xaa\x41\x98\xc4\xbd\x22\x96\xc0\x91\x8e\x52\xc5\x63\xc7\xd5\x4f\xae\x21\xac\xe5\xe5\x8b\xbc\x82\xfc\x16\x8a\xf1\x24\x71\x6c\xdc\xc3\x5b\x69\x50\x33\x23\x6f\x41\x20\x2a\x97\x39\xe5\x4b\x04\xa2\xd7\x90\x5c\x83\xf9\xd9\x40\x1d\x47\x3b\x93\x68\x42\xbe\x7c\x21\x51\xe4\x51\xda\xda\xfb\x09\x3e\x25\x51\x4a\x1b\x96\x6a\xec\x09\xdc\x94\xb6\xfc\x6d\x7d\xec\x10\xef\xd7\xe8\xc9\xe3\x81\xba\x9a\x59\x7b\xeb\x9f\x58\xed\x5f\x4e\x3b\x09\x73\x75\x80\xa2\x03\x4a\xd0\x8d\x46\x7d\xb8\x69\x27\xbd\x07\xd5\xc9\xfb\xb5\x1f\xd9\xba\xbd\xb5\x7d\x14\x6a\xb3\xed\x0a\x6c\x2b\xb8\x7b\x89\x61\x25\xbf\x56\x88\xff\xed\xac\xc6\xce\x7b\xb7\x47\x32\xfb\x09\xd1\xa7\x70\x4b\x11\x48\x1a\x05\x77\x08\xf8\x0a\x4a\x8a\xaf\xc2\x96\x1d\x3b\x3e\x3d\x5c\x7a\x37\x81\xac\xf5\x25\x96\xdc\x3e\xf5\xa7\xe3\xcb\x7d\x7e\xe6\x1c\xa8\x7a\x87\xdf\x1b\x71\xdb\xa2\xeb\x0a\x8d\x63\x81\x3e\x4a\x9b\x73\xdc\x14\x6d\x09\x04\x76\x49\x2d\xef\xc0\x09\x87\x16\x47\xe1\xca\x06\xd8\xbd\xdb\xf3\x08\xb6\x84\x07\x2b\xde\x67\x5f\x57\xf1\x8e\x56\x3f\xbe\xbe\xe8\x0a\xe7\xb9\x15\x91\xa7\x87\x39\x3a\x00\x42\xa3\x28\xf5\x63\x70\x4a\x82\xbf\x58\x71\xde\x6d\x0e\xf2\x32\xb4\xa7\xf3\x7b\xa8\xae\xce\x60\x1c\xe8\x82\x3f\xbb\x9b\x74\x9a\x9d\x02\x07\x7d\x6c\x11\x19\xa2\x1d\xbf\xc4\x47\x46\x1c\xf9\x9c\x59\xc2\x41\x5c\x9b\x0a\xc5\x4f\x9f\xf6\x32\x70\x29\xe0\xe0\xb3\x09\x74\xb6\x9f\xd8\xe5\xcb\x81\x01\x15\x79\x25\x55\x3f\xc9\x1c\x27\x9e\x81\xd7\x1c\xec\x2e\x1e\xd3\x5e\x7e\xce\xc3\xbe\x1c\xbe\x6e\x8e\xda\xbe\x83\x87\x4f\xec\xe7\x2f\x3a\x61\x36\x7b\x2a\xfb\xc5\x72\x8e\x1f\x7d\xe8\x7d\x3f\xd3\x04\xa9\xad\x8d\x8a\x4f\x9e\xf7\x10\x31\x40\x18\x38\x9e\x43\x1e\xa6\x67\xe0\xef\xa4\x6f\x82\x2e\x3b\xfd\x36\x2c\xb6\xa1\x1d\x03\xe3\x86\x6d\xba\x3d\xda\x4e\x7c\xb7\x0c\x3e\x5c\xfd\x07\x2b\xbe\x29\xdd\x7f\x16\xff\x05\x00\x00\xff\xff\x14\x26\xf7\x49\x6a\x0c\x00\x00")

func publicIndexHtmlBytes() ([]byte, error) {
	return bindataRead(
		_publicIndexHtml,
		"public/index.html",
	)
}

func publicIndexHtml() (*asset, error) {
	bytes, err := publicIndexHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "public/index.html", size: 3178, mode: os.FileMode(436), modTime: time.Unix(1453773298, 0)}
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
	"public/index.html": publicIndexHtml,
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
	"public": &bintree{nil, map[string]*bintree{
		"index.html": &bintree{publicIndexHtml, map[string]*bintree{}},
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


func assetFS() *assetfs.AssetFS {
	for k := range _bintree.Children {
		return &assetfs.AssetFS{Asset: Asset, AssetDir: AssetDir, AssetInfo: AssetInfo, Prefix: k}
	}
	panic("unreachable")
}
