// Code generated by go-bindata.
// sources:
// templates/en/eml/confirm_registration.eml
// templates/en/eml/proposal-accepted.eml
// templates/ja/eml/proposal-accepted.eml
// DO NOT EDIT!

package assets

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

var _templatesEnEmlConfirm_registrationEml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x4c\xcd\xb1\x4a\x04\x31\x10\xc6\xf1\xfe\x9e\xe2\xf3\x1e\xe0\xd2\x0b\x62\x61\x69\x63\x61\x2f\xb9\xec\xdc\x66\x48\x2e\x23\x93\x09\x6b\x08\x79\x77\xd1\xdb\x62\xab\x81\x81\xef\xf7\x1f\x03\x2b\x99\xd1\x8f\xe1\xfc\x19\x7d\x49\xe8\xd2\x70\x13\x85\xd2\xca\xd5\x48\xb9\xac\xe8\xa2\x0d\x74\xf7\x9c\xb1\xb1\x45\x5c\x1b\xe7\x85\xb4\x06\x29\x4f\x67\xcc\x79\x3a\x1d\x9d\x8f\x4c\xbe\x12\x42\xe6\x90\x20\x05\x16\x09\x99\x4b\xc2\x95\xb2\x6c\x30\x41\x90\x72\x63\xbd\xff\xb5\x74\x77\xfd\xb2\x28\xd5\xfa\xd0\x80\x68\xf6\x5d\x9f\x9d\x3b\x94\x2e\x2c\xae\x55\x52\xf7\x3f\x70\xbb\xf1\xba\x5f\x6f\x2c\xe5\x2b\x51\x7f\x19\x03\x97\xb7\xc3\xf3\x9d\x3a\xe6\xfc\x0d\x00\x00\xff\xff\x0c\xc7\x10\x45\xe9\x00\x00\x00")

func templatesEnEmlConfirm_registrationEmlBytes() ([]byte, error) {
	return bindataRead(
		_templatesEnEmlConfirm_registrationEml,
		"templates/en/eml/confirm_registration.eml",
	)
}

func templatesEnEmlConfirm_registrationEml() (*asset, error) {
	bytes, err := templatesEnEmlConfirm_registrationEmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/en/eml/confirm_registration.eml", size: 233, mode: os.FileMode(420), modTime: time.Unix(1477537231, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesEnEmlProposalAcceptedEml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xb4\x53\xc1\x4e\x1b\x31\x10\xbd\xfb\x2b\x46\x39\xa2\xb0\x09\xa8\xed\x21\x57\x02\x22\x6a\x49\x2a\x12\x84\x7a\x9c\x78\x27\xbb\x2e\x5e\xcf\xe2\x99\x4d\x94\xa2\xfc\x7b\x65\x6f\x48\x0b\xa8\xc7\x5e\xfd\xc6\x6f\x9e\xde\x7b\x33\x25\x8c\xf0\xf2\x02\xc5\xb2\x25\x7c\xa2\x58\xcc\x9d\x7d\x0a\xd8\x10\x1c\x0e\xc6\xfc\xe0\x2e\x42\x1b\xb9\x65\x41\x0f\x83\x34\xb7\x72\xea\x13\x38\x80\x1a\x05\xd6\x44\x01\xd0\x5a\x6a\x95\xca\xc2\x98\xc1\x2d\xef\x68\x4b\x71\x08\x67\x67\x67\xb0\x23\x88\xf4\xdc\xb9\x48\xb0\xe9\xa2\xd6\x14\x01\xad\x3a\x0e\x19\xe5\x00\xfb\xc4\xbf\xa6\x1a\xfd\x06\x94\xc1\x72\xd3\x7a\x52\x02\xad\xc9\x9c\xd6\xb6\x91\x2d\x89\x14\xf0\xdd\x13\x0a\xc1\xd6\x89\xd3\x34\x02\x1b\xf6\x9e\x77\x2e\x54\xd0\x62\x45\x43\xc0\x50\xc2\xd5\x62\x7e\x33\xbb\xbf\x03\xad\x51\x13\x3f\x60\x24\x83\x6b\x4f\x69\x01\xaa\x52\x28\xf3\x5c\x1b\x49\x28\x68\x2f\x41\x48\x24\xc9\xc2\x9e\x57\x6c\x4d\x65\x97\xbe\x24\x92\x1d\x41\x8d\xdb\x24\x28\xff\xa0\x32\x31\xed\xb9\x9b\x18\x03\x50\xab\xb6\x32\x19\x8d\xd6\x9d\xf3\x25\x45\xb1\x1c\x0a\xc7\xa3\xe4\xd4\x15\x87\x0d\x45\x0a\x96\x8a\x25\x45\x47\x52\x2c\x7d\x57\xc1\xe1\xf0\x01\x3d\x3e\x1f\x55\x8c\x2c\x87\x8d\x8b\x4d\xb6\x1f\xee\x1e\x96\x2b\x38\xbe\xbc\x15\xeb\x02\x70\x2c\x29\xbe\x77\x0e\x3e\x38\x67\x66\x9b\x6c\xc5\x74\x01\xf3\xc5\x3b\xb6\xde\x12\x0c\x96\xe0\x71\xb6\xba\x9d\xcd\x61\x31\xbf\x86\xc7\xeb\xeb\xaf\xc3\x3e\x3f\xa1\xb8\xed\x79\xa3\xab\x6a\x35\xca\x50\x3a\x79\xee\xd0\xbb\xcd\xbe\xe7\x38\x2d\xdc\x39\xad\xb9\x53\xd8\x61\x0c\x2e\x54\x7d\x22\x42\x9e\xac\x02\x06\xce\x0d\xe0\x90\xa3\x50\x7c\x22\xe3\x54\xa0\xf5\x68\xa9\x30\x66\x59\x73\xe7\xcb\xac\x33\xd9\x0d\x18\xf6\x59\xa9\x77\x56\x25\x33\xff\x3b\x1b\x38\x65\x33\x84\x36\xb7\xc4\xfc\x69\x89\x3d\x39\x9d\x6b\x92\x35\x59\x0e\x8a\xb6\xc7\x39\x56\x18\xdc\x2f\x8a\x02\x28\x20\x9c\x7a\x20\xd0\xb2\x88\x5b\x7b\x1a\x82\xb0\x79\xdd\xd5\xe0\x3e\x1b\xd6\xb4\x9a\xeb\x54\xfe\xec\xe4\x6d\x65\x8a\x54\x8a\x87\xfb\x6f\x93\xff\xd2\x8c\x04\xcf\xa6\xe9\x34\x01\xf2\x21\x4e\xe0\xef\x9b\x34\x00\x53\x54\x1a\xad\x5c\x73\x44\x96\x8a\x51\x65\x11\x8a\x1b\x8e\x0d\x2a\x0c\x2e\xc7\xe3\x2f\xe7\xe3\x8b\xf3\xf1\x25\x5c\x7c\x9e\x8c\x3f\x0d\xf2\x9d\xdf\x53\x85\xb1\x94\xa1\x79\xb7\xff\x95\x18\x56\x84\xcd\xef\x00\x00\x00\xff\xff\xdb\x9f\x0c\xac\x2b\x04\x00\x00")

