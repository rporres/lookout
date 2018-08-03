// Code generated by go-bindata.
// sources:
// store/migrations/1533123066_initial_schema.down.sql
// store/migrations/1533123066_initial_schema.up.sql
// store/migrations/1533132482_add_number_to_review.down.sql
// store/migrations/1533132482_add_number_to_review.up.sql
// store/migrations/1533135217_add_repo_id_to_review.down.sql
// store/migrations/1533135217_add_repo_id_to_review.up.sql
// store/migrations/lock.json
// DO NOT EDIT!

package store

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

var _migrations1533123066_initial_schemaDownSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x72\x72\x75\xf7\xf4\xb3\xe6\xe2\x72\x09\xf2\x0f\x50\x08\x71\x74\xf2\x71\x55\x28\x28\x2d\xce\x88\x4f\x2d\x4b\xcd\x2b\x41\x15\x4f\xce\xcf\xcd\xc5\x10\x2c\x4a\x2d\xcb\x4c\x2d\x87\x2b\x77\xf6\xf7\xf5\xf5\x0c\xb1\xe6\x02\x04\x00\x00\xff\xff\x67\xfb\x73\x7b\x57\x00\x00\x00")

func migrations1533123066_initial_schemaDownSqlBytes() ([]byte, error) {
	return bindataRead(
		_migrations1533123066_initial_schemaDownSql,
		"migrations/1533123066_initial_schema.down.sql",
	)
}

