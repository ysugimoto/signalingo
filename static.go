package signaling

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"strings"
	"os"
	"time"
	"io/ioutil"
	"path"
	"path/filepath"
)

func bindata_read(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	gz.Close()

	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	return buf.Bytes(), nil
}

type asset struct {
	bytes []byte
	info  os.FileInfo
}

type bindata_file_info struct {
	name string
	size int64
	mode os.FileMode
	modTime time.Time
}

func (fi bindata_file_info) Name() string {
	return fi.name
}
func (fi bindata_file_info) Size() int64 {
	return fi.size
}
func (fi bindata_file_info) Mode() os.FileMode {
	return fi.mode
}
func (fi bindata_file_info) ModTime() time.Time {
	return fi.modTime
}
func (fi bindata_file_info) IsDir() bool {
	return false
}
func (fi bindata_file_info) Sys() interface{} {
	return nil
}

var _conf_env_conf = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x9c\x54\x4d\x6f\xe3\x36\x10\xbd\xeb\x57\x0c\xe4\x4b\x0b\x78\xe5\xec\x5e\x5a\x18\xe8\x21\x28\x50\x34\x68\x9a\x06\x75\x80\xa0\x08\x72\xa0\xa5\x91\x44\x98\x22\x55\x72\x68\xd5\xff\xbe\x33\xa4\x25\xc7\x4d\x8a\x16\x6b\x1f\x28\x0e\xdf\x3c\xce\xc7\x1b\xae\x60\xa7\x3b\xab\x8c\xb6\x1d\x04\xf4\x47\xf4\x80\xf6\xa8\xbd\xb3\x03\x5a\x2a\x56\xc5\x0a\x7e\x55\xda\x82\x1a\x47\xa3\x6b\x45\xda\xd9\x19\x57\x3b\xdb\xea\x2e\xfa\x64\xac\x12\xb4\x77\x81\xb6\x70\xaf\x03\xa1\x4d\x9b\x0a\x76\x48\x70\xf7\x08\xb7\x4d\xe3\x31\x04\x70\x1e\x1a\x37\x08\xa3\x55\x03\xb2\xcb\xe8\xfc\xc5\x45\x36\x15\xfc\xe1\x22\x0c\xea\x04\x16\xb1\x81\x40\xca\x13\x4c\x9a\x7a\x28\x43\x6c\x5c\xb9\x06\xdd\xc2\x89\x21\x81\x99\x27\x34\xe6\xd3\xc1\xba\xe9\xec\xcb\x84\x68\x9b\xd1\x69\x7b\x21\x9d\x0d\x30\x2a\xea\x05\x41\x26\x2c\x87\x4f\xf7\xbb\x2b\xc2\x92\x7c\xc4\x52\x50\x35\x7a\x6a\xb5\xc1\x2d\xfc\xc8\x5f\x20\x9f\x99\x01\xee\x32\x3c\x06\x14\xaa\x35\x50\xaf\x03\x3b\xd7\xa9\x38\x29\x70\x47\x80\xc3\x48\x27\xa6\x39\xe0\x29\xb3\xfc\x82\xa7\xaf\x25\x29\x5e\x72\xc9\x5f\x0b\xa9\x29\xfc\x00\xe5\xe7\x2f\xdf\x55\x37\xfc\xff\x5c\x16\x92\x37\x9b\xbe\xe7\x5f\xb1\xa4\xca\x90\x4d\x98\x1b\x5b\x16\x7c\x03\x9b\x5a\x65\x02\x16\x73\x5e\x82\x29\x8b\x73\x78\x79\xc3\xe1\x3e\xe3\xfe\x67\xe7\x0e\xff\xa6\x85\xdb\x37\x32\xe8\xd5\x11\x17\x7c\xba\x76\x0d\xc9\x6c\x9b\xd0\xab\x03\x4a\x11\x9f\x7a\xf4\xd8\x3a\x8f\xeb\x2b\x05\x91\x57\x36\xb4\xcc\x1f\xdc\xc0\x05\xf1\xae\x16\x6d\x7c\x83\x55\x57\x81\x8a\xd4\xf3\x75\x02\xc5\x6f\x81\x9c\x94\xc9\xcb\x3d\x6f\x19\xb2\xdc\x72\x0d\xfd\x15\xb5\x47\x8a\xde\x72\x29\x63\x9d\x48\x59\x3f\x14\x03\x8b\xb5\x41\xf8\x72\x73\xb3\x66\x59\x72\xb8\x7a\x01\x54\xff\xc1\xd3\x2a\x6d\xa2\xc7\xb7\x3c\x6b\xc0\xbf\x70\xa4\x6b\x3a\xc1\x61\x93\xe3\x8a\xde\x6c\xe1\x77\xfc\x33\x22\x77\x8b\x37\xd5\x55\xdd\x26\x6d\x0c\xd7\xd7\x36\xd9\xd7\x9f\x71\x49\x00\x0c\xce\x5a\x1f\x95\xe7\xf9\x20\xf4\x21\x53\xce\x6c\xc9\x2f\x21\x1e\x7f\xdb\x3d\x01\x43\x7a\x97\x6e\x7d\x99\x70\x2f\x7c\xaf\xc5\x4a\x48\xb8\x9f\x3d\xd1\xb8\xdd\x6c\x8c\xab\x95\x11\xdd\x6c\xe4\x38\x35\xf9\xde\x75\x5d\x1e\x78\x22\x5e\xdf\x75\xd6\x45\x1a\x23\x05\x30\xae\x0b\xc0\xa3\x6a\x94\xef\x10\x6c\x1c\xf6\x39\x9e\x34\xd4\xe6\x4c\x62\xf0\x88\x06\xb8\xe9\x62\x01\x3a\x8d\x5c\x1f\xf9\x0a\x22\x0f\xd1\x96\xe8\x9d\xb3\x53\x24\x65\x42\xab\xf6\x2c\x37\x6e\xeb\xe4\x35\x71\xef\xd1\x0f\x3a\x84\xa5\xa5\xe2\xbf\x5d\x02\x94\x5d\x95\xe6\xa4\x56\x36\xcd\xe7\x51\x19\xae\x03\xec\xd1\xb8\x69\xcb\xf8\xfc\xfb\xc4\xed\x69\x38\x6a\x48\xae\x4b\xfc\x7c\x4b\xb6\xaf\x73\x7f\xf9\xc5\x0a\xce\x24\x61\xce\x7e\x49\xfc\xf0\xde\x4f\xec\xe9\x61\x48\x03\x2a\xad\xc9\x99\x49\x2c\x43\x4c\x8d\xe0\xc7\x62\x4e\xaf\x9c\x27\x57\xa8\x67\x63\xe6\x5c\xca\x70\x9e\xfb\x07\x79\xd4\x3e\xc8\x3d\xd7\xf1\x92\x7a\xda\xfe\xbf\xdc\x15\xeb\x49\x52\xb8\xe5\x55\x5a\x76\x95\x9f\x22\x65\xf8\xec\xa7\xb4\x46\xdb\xf0\xd0\xfd\x13\x33\x29\x6f\xc5\xff\x99\x57\xb9\xfa\x63\x94\xb6\xad\x13\xd4\x1d\xaf\x7e\xc8\x42\xf9\x18\x69\x9d\x45\x41\x3e\xb8\xf9\xa8\x78\xe1\x8f\xd7\x42\x4a\x28\xca\x94\x72\x94\xc5\xa2\x0d\x79\xac\x68\x18\x2f\x0f\x56\xc5\xe8\xb2\xc8\xc2\xe2\x43\xb9\x99\x75\xfb\x77\x00\x00\x00\xff\xff\xc9\x12\xc0\x50\xab\x06\x00\x00")