func templatesEnEmlProposalAcceptedEmlBytes() ([]byte, error) {
	return bindataRead(
		_templatesEnEmlProposalAcceptedEml,
		"templates/en/eml/proposal-accepted.eml",
	)
}

func templatesEnEmlProposalAcceptedEml() (*asset, error) {
	bytes, err := templatesEnEmlProposalAcceptedEmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/en/eml/proposal-accepted.eml", size: 1067, mode: os.FileMode(420), modTime: time.Unix(1477537231, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesJaEmlProposalAcceptedEml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xb4\x53\xcb\x4e\xdb\x5a\x14\x9d\xfb\x2b\x8e\x32\x44\xca\x03\x74\xef\x1d\x64\x7a\xd1\x95\xae\x54\xb5\x52\x43\x3f\x20\x4d\x4d\x6b\x41\x1c\x64\x9b\x11\x8a\xe4\x7d\xcc\x23\xc4\xa6\xa1\x90\x00\xe1\xd1\x24\x85\xf0\x88\x69\x52\x54\x8a\x42\xd2\x36\x1f\xb3\x7d\x1c\xe7\x2f\xaa\xe3\x38\x25\x40\x3b\xec\xec\xf8\x1c\xef\xb5\xd6\x5e\x6b\xef\xa5\x25\x12\x49\x2c\x88\xc9\x39\x51\x89\x3c\x95\x52\x73\x72\x32\x2d\x92\x6c\xd6\x3d\xdb\x13\x04\xd6\x3b\x64\x66\x05\x61\x19\xa1\x82\x50\xf5\x0f\xa7\xc1\xa7\xb1\x8b\xc6\x47\x34\x8e\xd0\xf8\x8a\xf4\x06\x0d\x9b\x84\x38\xd4\x8c\xa4\xcd\xf3\xfa\x10\x41\xb0\xdc\xb7\xb5\x7e\xf1\x1c\xa1\x84\xd4\xf2\x21\xb6\x11\xce\x91\x6e\x21\xe4\xfb\x95\x3a\xd2\x75\x84\x83\x11\xfa\x2e\xc2\x77\x84\x32\xea\x54\x10\x58\xeb\xc2\xad\x74\x03\x0a\xda\x45\x7a\x8b\x74\x8b\x35\x2d\xa7\xb3\xca\x7f\xa1\x26\x82\x8d\xd0\x22\x13\x13\x13\xa4\x5f\xee\xb2\xe3\xb5\xfe\x87\x8e\xd7\xd8\x40\xba\xe5\xd5\x2c\x84\xe3\x3b\x99\x5c\x75\x81\xf5\x56\xbc\x53\x40\xb0\x10\x28\xd2\xfc\x90\x89\x57\x73\x36\xd7\x5a\x63\xcd\xfd\x9f\x22\xdd\xdd\xba\x5b\xa6\x08\xb6\x7b\x78\x89\xf4\x04\x8d\x32\x1a\x9f\xd1\xc8\xf9\x94\x45\xf7\xa8\xce\xaa\xd7\x63\xd8\x1b\x43\xfe\xb1\x9b\x77\x48\x4d\x56\xbd\x66\x9b\x39\x84\x96\xd3\x36\xbd\xf3\x3d\x84\xa6\x80\x46\x09\xe9\x05\x77\xcb\xa8\x22\xcd\x21\xcd\xff\x5e\x79\xc1\x07\x2a\x71\x50\x6e\x07\x21\x6f\x34\x6d\x41\x8d\x47\xa3\x2f\x17\xa5\xf9\x57\xa2\xa2\xa6\x32\x72\x44\xca\x44\xb9\xe1\xff\x66\xe4\x59\x51\x11\xe5\x94\x18\x49\x88\x8a\x24\xaa\x91\xc4\xfc\xe2\x6b\x92\xcd\x3e\x7a\x0d\xae\x55\x51\x55\xa5\x8c\x1c\x4d\x65\xe4\x59\x49\x49\x0b\xc2\x3d\x1d\x60\x4d\x0e\xf4\xab\xc1\xce\xb6\xd3\xad\xb3\xd5\x15\xbf\xeb\xc6\xc8\x9c\x06\xc2\x72\xd0\x9a\x0e\x8f\x27\x00\xa1\x39\x4a\xdc\x62\x85\x1d\xa4\x79\xf7\x26\x17\xd4\x52\x93\x15\x5a\x9e\xf1\xcd\xd5\xcf\x10\x2c\x61\x3c\x08\x84\x26\xc2\x19\xbf\xe1\x43\x60\xfa\x63\xb1\x8f\x50\x74\x3a\xab\xee\x7a\xcf\x69\x9b\x63\x56\x38\xdd\x3c\x3b\x78\xef\x16\x36\xfb\x27\x1d\xff\xfa\xe0\x7e\xd4\xfe\x81\x8f\xcb\x15\xd2\x36\x1a\x75\xdf\x6f\x1b\xc1\x66\xa5\xc2\x80\x1b\x3c\x1c\x81\xbb\x80\x06\x7a\x05\xe9\x32\xa7\x05\x1b\xa9\xed\x67\x5d\x42\x5a\x43\xe3\x92\x9f\xe9\xad\xe0\xb4\xbb\x8c\x5e\x7a\xfa\xd0\x8b\xe2\x40\x3f\xee\x7f\xa9\x3d\xc8\x68\xd8\x5b\xd0\x05\x77\xea\xd3\xaf\x44\x34\xbd\x46\xcf\x2d\x5d\xfb\x0b\xf0\x58\xfa\xc6\xd8\x02\x10\xf2\xe2\xf9\x93\xf8\x1f\x89\x9d\x3f\xff\x3f\x4d\xb2\x59\x81\x10\x7f\x59\xe3\x64\x7c\x6f\x05\x42\xa6\x93\x9a\x18\x9d\x91\xd2\xc1\x4b\x42\x4b\x2a\x9a\xfa\x4c\x8e\xfc\x97\x51\xd2\x49\x8d\x84\xa6\x62\xb1\x7f\xc2\xb1\xc9\x70\x6c\x8a\x4c\xfe\x1d\x8f\xfd\x15\xe2\x65\x42\x38\x1c\x16\x1e\x50\x8f\x30\xc9\xd0\xc2\x1f\x01\x00\x00\xff\xff\x3e\x56\xb1\x50\x6a\x04\x00\x00")

