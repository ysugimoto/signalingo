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

var _conf_env_conf = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x9c\x54\x4d\x8b\xe4\x36\x10\xbd\xeb\x57\x14\x9e\x4b\x02\xbd\xde\xd9\xbd\x24\x34\xcc\x61\x08\x84\x0c\x99\x2c\x21\xb3\xb0\x84\x61\x0f\x6a\xbb\x6c\x8b\x96\x25\x47\x2a\xb5\xd3\xff\x3e\x55\xd2\xd8\xed\xde\x9d\x7c\x90\xee\x83\xac\xd2\xab\x57\x1f\x7a\xa5\x1b\x78\x32\xbd\xd3\xd6\xb8\x1e\x22\x86\x13\x06\x40\x77\x32\xc1\xbb\x11\x1d\xa9\x1b\x75\x03\xbf\x68\xe3\x40\x4f\x93\x35\x8d\x26\xe3\xdd\x82\x6b\xbc\xeb\x4c\x9f\x42\x36\xd6\x19\x3a\xf8\x48\x7b\x78\x34\x91\xd0\xe5\x4d\x0d\x4f\x48\xf0\xf0\x2b\xdc\xb7\x6d\xc0\x18\xc1\x07\x68\xfd\x28\x8c\x4e\x8f\xc8\x2e\x93\x0f\x17\x17\xd9\xd4\xf0\xbb\x4f\x30\xea\x33\x38\xc4\x16\x22\xe9\x40\x30\x1b\x1a\xa0\x8a\xa9\xf5\xd5\x0e\x4c\x07\x67\x86\x44\x66\x9e\xd1\xda\x37\x47\xe7\xe7\x17\x5f\x26\x44\xd7\x4e\xde\xb8\x0b\xe9\x62\x80\x49\xd3\x20\x08\xb2\x71\x3d\xfc\xf8\xf8\x74\x45\x58\x51\x48\x58\x09\xaa\xc1\x40\x9d\xb1\xb8\x87\x1f\xf8\x0b\xe4\xb3\x30\xc0\x43\x81\xa7\x88\x42\xb5\x03\x1a\x4c\x64\xe7\x26\x37\x27\x27\xee\x09\x70\x9c\xe8\xcc\x34\x47\x3c\x17\x96\x9f\xf1\xfc\x7f\x49\xd4\x73\x69\xf9\x67\x25\x3d\x85\x3b\xa8\xde\xbd\xff\xae\xbe\xe5\xff\xbb\x4a\x49\xdd\x6c\xfa\x9e\x7f\x6a\x2d\x95\x21\x6f\xe3\x72\xb1\x95\xe2\x08\x6c\xea\xb4\x8d\xa8\x96\xba\x04\x53\xa9\x97\xf4\xca\x86\xd3\xfd\x84\x87\x9f\xbc\x3f\xfe\x9d\x16\xee\x37\x32\x18\xf4\x09\x21\xfa\x11\x57\xa7\x1c\x7b\xc7\x07\xae\x8d\x83\x3e\xe2\x0e\x7c\xd7\x31\xd3\x0e\xb4\x8b\x33\xaf\x75\x2d\x9d\xfd\x38\x60\xc0\xce\x07\x3e\xdf\xca\x8a\x02\xa3\x3a\x0e\x9a\x39\xa7\xe0\x1b\x11\xcc\x37\x58\xf7\x35\xe8\x44\x03\xe7\x20\x50\xfc\x16\xc8\x4b\xef\x82\xc4\xdd\x32\x14\x0d\x96\xc6\x86\x2b\xea\x80\x94\x82\xe3\xfe\xa6\x26\x93\xb2\xa8\x28\x45\x56\x70\x8b\xf0\xfe\xf6\x96\x53\x96\xf4\xcd\x0a\xa8\xff\x85\xa7\xd3\xc6\xa6\x80\x5b\x9e\x1d\xe0\x9f\x38\xd1\x35\x9d\xe0\xb0\xdd\xce\xc6\x6f\xf8\x47\x42\xbe\xc3\xcb\x70\x7c\x15\xa4\xcc\xc7\x3a\x1a\x8b\xc7\xa2\x6f\x91\xcf\xc6\xba\xd5\xb4\xd0\x89\x7e\x17\x45\x5f\xf3\xce\x3e\x1c\x65\xcc\xab\x81\x68\x8a\x95\x74\x98\x7c\xe3\xad\xa4\xf7\x3c\xe3\x41\x92\xfe\x47\x85\xdd\x2a\x89\x96\xb5\x25\xd8\x7a\x1a\xa6\x6b\x69\x71\x1e\x8f\xbe\xef\xcb\x5b\x42\xc4\xeb\x57\xa2\xf1\x89\xa6\x44\x11\xac\xef\x23\xf0\x2b\x60\x75\xe8\x11\x5c\x1a\x0f\x18\x72\xdb\xa5\x06\xfb\x42\x62\xf1\x84\x96\xb5\xd3\x8a\x05\xe8\x3c\x71\x65\xf2\x15\x45\x79\x22\xdb\x9c\x0f\x0d\x9a\xa4\xd9\xe8\xf4\x81\x95\xcc\xe2\x98\x83\x21\x56\x10\x86\xd1\xc4\xb8\x0a\x43\xfc\xf7\x6b\x82\xb2\xab\xf3\x08\x36\xda\xe5\xd1\x3f\x69\xcb\x2d\x85\x03\x5a\x3f\xef\x19\x5f\x7e\x6f\xf8\x92\x5b\xce\x1a\xf6\x39\xf4\x92\x7f\x31\xee\x4a\x97\xf9\x25\x8c\xde\x62\xbd\x71\xca\x43\x05\x5f\x38\x71\x6a\x62\xcf\xd7\x93\x07\x5f\x66\xbe\x94\x25\x89\x8c\x89\x5b\x9f\x1f\xa1\xa5\xb6\x6a\x79\x11\x84\x7a\x31\xe6\x1a\x2e\x3d\x78\x79\x4f\x3e\xc8\x63\xf9\x4a\xe1\xa5\x89\x97\xba\xf3\xf6\xbf\x15\xae\xad\xcd\x25\xdc\xf3\x2a\xf7\x75\x55\x9f\x26\x6d\xf9\xec\xc7\xbc\x26\xd7\xf2\xdc\x7e\x89\x99\x75\x70\xe2\xff\x89\x57\x09\xfd\x3a\xca\xb8\xce\x0b\xea\x81\xd7\x30\x16\x95\xbc\x8e\x74\xde\xa1\x20\x3f\xf8\xe5\x48\x3d\xf3\xc7\x67\x25\x2d\x64\x11\x96\x3b\x51\xab\x30\xee\xe0\x2d\x8d\xd3\xe5\x1d\xac\x19\xac\x8a\xa6\xee\x72\x58\xa5\xfe\x0a\x00\x00\xff\xff\x0d\x29\xf2\xf1\xff\x06\x00\x00")

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

	info := bindata_file_info{name: "conf/env.conf", size: 1791, mode: os.FileMode(420), modTime: time.Unix(1423675216, 0)}
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