func migrations1533123066_initial_schemaDownSql() (*asset, error) {
	bytes, err := migrations1533123066_initial_schemaDownSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "migrations/1533123066_initial_schema.down.sql", size: 87, mode: os.FileMode(509), modTime: time.Unix(1533216138, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _migrations1533123066_initial_schemaUpSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xcc\x92\xcf\x6e\xf2\x30\x10\xc4\xcf\xf1\x53\xec\x11\xa4\xbc\x01\x27\x40\xfe\x3e\xa1\x42\xa8\xd2\xf4\xc0\x29\x72\xec\x25\x6c\x95\xd8\x91\xbd\xa6\x55\x9f\xbe\x02\xa9\x55\x1a\xf7\xdf\xa1\x87\x5e\x67\x7e\x7b\x98\x99\x5d\xc9\xff\x9b\x62\x21\xc4\xba\x94\xcb\x4a\x42\xb5\x5c\x6d\x25\x78\x3c\x13\x3e\xd6\x78\x46\xcb\x30\x13\x19\x19\x88\x91\x0c\x14\xfb\x0a\x8a\xfb\xed\x16\x6e\xcb\xcd\x6e\x59\x1e\xe0\x46\x1e\x72\x91\x05\x56\x1c\x03\x30\x3e\xf1\x1b\x92\x8b\x6c\xf0\xee\x4c\x06\x7d\x62\x90\x65\xf4\x56\x75\x35\x99\xc4\xd3\x1e\x15\xa3\xa9\x15\x03\x53\x8f\x81\x55\x3f\xf0\xf3\x98\x88\x83\xf9\x86\xa0\x50\xf7\xe8\x5b\x54\x4d\x87\xd0\x38\xd7\xa1\xb2\x63\x3f\xb8\xe8\x35\xc2\x43\x70\xb6\x19\xeb\xd7\xa3\x54\xd6\xce\x1e\xa9\x8d\x5e\x31\x39\x9b\xda\x8d\x0a\x1f\x1c\x9d\x50\x99\x89\x2a\xe6\x0b\x31\x29\x5a\xbb\xbe\xff\x59\xc7\xe3\x49\xea\x57\xb6\x94\xff\x64\x29\x8b\xb5\xbc\x7b\x37\xd9\x8c\xcc\x3c\x17\xd9\x91\x3a\x4c\xfa\xed\xc8\x22\x5c\x06\x68\xd1\x8f\xf5\x2b\x97\x8c\x71\x49\x6e\xd0\x6a\x84\x86\x5a\xb2\xfc\x55\x96\x21\x86\xd3\x1f\x7d\x99\x4b\xcd\xc4\x61\x1a\x22\x17\x99\xa1\xc0\x64\x35\xd7\x9f\x23\xbf\x3c\xff\x7e\xb7\xdb\x54\x0b\xf1\x12\x00\x00\xff\xff\xca\x98\xf7\xfa\x78\x03\x00\x00")

func migrations1533123066_initial_schemaUpSqlBytes() ([]byte, error) {
	return bindataRead(
		_migrations1533123066_initial_schemaUpSql,
		"migrations/1533123066_initial_schema.up.sql",
	)
}

func migrations1533123066_initial_schemaUpSql() (*asset, error) {
	bytes, err := migrations1533123066_initial_schemaUpSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "migrations/1533123066_initial_schema.up.sql", size: 888, mode: os.FileMode(509), modTime: time.Unix(1533216138, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _migrations1533132482_add_number_to_reviewDownSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x72\x72\x75\xf7\xf4\xb3\xe6\xe2\x72\xf4\x09\x71\x0d\x52\x08\x71\x74\xf2\x71\x55\x28\x4a\x2d\xcb\x4c\x2d\x8f\x4f\x2d\x4b\xcd\x2b\x51\x70\x09\xf2\x0f\x50\x70\xf6\xf7\x09\xf5\xf5\x53\xc8\x2b\xcd\x4d\x4a\x2d\xb2\xe6\xe2\x72\xf6\xf7\xf5\xf5\x0c\xb1\xe6\x02\x04\x00\x00\xff\xff\x1a\x1f\x67\xc9\x3e\x00\x00\x00")

func migrations1533132482_add_number_to_reviewDownSqlBytes() ([]byte, error) {
	return bindataRead(
		_migrations1533132482_add_number_to_reviewDownSql,
		"migrations/1533132482_add_number_to_review.down.sql",
	)
}

func migrations1533132482_add_number_to_reviewDownSql() (*asset, error) {
	bytes, err := migrations1533132482_add_number_to_reviewDownSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "migrations/1533132482_add_number_to_review.down.sql", size: 62, mode: os.FileMode(509), modTime: time.Unix(1533216138, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _migrations1533132482_add_number_to_reviewUpSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x04\xc0\x3b\x0a\x03\x20\x0c\x06\xe0\x3d\xa7\xf8\xef\xe1\xe4\x8b\x22\xc4\x08\x25\xce\x05\x21\x14\x87\x3a\x88\xb5\xd7\xef\x17\xf2\xa3\x88\x23\xf2\xac\xf9\x09\xf5\x81\x33\xb6\xdd\x69\xbf\x97\x5d\x5b\x07\x3e\x25\xc4\xc6\xbd\x0a\xd6\xf7\x33\x6c\x63\xcc\xf7\x5c\x07\xd2\x14\xd2\x99\x1d\x51\x6c\xb5\x16\x75\xf4\x0f\x00\x00\xff\xff\x0e\xda\x46\x8c\x4d\x00\x00\x00")

func migrations1533132482_add_number_to_reviewUpSqlBytes() ([]byte, error) {
	return bindataRead(
		_migrations1533132482_add_number_to_reviewUpSql,
		"migrations/1533132482_add_number_to_review.up.sql",
	)
}

func migrations1533132482_add_number_to_reviewUpSql() (*asset, error) {
	bytes, err := migrations1533132482_add_number_to_reviewUpSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "migrations/1533132482_add_number_to_review.up.sql", size: 77, mode: os.FileMode(509), modTime: time.Unix(1533216138, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _migrations1533135217_add_repo_id_to_reviewDownSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x72\x72\x75\xf7\xf4\xb3\xe6\xe2\x72\xf4\x09\x71\x0d\x52\x08\x71\x74\xf2\x71\x55\x28\x4a\x2d\xcb\x4c\x2d\x8f\x4f\x2d\x4b\xcd\x2b\x51\x70\x09\xf2\x0f\x50\x70\xf6\xf7\x09\xf5\xf5\x53\x28\x4a\x2d\xc8\x2f\xce\x2c\xc9\x2f\xaa\x8c\xcf\x4c\xb1\xe6\xe2\x72\xf6\xf7\xf5\xf5\x0c\xb1\xe6\x02\x04\x00\x00\xff\xff\x31\x5a\x11\xc5\x45\x00\x00\x00")

func migrations1533135217_add_repo_id_to_reviewDownSqlBytes() ([]byte, error) {
	return bindataRead(
		_migrations1533135217_add_repo_id_to_reviewDownSql,
		"migrations/1533135217_add_repo_id_to_review.down.sql",
	)
}

func migrations1533135217_add_repo_id_to_reviewDownSql() (*asset, error) {
	bytes, err := migrations1533135217_add_repo_id_to_reviewDownSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "migrations/1533135217_add_repo_id_to_review.down.sql", size: 69, mode: os.FileMode(509), modTime: time.Unix(1533216138, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _migrations1533135217_add_repo_id_to_reviewUpSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x04\xc0\xc1\x0a\xc3\x20\x0c\x06\xe0\x7b\x9e\xe2\x7f\x0f\x4f\xb6\x95\x51\x88\x11\x46\x7a\x16\xc6\xc2\xc8\x45\x87\x13\xc7\xde\x7e\xdf\x96\x6e\xa7\x04\xa2\xc8\x9a\xee\xd0\xb8\x71\xc2\xb0\xe5\xf6\xad\xb6\xac\x4d\xc4\xe3\xc0\x5e\xf8\xca\x82\x61\xef\xfe\xf1\xd9\xc7\xaf\xfa\x13\x0f\x7f\x79\x9b\x90\xa2\x90\x8b\x39\x10\xed\x25\xe7\x53\x03\xfd\x03\x00\x00\xff\xff\xab\x20\x0d\xc4\x54\x00\x00\x00")

func migrations1533135217_add_repo_id_to_reviewUpSqlBytes() ([]byte, error) {
	return bindataRead(
		_migrations1533135217_add_repo_id_to_reviewUpSql,
		"migrations/1533135217_add_repo_id_to_review.up.sql",
	)
}

func migrations1533135217_add_repo_id_to_reviewUpSql() (*asset, error) {
	bytes, err := migrations1533135217_add_repo_id_to_reviewUpSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "migrations/1533135217_add_repo_id_to_review.up.sql", size: 84, mode: os.FileMode(509), modTime: time.Unix(1533216138, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _migrationsLockJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x58\xc1\x8a\xdb\x30\x10\xbd\xfb\x2b\x84\xcf\xfb\x05\xb9\xf6\x58\x08\xa5\x6c\x4f\xa5\x18\xd9\x1e\x7b\xa7\x48\x23\x57\x1a\xa5\x75\x97\xfc\x7b\xb1\xd8\x4d\x6d\xc7\x6a\x4a\x4a\x37\x72\xf0\x25\x08\x0d\x93\x79\xef\x49\x9a\x37\xc9\x73\x26\x44\xfe\x28\x4b\x05\x2e\xdf\x89\xcf\x99\x10\x42\x3c\x87\x4f\x21\xf2\xbd\xd4\x90\xef\x44\x5e\x19\xad\x81\x38\x7f\x78\x0d\xbc\x33\xca\x6b\xfa\x9d\x31\xce\x9a\x64\x62\x7d\x4a\x0a\xfb\x8f\x7d\x17\xf6\xbd\x9f\x47\x3e\x58\xd4\xd2\xf6\xef\xa1\xcf\x77\x82\xad\x87\x49\xf4\x23\x34\x60\x81\xaa\x21\x99\xbc\x52\x93\xe0\xde\xf0\xde\x2b\xb5\x94\xf7\x89\xf0\x9b\x1f\x92\x1a\xa9\x1c\x9c\x22\xc7\x87\x3f\xc3\xb6\x70\x40\xf8\x5e\xc0\x01\x88\x8b\x2b\x39\x84\x8a\x51\x12\xe3\xb2\xaf\x27\x30\xaf\x3c\xf9\xf2\x93\xea\x2f\xb2\x8e\x42\xc7\x88\x1a\xe7\x08\xae\x94\xa3\x41\x05\xcb\x1a\x30\xfc\xe0\xab\x35\x78\x83\x83\x54\x48\x11\xe4\x48\x0c\x2d\xd8\x94\xc1\x9f\x8b\xbb\x16\xd9\x2b\x43\x0d\xd6\xa1\xde\x22\xfe\x12\x5b\xa4\xdb\x32\x78\x59\x7d\xc9\x46\x7c\xce\x1a\x5f\xe7\xdd\xd3\xec\x2d\xde\x7d\xef\x73\x2c\xd9\xbb\x75\xde\xbb\xce\x9a\x03\xd6\xf3\x57\xbd\x16\xf4\x43\x4b\xb2\x24\x55\xd4\x71\x52\x27\x50\x59\x90\x0c\x75\x21\x63\x6d\x0b\x35\x38\x96\xba\xe3\x9f\x49\xd3\x30\x5a\x23\x47\x9e\x40\x02\xad\xeb\x02\xfc\x1a\x1d\x23\x55\x5c\xac\x9c\x47\x30\x91\xd6\x5b\xc9\x68\x68\x99\xc4\x57\x67\xa8\x4c\x99\x43\x29\x5d\xc4\x02\x93\x87\xfe\x04\x32\xd2\x86\x6e\x0f\xfd\xaf\xcc\x7b\x71\x94\xde\xec\x3b\xe1\x2b\xb7\xd9\xf7\x66\xdf\xff\x4e\xc3\x77\xf5\x3d\xd0\x40\x57\x68\xb0\x2d\x84\x7f\x06\x96\x2d\xdc\x18\x05\x92\x52\x26\xe1\x8c\xb7\xb1\x1f\x81\xb7\xb7\x91\x0b\xe0\x83\xfc\x2b\xc5\x7e\x0f\xc3\x93\x85\xce\x38\x64\x63\xfb\x68\x4b\x4d\x7f\x8a\x25\xaf\xcb\x98\xa3\xa5\x8f\x7e\x9b\x5f\xff\x0b\xf4\xe9\xfc\x9a\x0d\xab\xe3\xaf\x00\x00\x00\xff\xff\x01\x5e\x33\x03\x8a\x17\x00\x00")

func migrationsLockJsonBytes() ([]byte, error) {
	return bindataRead(
		_migrationsLockJson,
		"migrations/lock.json",
	)
}

func migrationsLockJson() (*asset, error) {
	bytes, err := migrationsLockJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "migrations/lock.json", size: 6026, mode: os.FileMode(509), modTime: time.Unix(1533216138, 0)}
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
	"migrations/1533123066_initial_schema.down.sql": migrations1533123066_initial_schemaDownSql,
	"migrations/1533123066_initial_schema.up.sql": migrations1533123066_initial_schemaUpSql,
	"migrations/1533132482_add_number_to_review.down.sql": migrations1533132482_add_number_to_reviewDownSql,
	"migrations/1533132482_add_number_to_review.up.sql": migrations1533132482_add_number_to_reviewUpSql,
	"migrations/1533135217_add_repo_id_to_review.down.sql": migrations1533135217_add_repo_id_to_reviewDownSql,
	"migrations/1533135217_add_repo_id_to_review.up.sql": migrations1533135217_add_repo_id_to_reviewUpSql,
	"migrations/lock.json": migrationsLockJson,
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
	"migrations": &bintree{nil, map[string]*bintree{
		"1533123066_initial_schema.down.sql": &bintree{migrations1533123066_initial_schemaDownSql, map[string]*bintree{}},
		"1533123066_initial_schema.up.sql": &bintree{migrations1533123066_initial_schemaUpSql, map[string]*bintree{}},
		"1533132482_add_number_to_review.down.sql": &bintree{migrations1533132482_add_number_to_reviewDownSql, map[string]*bintree{}},
		"1533132482_add_number_to_review.up.sql": &bintree{migrations1533132482_add_number_to_reviewUpSql, map[string]*bintree{}},
		"1533135217_add_repo_id_to_review.down.sql": &bintree{migrations1533135217_add_repo_id_to_reviewDownSql, map[string]*bintree{}},
		"1533135217_add_repo_id_to_review.up.sql": &bintree{migrations1533135217_add_repo_id_to_reviewUpSql, map[string]*bintree{}},
		"lock.json": &bintree{migrationsLockJson, map[string]*bintree{}},
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
