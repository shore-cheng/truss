// Code generated by go-bindata. DO NOT EDIT.
// sources:
// NAME-service/grpc/NAME/client.gotemplate (3.184kB)
// NAME-service/http/NAME/client.gotemplate (105B)

package template

import (
	"bytes"
	"compress/gzip"
	"crypto/sha256"
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
		return nil, fmt.Errorf("read %q: %w", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("read %q: %w", name, err)
	}
	if clErr != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

type asset struct {
	bytes  []byte
	info   os.FileInfo
	digest [sha256.Size]byte
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

var _grpcNameClientGotemplate = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\x56\x4d\x6f\xdb\x38\x10\x3d\x8b\xbf\x62\xd6\x08\x16\x52\xa0\xd0\xf7\x2c\x7c\xa9\xd3\x2d\xba\xd8\xa6\x46\x1a\x74\x0f\x45\x51\x30\xd4\x58\x26\x2c\x93\x2a\x49\x3b\x31\x04\xfd\xf7\xc5\x90\x94\x23\x27\x8e\xdb\x43\x10\x8b\xf3\x38\x1f\xef\x0d\x39\x9c\x4e\x61\x6e\x2a\x84\x1a\x35\x5a\xe1\xb1\x82\x87\x3d\x78\xbb\x75\x8e\xc3\xcd\x67\xb8\xfd\x7c\x0f\xef\x6f\x3e\xde\x73\x36\x9d\xc2\x1d\xda\xad\xd6\x4a\xd7\x11\x00\x8f\xaa\x69\xc0\xec\xd0\x3e\x5a\xe5\x11\xfc\x4a\x39\x58\xaa\x06\x03\xf8\x2b\x5a\xa7\x8c\xbe\x86\xae\xe3\xe9\x77\xdf\x8f\x0c\x70\x23\x3c\x8e\xad\xf4\xdd\xf7\x8c\x20\x0b\x21\xd7\xa2\x46\xa8\x6d\x2b\xa1\xb5\x66\xa7\x2a\x74\x20\xa0\xbe\x5b\xcc\x41\x36\x0a\xb5\x87\xa5\xb1\xe0\x57\x48\x0e\xbe\xa0\xdd\x29\x89\xfc\x56\x6c\xb0\xef\xc1\xa5\x4f\xd6\x8e\xdc\x30\xa6\x36\xad\xb1\x1e\x72\x96\x4d\xa4\xd1\x1e\x9f\xfc\x84\x65\x93\xda\x98\xba\x41\x5e\x9b\x46\xe8\x9a\x1b\x5b\x4f\x09\xfd\xb6\x65\xba\x41\x2f\x2a\xe1\x45\x80\x28\xbf\xda\x3e\x70\x69\x36\xd3\x76\x5d\x4f\xd1\x5a\x63\xdd\x84\x1d\x5b\x6a\x73\xb5\x56\x7e\x4a\x7f\xa8\xab\xd6\x28\x4d\x81\xc9\x97\xb7\x42\xbb\x90\xd4\x1b\xf8\x03\x20\x25\xc5\xb2\xe9\x14\xee\x89\xe6\x54\x32\xcb\x26\x5d\xc7\x3f\x86\xca\x16\xc2\xaf\xe0\xaa\xef\x61\xea\x76\x54\x40\xfb\x00\x64\x5c\xbc\x3b\x36\x4f\x58\x11\x38\xbe\xc5\x47\xb0\xe8\xb7\x56\x3b\x10\x7a\x20\x0d\x1e\x84\x5c\xc7\x26\x38\xa6\x5b\x1a\xad\x51\x7a\x65\x34\x87\x8f\x1e\x94\x23\xf2\xc9\x8f\x45\xd7\x1a\xed\xd4\x83\x6a\x94\xdf\x83\x59\x06\x55\xa4\x68\x1a\xb4\xe0\x0d\x54\x4a\x34\x25\x08\x5d\x41\x23\x3c\x5a\x90\x8d\x71\x58\x46\xd0\xb3\x4f\xb6\xdc\x6a\x49\x39\xe5\xb4\x08\x97\x54\x2f\x9f\x87\xd0\x73\xa3\x75\x09\xa6\x25\x9c\x03\xce\xd3\xf2\xe7\xb0\x50\x40\xde\x3e\xf0\x57\x3d\x40\x5f\x68\x4b\x08\x8a\x14\xd0\xb1\x6c\x27\x2c\x48\x99\xaa\x99\x1b\xbd\x54\x35\x63\x19\x35\xd1\x8f\x12\x96\x70\x3d\x03\x2b\x74\x8d\x87\x38\x1d\xcb\x32\xb4\x96\x0c\xcb\xfc\x4f\x29\x0b\x96\x65\x6a\x49\x0e\xe1\x8f\x19\x68\xd5\x04\x44\x16\x19\xa4\xef\x14\xcc\xf1\xff\xac\x68\x73\xb4\xb6\x84\x89\x14\x5a\x1b\x0f\xa2\x6d\x9b\x7d\xf2\x3c\x21\x47\x3d\xcb\x7a\xc6\x32\x39\x2a\xc4\x51\xa4\x6f\xdf\x8f\xda\xe2\xa8\x52\x0a\x77\xca\xfa\x0e\x97\xc6\x62\x4e\xc9\xa4\xb6\xfe\x2a\x9a\x2d\xba\x7b\xf3\xe1\x6e\x31\xff\x94\xba\x35\x97\x92\xaf\x50\x54\x68\x5d\x51\x94\x14\x3e\xeb\xba\x2b\x78\x54\x7e\x05\x17\x1e\x29\x38\xef\x7b\x96\x8d\x56\xdb\x75\x4d\x64\x92\xe9\xc2\x23\x4f\x67\x32\xf2\x4b\xd1\x08\x19\x39\xbb\x50\x03\x68\x50\xe1\x13\xfa\x95\xa9\x5c\x04\x06\xee\xbb\xee\xde\xfc\x6b\x1e\xd1\xc2\x85\x4a\x22\xbd\x4f\xa7\x01\x86\x63\xc1\x87\x95\xb0\x2b\xf0\x4b\x61\xde\xde\x38\x83\x63\x46\x6e\xf1\x31\x92\x92\xc7\xbd\xc4\x88\x2e\xd3\xef\x49\xd7\x0d\x35\xf5\x3d\xef\xba\x71\xbe\x71\x71\x32\x86\xaa\x97\x8b\xef\xb5\x34\x15\x12\xa9\x23\xeb\x1d\xfe\xdc\xa2\xf3\x03\xe6\x06\x4f\x62\xc2\x09\xc1\x01\x14\x1a\xf6\x83\x09\xe4\x5e\x28\x3e\x98\xef\xf7\xed\x90\x48\xd7\x0f\xd8\xa3\x16\xe1\x9c\xa7\xf5\xe2\x40\x55\x5e\x84\x95\xa4\x08\xea\x2a\xa9\x98\x7e\x0d\x3f\xd8\xd0\xa9\x6e\x27\x0f\x7b\x5d\x47\x80\xb1\x86\x2f\x05\xa4\x0b\x23\xb8\x7b\xc5\xfd\x35\x00\x9c\x13\xb5\x7c\x8e\x9d\xf5\x25\x1d\x10\x16\xef\x76\x22\x07\xa2\x4a\x10\xe9\x62\xe7\x73\x88\x53\xe3\x2c\xb3\x74\x1d\x09\x38\xbe\x2d\x79\xdc\x31\x40\xfe\xa6\xfb\xc5\xaf\x44\xb8\xc9\x76\x68\xbd\x03\x41\x7e\xc3\x1d\x77\xa2\x0e\xb0\x48\x87\xd6\x1b\x10\xb0\x75\x68\xaf\x2a\xb3\x11\x4a\xbf\x01\x8d\x31\x38\x2c\xac\xda\x08\xab\x9a\x3d\xed\x59\x6e\x1b\x50\x1a\x44\xba\x74\xd2\x1d\x77\xb6\x90\xfc\x07\xa4\x43\xcc\xe7\xf1\x7f\x19\x5a\xfc\x2e\x24\xa3\xb4\x47\xbb\x14\x12\xbb\xbe\x80\x7c\xf4\x35\xbe\xe8\x62\xde\xd7\xb3\xe7\x7d\x3c\xbf\xfc\x75\xcb\x15\x87\x0e\x09\x0e\x06\xc5\x0e\xfd\xf3\x42\xb9\x78\x18\x7e\x4b\xb9\x73\xe7\xe6\xa4\x70\x71\x43\x42\xbc\xa5\xdb\xaf\x35\x89\x01\x82\x80\x67\x44\x0e\xa8\xdf\x12\xee\x5c\x1d\xa7\x74\x1b\x32\xf8\x4d\xd5\x7e\x86\x19\x94\xf2\x39\xa1\x58\x30\xbc\x21\xd8\xcf\x57\x72\x31\xbf\x6f\xf1\x68\xda\x81\xf3\x76\x2b\x3d\x05\x4b\x83\x00\xbe\x7d\x77\xde\x2a\x5d\xa7\x93\x39\x9e\x36\x51\x18\xaa\x3b\x7c\x05\x01\x36\xa6\x52\x4b\x85\x2e\xce\xee\xc3\xb3\x80\x26\x69\x88\x76\xb4\x9f\xb6\xe6\x97\xe3\x04\x8a\x58\x2e\x8b\x6c\xce\xfd\xd3\x30\xa7\xbe\xa0\xae\xf2\x35\xee\xc3\x70\x8f\x19\x15\xc7\xce\xba\x43\xad\xc1\xad\x81\x53\x8e\xc3\x40\x36\xc3\x94\x83\x19\x90\x4b\x36\x1e\xd1\x34\xf6\xfa\x14\xff\xdc\xac\x0c\xb9\x0c\xe4\x14\x70\x6a\xea\x8e\xbb\xf3\x45\x76\xd2\x3f\xbd\x6e\x86\x4d\x05\x97\xc3\xcb\x91\x7f\xba\x29\x5e\x22\x42\xf2\x34\x27\x5b\xa1\xc6\xca\x64\xc3\x13\x65\xfd\xfc\x44\x09\xe9\x85\xe9\xa8\x96\xb0\x2b\xc1\x04\x9b\xf4\x4f\x3c\x54\x93\xaf\x0b\x9e\xa7\xdc\xff\x22\x63\x1c\xa4\xd1\xf1\x8c\x1e\x23\xc4\x77\xf8\x2c\x61\x5d\xc2\x2e\x4c\x90\x3e\x3c\x4b\xe2\x23\x27\x42\xc7\xcf\x9c\xcb\x4d\x05\x33\x38\x14\xf0\x8f\x51\x3a\xbf\xdc\x54\xe5\xf3\xd2\x82\xf6\x44\xaf\x9c\xf3\xa2\x18\xdc\x25\x66\xa4\x7f\x8a\xec\xff\x1f\x00\x00\xff\xff\x00\xce\x0e\xa6\x70\x0c\x00\x00")

func grpcNameClientGotemplateBytes() ([]byte, error) {
	return bindataRead(
		_grpcNameClientGotemplate,
		"grpc/NAME/client.gotemplate",
	)
}

func grpcNameClientGotemplate() (*asset, error) {
	bytes, err := grpcNameClientGotemplateBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "grpc/NAME/client.gotemplate", size: 3184, mode: os.FileMode(0644), modTime: time.Unix(1464111000, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x63, 0x72, 0x1f, 0xe5, 0x3a, 0x45, 0x1, 0x91, 0xd8, 0x5b, 0xa8, 0x47, 0x45, 0x45, 0x98, 0xee, 0x0, 0xf5, 0xc1, 0x3c, 0x43, 0xf0, 0x86, 0x3c, 0xec, 0xbe, 0x2d, 0x84, 0xed, 0x1a, 0x17, 0x6c}}
	return a, nil
}

var _httpNameClientGotemplate = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x2c\xcb\xb1\x0d\xc2\x30\x10\x05\xd0\x3e\x53\x5c\x1d\x09\xdf\x10\x34\x29\x91\xc8\x02\x56\xf8\x98\x88\xc3\x67\x9d\x3f\x95\xe5\xdd\x69\x18\xe0\x8d\xa1\xab\xdc\x01\x29\x7e\x61\x7c\x7b\xd7\x82\x5a\xfc\x7d\x52\x5f\x64\x63\xe4\xda\x9b\x07\x95\xf8\x34\xcb\x44\x2a\x2e\x4f\x0f\x39\xfc\x01\x59\x75\xce\x65\x8c\x23\x9b\x49\xda\xf6\xfd\xb6\xc1\x1a\x22\x5d\xed\x44\xe5\xfe\x27\x92\xe6\x5c\x7e\x01\x00\x00\xff\xff\x0b\x3c\x4c\x9e\x69\x00\x00\x00")

func httpNameClientGotemplateBytes() ([]byte, error) {
	return bindataRead(
		_httpNameClientGotemplate,
		"http/NAME/client.gotemplate",
	)
}

func httpNameClientGotemplate() (*asset, error) {
	bytes, err := httpNameClientGotemplateBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "http/NAME/client.gotemplate", size: 105, mode: os.FileMode(0644), modTime: time.Unix(1464111000, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xa1, 0xf0, 0x36, 0xf9, 0x16, 0xea, 0x9d, 0x4e, 0x73, 0x64, 0xc5, 0xad, 0xb3, 0x1b, 0x4, 0xe, 0xd8, 0xc8, 0x1e, 0xf7, 0x7a, 0x39, 0x40, 0x4c, 0xb2, 0x12, 0x83, 0x35, 0xca, 0x82, 0x6f, 0xd0}}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[canonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// AssetString returns the asset contents as a string (instead of a []byte).
func AssetString(name string) (string, error) {
	data, err := Asset(name)
	return string(data), err
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

// MustAssetString is like AssetString but panics when Asset would return an
// error. It simplifies safe initialization of global variables.
func MustAssetString(name string) string {
	return string(MustAsset(name))
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[canonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetDigest returns the digest of the file with the given name. It returns an
// error if the asset could not be found or the digest could not be loaded.
func AssetDigest(name string) ([sha256.Size]byte, error) {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[canonicalName]; ok {
		a, err := f()
		if err != nil {
			return [sha256.Size]byte{}, fmt.Errorf("AssetDigest %s can't read by error: %v", name, err)
		}
		return a.digest, nil
	}
	return [sha256.Size]byte{}, fmt.Errorf("AssetDigest %s not found", name)
}

// Digests returns a map of all known files and their checksums.
func Digests() (map[string][sha256.Size]byte, error) {
	mp := make(map[string][sha256.Size]byte, len(_bindata))
	for name := range _bindata {
		a, err := _bindata[name]()
		if err != nil {
			return nil, err
		}
		mp[name] = a.digest
	}
	return mp, nil
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
	"grpc/NAME/client.gotemplate": grpcNameClientGotemplate,
	"http/NAME/client.gotemplate": httpNameClientGotemplate,
}

// AssetDebug is true if the assets were built with the debug flag enabled.
const AssetDebug = false

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//     data/
//       foo.txt
//       img/
//         a.png
//         b.png
// then AssetDir("data") would return []string{"foo.txt", "img"},
// AssetDir("data/img") would return []string{"a.png", "b.png"},
// AssetDir("foo.txt") and AssetDir("notexist") would return an error, and
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		canonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(canonicalName, "/")
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
	"grpc": {nil, map[string]*bintree{
		"NAME": {nil, map[string]*bintree{
			"client.gotemplate": {grpcNameClientGotemplate, map[string]*bintree{}},
		}},
	}},
	"http": {nil, map[string]*bintree{
		"NAME": {nil, map[string]*bintree{
			"client.gotemplate": {httpNameClientGotemplate, map[string]*bintree{}},
		}},
	}},
}}

// RestoreAsset restores an asset under the given directory.
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
	return os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
}

// RestoreAssets restores an asset under the given directory recursively.
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
	canonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(canonicalName, "/")...)...)
}