func templatesJaEmlProposalAcceptedEmlBytes() ([]byte, error) {
	return bindataRead(
		_templatesJaEmlProposalAcceptedEml,
		"templates/ja/eml/proposal-accepted.eml",
	)
}

func templatesJaEmlProposalAcceptedEml() (*asset, error) {
	bytes, err := templatesJaEmlProposalAcceptedEmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/ja/eml/proposal-accepted.eml", size: 1130, mode: os.FileMode(420), modTime: time.Unix(1477537231, 0)}
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
	"templates/en/eml/confirm_registration.eml": templatesEnEmlConfirm_registrationEml,
	"templates/en/eml/proposal-accepted.eml": templatesEnEmlProposalAcceptedEml,
	"templates/ja/eml/proposal-accepted.eml": templatesJaEmlProposalAcceptedEml,
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
	"templates": {nil, map[string]*bintree{
		"en": {nil, map[string]*bintree{
			"eml": {nil, map[string]*bintree{
				"confirm_registration.eml": {templatesEnEmlConfirm_registrationEml, map[string]*bintree{}},
				"proposal-accepted.eml": {templatesEnEmlProposalAcceptedEml, map[string]*bintree{}},
			}},
		}},
		"ja": {nil, map[string]*bintree{
			"eml": {nil, map[string]*bintree{
				"proposal-accepted.eml": {templatesJaEmlProposalAcceptedEml, map[string]*bintree{}},
			}},
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
