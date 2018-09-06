// Code generated by go-bindata.
// sources:
// assets/dns/cluster-role-binding.yaml
// assets/dns/cluster-role.yaml
// assets/dns/configmap.yaml
// assets/dns/daemonset.yaml
// assets/dns/namespace.yaml
// assets/dns/service-account.yaml
// assets/dns/service.yaml
// DO NOT EDIT!

package manifests

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

var _assetsDnsClusterRoleBindingYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\x8f\x31\x4b\x04\x41\x0c\x46\xfb\xf9\x15\xc3\xf5\xbb\x62\x27\xe9\xd4\xde\xe2\x04\xfb\xdc\x4c\xce\x8b\xb7\x9b\x2c\x49\x66\x41\x7f\xbd\x0c\x83\x20\x28\x68\x37\xc5\x37\xef\xbd\x5c\x59\x2a\xe4\xc7\xa5\x79\x90\x1d\x75\xa1\x07\x96\xca\xf2\x9a\x70\xe3\x17\x32\x67\x15\xc8\x76\xc2\x32\x63\x8b\x8b\x1a\x7f\x60\xb0\xca\x7c\xbd\xf3\x99\xf5\x66\xbf\x4d\x2b\x05\x56\x0c\x84\x94\x73\xce\x82\x2b\x41\x2e\x83\x37\x55\xf1\x49\x37\x32\x0c\x35\xa8\xe2\xc9\xdb\xe9\x8d\x4a\x38\xa4\x29\x0f\xf5\x33\xd9\xce\x85\xee\x4b\xd1\x26\x91\xbe\x08\x7d\x3c\xde\xbe\x61\x21\xc8\xba\x91\xf8\x85\xcf\x31\x7d\x83\x27\xd3\x85\x8e\x74\xee\xee\x1f\x97\xa4\x3f\x6b\xfe\x21\x68\x4e\xf6\xd4\x47\xbd\xf8\xe0\xef\x1e\xb4\x82\x8f\x66\x1c\xcd\xf0\xeb\xcf\x2e\x38\xa4\xcf\x00\x00\x00\xff\xff\x71\x91\xce\x9d\x5f\x01\x00\x00")

func assetsDnsClusterRoleBindingYamlBytes() ([]byte, error) {
	return bindataRead(
		_assetsDnsClusterRoleBindingYaml,
		"assets/dns/cluster-role-binding.yaml",
	)
}

