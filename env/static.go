package env

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

var _conf_env_conf = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x9c\x54\xdf\x8f\xe3\x34\x10\x7e\xcf\x5f\x31\x4a\x5f\x40\xea\xa5\x7b\x87\xc4\x41\x25\x1e\x56\x48\x88\x15\xcb\x71\xa2\x2b\x9d\x50\xb5\x0f\x6e\x32\x49\xac\x3a\x76\xb0\xc7\x0d\xfd\xef\x99\xb1\x9b\x74\xcb\x2d\x3f\x44\xfb\x10\x7b\xfc\xcd\xe7\xf9\xf1\x8d\x57\xb0\xd3\x9d\x55\x46\xdb\x0e\x02\xfa\x13\x7a\x40\x7b\xd2\xde\xd9\x01\x2d\x15\xab\x62\x05\x3f\x2b\x6d\x41\x8d\xa3\xd1\xb5\x22\xed\xec\x8c\xab\x9d\x6d\x75\x17\x7d\x32\x56\x09\xda\xbb\x40\x5b\x78\xd4\x81\xd0\xa6\x4d\x05\x3b\x24\x78\xf8\x08\xf7\x4d\xe3\x31\x04\x70\x1e\x1a\x37\x08\xa3\x55\x03\xb2\xcb\xe8\xfc\xd5\x45\x36\x15\xfc\xe6\x22\x0c\xea\x0c\x16\xb1\x81\x40\xca\x13\x4c\x9a\x7a\x28\x43\x6c\x5c\xb9\x06\xdd\xc2\x99\x21\x81\x99\x27\x34\xe6\xcd\xd1\xba\xe9\xe2\xcb\x84\x68\x9b\xd1\x69\x7b\x25\x9d\x0d\x30\x2a\xea\x05\x41\x26\x2c\x87\x4f\x8f\xbb\x1b\xc2\x92\x7c\xc4\x52\x50\x35\x7a\x6a\xb5\xc1\x2d\x7c\xcf\x2b\x90\x65\x66\x80\x87\x0c\x8f\x01\x85\x6a\x0d\xd4\xeb\xc0\xce\x75\x2a\x4e\x0a\xdc\x11\xe0\x30\xd2\x99\x69\x8e\x78\xce\x2c\x3f\xe1\xf9\xff\x92\x14\xfb\x5c\xf2\xe7\x42\x6a\x0a\xdf\x41\xf9\xf6\xdd\xfb\xea\x8e\xff\x6f\xcb\x42\xf2\x66\xd3\x37\xfc\x2b\x96\x54\x19\xb2\x09\x73\x63\xcb\x82\x6f\x60\x53\xab\x4c\xc0\x62\xce\x4b\x30\x65\x71\x09\x2f\x6f\x38\xdc\x4f\x78\xf8\xd1\xb9\xe3\xdf\x69\xe1\xfe\x85\x0c\x7a\x75\xc2\x05\x9f\xae\x5d\x43\x32\xdb\x26\xf4\xea\x88\x52\xc4\xa7\x1e\x3d\xb6\xce\xe3\xfa\x46\x41\xe4\x95\x0d\x2d\xf3\x07\x37\x70\x41\xbc\xab\x45\x1b\x5f\x60\xd5\x55\xa0\x22\xf5\x7c\x9d\x40\xf1\x4b\x20\x27\x65\xf2\x72\xcf\x4b\x86\x2c\xb7\x5c\x43\x7f\x43\xed\x91\xa2\xb7\x5c\xca\x58\x27\x52\xd6\x0f\xc5\xc0\x62\x6d\x10\xde\xdd\xdd\xad\x59\x96\x1c\xae\x5e\x00\xd5\xbf\xf0\xb4\x4a\x9b\xe8\xf1\x25\xcf\x1a\xf0\x0f\x1c\xe9\x96\x4e\x70\xd8\xe4\xb8\xa2\x37\x5b\xf8\x15\x7f\x8f\xc8\xdd\xe2\x4d\x75\x53\xb7\x49\x1b\xc3\xf5\xb5\x4d\xf6\xf5\x17\x5c\x12\x00\x83\xb3\xd6\x47\xe5\x79\x3e\x08\x7d\xc8\x94\x33\x5b\xf2\x4b\x88\x8f\xbf\xec\x9e\x80\x21\xbd\x4b\xb7\xee\x27\x3c\x08\xdf\x73\xb1\x12\x12\xee\x67\x4f\x34\x6e\x37\x1b\xe3\x6a\x65\x44\x37\x1b\x39\x4e\x4d\x7e\x74\x5d\x97\x07\x9e\x88\xbf\x9f\x75\xd6\x45\x1a\x23\x05\x30\xae\x0b\xc0\xa3\x6a\x94\xef\x10\x6c\x1c\x0e\x39\x9e\x34\xd4\xe6\x42\x62\xf0\x84\x06\xb8\xe9\x62\x01\x3a\x8f\x5c\x1f\x59\x05\x91\x87\x68\x4b\xf4\xce\xd9\x29\x92\x32\xa1\x55\x07\x96\x1b\xb7\x75\xf2\x9a\xb8\xf7\xe8\x07\x1d\xc2\xd2\x52\xf1\xdf\x2e\x01\xca\xae\x4a\x73\x52\x2b\x9b\xe6\xf3\xa4\x0c\xd7\x01\x0e\x68\xdc\xb4\x65\x7c\xfe\xbd\xe1\xf6\x34\x1c\x35\x24\xd7\x25\x7e\xbe\x25\xdb\xd7\xb9\xbf\xfc\x62\x05\x67\x92\x30\x67\xbf\x24\x7e\xf8\xdc\x4f\xec\xe9\x61\x48\x03\x2a\xad\xc9\x99\x49\x2c\x43\x4c\x8d\xe0\xc7\x62\x4e\xaf\x9c\x27\x57\xa8\x67\x63\xe6\x5c\xca\x70\x99\xfb\x0f\xf2\xa8\xbd\x92\x7b\xae\xe3\x35\xf5\xb4\xfd\x6f\xb9\x2b\xd6\x93\xa4\x70\xcf\x5f\x69\xd9\x4d\x7e\x8a\x94\xe1\xb3\x1f\xd2\x37\xda\x86\x87\xee\xaf\x98\x49\x79\x2b\xfe\x9f\xf8\x2b\x57\xbf\x8e\xd2\xb6\x75\x82\x7a\xe0\xaf\x1f\xb2\x50\x5e\x47\x5a\x67\x51\x90\x1f\xdc\x7c\x54\xec\x79\xf1\x5c\x48\x09\x45\x99\xb9\x29\x65\xb1\xa8\x43\x9e\x2b\x1a\xc6\xeb\x93\x55\x31\xbe\x2c\xb2\xb4\xf8\x50\xee\x66\xe5\xee\x03\x39\xaf\x3a\xbc\x32\x79\x6c\x74\x90\x93\xb4\xf8\xa7\xf7\xf1\xeb\xaf\xde\x7f\x5b\xfc\x19\x00\x00\xff\xff\x46\x83\x15\x8b\xee\x06\x00\x00")

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

	info := bindata_file_info{name: "conf/env.conf", size: 1774, mode: os.FileMode(420), modTime: time.Unix(1423947212, 0)}
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