func conf_env_conf_bytes() ([]byte, error) {
	return bindata_read(
		_conf_env_conf,
		"conf/env.conf",
	)
}

func conf_env_conf() (*asset, error) {
	bytes, err := conf_env_conf_bytes()
	if err != nil {
		return nil, err
	}

	info := bindata_file_info{name: "conf/env.conf", size: 1707, mode: os.FileMode(420), modTime: time.Unix(1423759153, 0)}
	a := &asset{bytes: bytes, info:  info}
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
	"conf/env.conf": conf_env_conf,
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
	for name := range node.Children {
		rv = append(rv, name)
	}
	return rv, nil
}

type _bintree_t struct {
	Func func() (*asset, error)
	Children map[string]*_bintree_t
}
var _bintree = &_bintree_t{nil, map[string]*_bintree_t{
	"conf": &_bintree_t{nil, map[string]*_bintree_t{
		"env.conf": &_bintree_t{conf_env_conf, map[string]*_bintree_t{
		}},
	}},
}}

// Restore an asset under the given directory
func RestoreAsset(dir, name string) error {
        data, err := Asset(name)
        if err != nil {
                return err
        }
        info, err := AssetInfo(name)
        if err != nil {
                return err
        }
        err = os.MkdirAll(_filePath(dir, path.Dir(name)), os.FileMode(0755))
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

// Restore assets under the given directory recursively
func RestoreAssets(dir, name string) error {
        children, err := AssetDir(name)
        if err != nil { // File
                return RestoreAsset(dir, name)
        } else { // Dir
                for _, child := range children {
                        err = RestoreAssets(dir, path.Join(name, child))
                        if err != nil {
                                return err
                        }
                }
        }
        return nil
}

func _filePath(dir, name string) string {
        cannonicalName := strings.Replace(name, "\\", "/", -1)
        return filepath.Join(append([]string{dir}, strings.Split(cannonicalName, "/")...)...)
}