func assetsDnsClusterRoleBindingYaml() (*asset, error) {
	bytes, err := assetsDnsClusterRoleBindingYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "assets/dns/cluster-role-binding.yaml", size: 351, mode: os.FileMode(420), modTime: time.Unix(1, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _assetsDnsClusterRoleYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x2c\x8e\xb1\x6e\xc3\x40\x08\x40\x77\xbe\x02\x65\xb7\xab\x6e\xd5\xad\x1d\xba\x67\xe8\x8e\xef\x90\x82\x72\x81\x13\x70\xae\xd4\xaf\x8f\x1c\x67\x7b\x0f\xe9\x01\x77\xd1\x56\xf0\xbb\xcf\x48\xf6\xab\x75\x06\x1a\xf2\xcb\x1e\x62\x5a\xd0\x37\xaa\x2b\xcd\xbc\x99\xcb\x3f\xa5\x98\xae\xf7\xaf\x58\xc5\x3e\xf6\x4f\x78\x70\x52\xa3\xa4\x02\x88\x4a\x0f\x2e\x58\xcf\x35\x4b\xd3\x58\x6c\xb0\x53\x9a\x97\xa6\x01\x3e\x3b\x47\x81\x05\x69\xc8\x8f\xdb\x1c\x71\x44\x0b\x5e\x2e\x80\x48\x99\x2e\xdb\x4c\xbe\x72\xa4\x4b\x3d\xce\x44\x41\x9d\xbd\x03\xa2\x73\xd8\xf4\xca\xef\x82\xb5\x0d\x13\xcd\x78\x59\xb0\xef\x52\xf9\x94\x61\xed\x84\xe3\x99\x18\x74\xce\x77\xf6\xed\xdd\x76\x89\x7c\xc1\x1f\x65\xbd\xc1\x33\x00\x00\xff\xff\x13\x4d\xe1\x31\xfb\x00\x00\x00")

func assetsDnsClusterRoleYamlBytes() ([]byte, error) {
	return bindataRead(
		_assetsDnsClusterRoleYaml,
		"assets/dns/cluster-role.yaml",
	)
}

func assetsDnsClusterRoleYaml() (*asset, error) {
	bytes, err := assetsDnsClusterRoleYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "assets/dns/cluster-role.yaml", size: 251, mode: os.FileMode(420), modTime: time.Unix(1, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _assetsDnsConfigmapYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x6c\x8c\xcd\x4e\xc3\x30\x0c\xc7\xef\x7d\x0a\xbf\x40\x3b\xa6\xaa\x48\xe4\xba\x33\x57\xee\x26\x71\x96\x68\x69\x1c\xd9\xc9\x04\x02\xde\x1d\x95\x8d\xc0\xa4\xf9\xf4\xff\xf2\xef\x14\xb3\x33\x70\xe0\xec\xe3\xf1\x19\xcb\x80\x25\xbe\x90\x68\xe4\x6c\xe0\xbc\x1f\x56\xaa\xe8\xb0\xa2\x19\x00\x32\xae\x64\xc0\x65\xbd\x6a\x2d\x68\xc9\x00\x17\xca\x1a\xa2\xaf\xa3\x4d\x4d\x2b\xc9\xb8\x4d\x7e\x9f\x0e\x2c\xe4\x63\x22\x03\x9f\x03\x00\xc0\x64\x96\x79\x99\xe1\xe3\xc7\x6c\x47\x22\x2c\xda\x6d\x20\x4c\x35\x74\x7b\x6a\xaf\x24\x99\x2a\x29\x5c\xe9\x53\x62\x8b\x09\x62\x1e\xd1\x39\x99\x50\x0a\x42\x2c\x8f\x17\xf1\x87\xdd\xae\xb0\x53\x88\x59\xc9\x36\xa1\x9b\xa6\x15\xad\x42\xb8\xde\x84\x1e\x53\xaa\x41\xb8\x1d\xc3\x7d\x7c\x5f\x7f\x75\x55\x84\x57\xaa\x81\x9a\x82\x79\xda\x2f\xf3\xff\xe2\xed\x1d\x26\xd8\x51\xb5\x3b\x21\xe5\x74\x9e\x2c\x67\xdf\x07\x16\x6d\x20\x98\x1f\x7a\x20\x94\x18\xdd\x70\xe1\x7f\x07\x00\x00\xff\xff\x9c\xe9\xd3\xc8\x97\x01\x00\x00")

func assetsDnsConfigmapYamlBytes() ([]byte, error) {
	return bindataRead(
		_assetsDnsConfigmapYaml,
		"assets/dns/configmap.yaml",
	)
}

func assetsDnsConfigmapYaml() (*asset, error) {
	bytes, err := assetsDnsConfigmapYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "assets/dns/configmap.yaml", size: 407, mode: os.FileMode(420), modTime: time.Unix(1, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _assetsDnsDaemonsetYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x94\x52\x4d\x8b\xdb\x30\x10\xbd\xef\xaf\x18\x7c\xae\xd7\x09\x21\x65\xab\x5b\xd9\x40\x5b\xe8\x6e\x0d\x9b\xf6\x52\x7a\x98\x95\xc7\xb1\x88\xa4\x11\xd2\x38\x90\x7f\x5f\x1c\xc7\xae\x4c\xd3\xc2\xea\x64\xf4\x3e\x34\xf3\xfc\x8e\xc6\x37\x0a\x76\x48\x8e\xfd\x0b\xc9\x1d\x06\xf3\x83\x62\x32\xec\x15\x60\x08\xa9\x3a\xad\xef\x1c\x09\x36\x28\xa8\xee\x00\x3c\x3a\x52\xd0\xf8\x74\xfd\x4e\x01\x35\x29\xe0\x40\x3e\x75\xa6\x95\x52\xdb\x3e\x09\xc5\x72\xa4\x58\x7c\x25\x9b\x06\x21\x64\x1c\x0c\x61\xf4\x48\x81\xf4\x00\x26\xb2\xa4\x85\xe3\x48\x74\x28\xba\xfb\x9a\x29\x6f\x6a\x01\x84\x5c\xb0\x28\x74\x55\x65\x53\x0e\xc7\x2e\x0c\xfe\x61\x01\x30\x8d\x70\xf9\xa6\x78\x32\x9a\x3e\x6a\xcd\xbd\x97\xe7\x6c\xd5\xe1\x68\xf6\x82\xc6\x53\x9c\x4d\xcb\x45\x1c\xe3\x31\x0e\x0f\xa4\xa0\x68\x58\x1f\x29\xde\x1b\xae\xe6\x87\x2b\x8e\xe6\x60\x7c\xa9\x39\x52\xe3\x93\x3a\x6d\xee\xd7\xeb\xfb\x55\xb1\xd4\xd6\xbd\xb5\x35\x5b\xa3\xcf\x0a\xbe\xb4\xcf\x2c\x75\xa4\x44\x5e\x66\x96\x66\xe7\x70\xf8\x6b\x3f\xa1\xb8\x5a\x15\xf0\x6b\x86\x31\x1e\xd2\x05\x2b\x35\xfb\xb6\x78\x07\x45\x45\xa2\xab\x2b\xb3\x7a\xe4\x48\xad\xb1\x94\x4b\x4e\x6c\x7b\x47\x4f\xc3\xd2\x59\x60\xd3\x76\x83\x8d\x39\x94\x23\x69\x46\x01\xdc\xc0\xaf\x51\x3a\x05\xf9\x0b\x19\x23\x12\x36\xdf\xbc\x3d\x2b\x90\xd8\xff\x91\x06\x8e\xcb\x77\xe6\x64\x6b\x8e\xa2\x60\xbb\xd9\x6e\x32\x97\xbf\x33\x06\x08\x91\x85\x35\x5b\x05\xdf\x77\xf5\xdb\x9d\x4a\xd1\xe1\xa6\xdb\xfe\xf1\x3f\x6e\x1f\xd6\x37\xdc\x1c\x49\x34\xfa\xf6\x6c\xb9\x9b\x35\x27\xf2\x94\x52\x1d\xf9\x95\x54\x46\xef\x44\xc2\x27\x92\xfc\x0a\x20\x8c\xb1\x76\x84\x56\xba\x25\x72\x19\xe5\x61\xf5\xb0\x5a\x5c\x27\xdd\xd1\x30\xce\xe7\xfd\xbe\xce\x00\xe3\x8d\x18\xb4\x3b\xb2\x78\x7e\x21\xcd\xbe\x49\x0a\xde\xe7\x52\x31\x8e\xb8\x97\x19\xdc\x66\x58\xea\xb5\xa6\x94\xf6\x5d\xa4\xd4\xb1\x6d\x14\xac\x33\xb4\x45\x63\xfb\x48\x19\x3a\x69\x1b\x9f\xa6\x06\xef\xa8\xc5\xde\x4e\xe5\x1d\x3b\xf4\x86\x8e\x8d\xf7\x4f\x18\x96\xf1\xdc\xaa\x04\x80\x11\x72\x69\x49\x2c\xe1\x48\x67\x05\x53\xe9\x17\xd8\x94\xf2\x0c\xfe\x0e\x00\x00\xff\xff\x98\x16\x93\x48\x0c\x05\x00\x00")

func assetsDnsDaemonsetYamlBytes() ([]byte, error) {
	return bindataRead(
		_assetsDnsDaemonsetYaml,
		"assets/dns/daemonset.yaml",
	)
}

func assetsDnsDaemonsetYaml() (*asset, error) {
	bytes, err := assetsDnsDaemonsetYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "assets/dns/daemonset.yaml", size: 1292, mode: os.FileMode(420), modTime: time.Unix(1, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _assetsDnsNamespaceYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x44\xca\xb1\xae\x42\x31\x08\x06\xe0\x9d\xa7\x20\x67\xef\xbd\x71\xe5\x21\x1c\xdd\x49\xfb\x1b\x89\xa7\xd0\x14\xf4\xf9\x8d\x93\xfb\xf7\x34\x1f\xc2\x57\x9d\xc8\xa5\x1d\xa4\xcb\x6e\xd8\x69\xe1\xc2\xef\x0b\x4d\x94\x0e\x2d\x15\x62\x76\x9d\x10\x8e\x05\xcf\x87\xdd\xab\xf5\xf3\x95\x85\xdd\x86\x27\x31\xab\x7b\x94\x96\x85\xe7\x17\xf3\x0f\xfe\x59\xfc\x7b\x0c\xb4\xc4\x89\x5e\xb1\x85\x8f\x83\x3e\x01\x00\x00\xff\xff\xb5\x9f\xce\xf1\x79\x00\x00\x00")

func assetsDnsNamespaceYamlBytes() ([]byte, error) {
	return bindataRead(
		_assetsDnsNamespaceYaml,
		"assets/dns/namespace.yaml",
	)
}

func assetsDnsNamespaceYaml() (*asset, error) {
	bytes, err := assetsDnsNamespaceYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "assets/dns/namespace.yaml", size: 121, mode: os.FileMode(420), modTime: time.Unix(1, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _assetsDnsServiceAccountYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x2c\xca\x31\x0e\x82\x21\x0c\x06\xd0\x9d\x53\xf4\x02\x0c\xae\xdd\x3c\x83\x89\x7b\x53\x3e\x63\xa3\x14\x42\x0b\xe7\x37\x26\xff\xf6\x86\xf7\x31\x6f\x4c\x0f\xac\x63\x8a\xbb\xea\xd8\x9e\x45\xa6\x3d\xb1\xc2\x86\x33\x9d\x5b\xe9\x48\x69\x92\xc2\x85\xc8\xa5\x83\xa9\x79\x5c\x8e\x29\x0a\xa6\x31\xe1\xf1\xb6\x57\x56\xfd\xee\x48\xac\xfa\x2f\xbf\x00\x00\x00\xff\xff\x35\xeb\xbe\x6a\x5d\x00\x00\x00")

func assetsDnsServiceAccountYamlBytes() ([]byte, error) {
	return bindataRead(
		_assetsDnsServiceAccountYaml,
		"assets/dns/service-account.yaml",
	)
}

func assetsDnsServiceAccountYaml() (*asset, error) {
	bytes, err := assetsDnsServiceAccountYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "assets/dns/service-account.yaml", size: 93, mode: os.FileMode(420), modTime: time.Unix(1, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _assetsDnsServiceYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x84\x8e\x3d\xae\x02\x31\x0c\x84\xfb\x9c\xc2\x17\x48\xf1\xf4\x44\x93\x16\x0e\xb0\x12\x3f\xbd\x49\x0c\x44\x64\x13\xcb\x36\x7b\x7e\xb4\x11\xda\x85\x02\x51\x8e\xe7\xf3\xa7\xb9\xe7\x9a\x02\xec\x49\xa6\x1c\xc9\x21\xe7\x13\x89\xe6\x56\x03\x4c\x7f\x6e\x24\xc3\x84\x86\xc1\x01\x54\x1c\x29\x40\xaa\x0a\xaf\xa0\x8c\x91\x02\x34\xa6\xaa\xb7\x7c\x31\x1f\xcb\x43\x8d\xc4\xa7\xaa\x0e\xa0\xe0\x99\x8a\xce\x9f\xf0\xc6\x20\x73\x97\x38\x65\x8a\x73\xa9\x54\x28\x5a\x93\xaf\x20\x00\x37\xb1\x2e\xf2\xeb\x88\x4e\xcf\x45\x80\xcd\x7f\x0f\x86\x72\x25\x1b\xfa\x69\x01\xa4\x59\x8b\xad\x04\x38\xee\x86\x4f\x81\xb7\xc8\x3f\x25\x2b\xb4\x88\x0e\xdb\xc1\x3d\x03\x00\x00\xff\xff\xfe\x04\x0c\xac\x34\x01\x00\x00")

func assetsDnsServiceYamlBytes() ([]byte, error) {
	return bindataRead(
		_assetsDnsServiceYaml,
		"assets/dns/service.yaml",
	)
}

func assetsDnsServiceYaml() (*asset, error) {
	bytes, err := assetsDnsServiceYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "assets/dns/service.yaml", size: 308, mode: os.FileMode(420), modTime: time.Unix(1, 0)}
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
	"assets/dns/cluster-role-binding.yaml": assetsDnsClusterRoleBindingYaml,
	"assets/dns/cluster-role.yaml": assetsDnsClusterRoleYaml,
	"assets/dns/configmap.yaml": assetsDnsConfigmapYaml,
	"assets/dns/daemonset.yaml": assetsDnsDaemonsetYaml,
	"assets/dns/namespace.yaml": assetsDnsNamespaceYaml,
	"assets/dns/service-account.yaml": assetsDnsServiceAccountYaml,
	"assets/dns/service.yaml": assetsDnsServiceYaml,
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
	"assets": &bintree{nil, map[string]*bintree{
		"dns": &bintree{nil, map[string]*bintree{
			"cluster-role-binding.yaml": &bintree{assetsDnsClusterRoleBindingYaml, map[string]*bintree{}},
			"cluster-role.yaml": &bintree{assetsDnsClusterRoleYaml, map[string]*bintree{}},
			"configmap.yaml": &bintree{assetsDnsConfigmapYaml, map[string]*bintree{}},
			"daemonset.yaml": &bintree{assetsDnsDaemonsetYaml, map[string]*bintree{}},
			"namespace.yaml": &bintree{assetsDnsNamespaceYaml, map[string]*bintree{}},
			"service-account.yaml": &bintree{assetsDnsServiceAccountYaml, map[string]*bintree{}},
			"service.yaml": &bintree{assetsDnsServiceYaml, map[string]*bintree{}},
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
