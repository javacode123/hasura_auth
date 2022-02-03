package schema

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"strings"
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

var _schema_schema_graphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x2a\x4e\xce\x48\xcd\x4d\x54\xa8\xe6\x52\x50\x50\x50\x28\x2c\x4d\x2d\xaa\xb4\x52\x08\x04\x51\x60\x81\xdc\xd2\x92\xc4\x92\xcc\xfc\x3c\x2b\x05\x5f\x28\x8b\xab\x96\x8b\xab\xa4\xb2\x20\x15\xa2\x08\xaa\xcf\x27\x3f\x3d\x33\x4f\x41\x23\x2f\x31\x37\xd5\x4a\x21\xb8\xa4\x28\x33\x2f\x5d\x51\xa1\x20\xb1\xb8\xb8\x3c\xbf\x28\x05\x2e\xa2\x69\xa5\xe0\x55\x9c\x9f\x17\x9e\x9a\x14\x92\x9f\x9d\x8a\x30\x09\x66\x34\x85\x86\x01\x02\x00\x00\xff\xff\x8f\xf5\xa0\xac\xca\x00\x00\x00")

func schema_schema_graphql() ([]byte, error) {
	return bindata_read(
		_schema_schema_graphql,
		"schema/schema.graphql",
	)
}

var _schema_type_auth_graphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x2a\xa9\x2c\x48\x55\xf0\x2a\xce\xcf\x0b\x4f\x4d\x0a\xc9\xcf\x4e\xcd\x53\xa8\xe6\x52\x50\x50\x50\x28\x01\xb3\xad\x14\x82\x4b\x8a\x32\xf3\xd2\x15\xb9\x6a\x01\x01\x00\x00\xff\xff\x47\xba\x84\xb5\x29\x00\x00\x00")

func schema_type_auth_graphql() ([]byte, error) {
	return bindata_read(
		_schema_type_auth_graphql,
		"schema/type/auth.graphql",
	)
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		return f()
	}
	return nil, fmt.Errorf("Asset %s not found", name)
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
var _bindata = map[string]func() ([]byte, error){
	"schema/schema.graphql": schema_schema_graphql,
	"schema/type/auth.graphql": schema_type_auth_graphql,
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
	Func func() ([]byte, error)
	Children map[string]*_bintree_t
}
var _bintree = &_bintree_t{nil, map[string]*_bintree_t{
	"schema": &_bintree_t{nil, map[string]*_bintree_t{
		"schema.graphql": &_bintree_t{schema_schema_graphql, map[string]*_bintree_t{
		}},
		"type": &_bintree_t{nil, map[string]*_bintree_t{
			"auth.graphql": &_bintree_t{schema_type_auth_graphql, map[string]*_bintree_t{
			}},
		}},
	}},
}}
