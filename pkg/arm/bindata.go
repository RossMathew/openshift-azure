// Code generated by go-bindata. DO NOT EDIT.
// sources:
// data/azuredeploy.json
// data/startup.sh
package arm

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

var _azuredeployJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x5b\xdd\x6f\xdb\x38\x12\x7f\xef\x5f\x21\xe8\x0e\x70\xb3\x88\x6c\x39\x71\x7b\x45\xde\xd2\xe4\x7a\x35\x6e\xd3\x1a\x51\xbb\xf7\xb0\x08\xba\x34\x35\xb6\x79\x91\x48\x81\xa4\x9c\xba\x81\xff\xf7\x05\xa9\x8f\xda\xfa\x76\x2c\x27\x4e\x6a\xf5\x21\xa9\x44\x0e\x67\x86\xbf\x19\xce\x07\x73\xff\xca\x30\x0c\xc3\xfc\xa7\xc0\x33\xf0\x91\x79\x66\x98\x33\x29\x03\x71\xd6\xeb\x45\x6f\xba\x3e\xa2\x68\x0a\x3e\x50\xd9\x45\x3f\x42\x0e\x5d\xcc\xfc\xf8\x9b\xe8\x9d\xd8\xfd\x37\x96\xdd\xb7\xec\x7e\xcf\x85\xc0\x63\x0b\x35\xee\x0b\xf8\x81\x87\x24\x74\xff\x2f\x18\xfd\x87\x79\x1c\xad\x80\x19\x95\x40\xe5\x1f\xc0\x05\x61\x54\x2d\xd4\xef\xda\xea\x5f\x32\x20\x40\x1c\xf9\x20\x81\x0b\xf3\xcc\xb8\x5f\xc6\x6f\xe7\x88\x13\x34\xf6\x60\xed\x25\x07\xc1\x42\x8e\xf5\xcb\x3f\xf5\x2b\xf5\xdc\xa7\xbf\xe9\x41\x72\x11\x80\x5a\xe6\x8a\x60\xce\x04\x9b\xc8\xee\x27\x90\x77\x8c\xdf\xf6\x68\xf4\xd3\x01\x1c\x72\x22\x17\xff\xe1\x2c\x0c\x44\xcc\x46\x3a\x1d\x05\x64\x85\xd7\x13\xbb\xff\xd6\xb2\x4f\xad\x53\x3b\x3b\xce\x63\x18\xc9\x78\xd4\xfd\xbd\xd1\xbd\x42\x94\x4c\x40\xc8\xee\xef\xf1\x07\x63\xb9\xcc\xce\xa1\xc8\xd7\xac\x51\x31\xb5\x30\xf3\x83\x50\x42\x76\x48\xc0\x59\x00\x5c\x92\x48\xf0\xb5\x6f\xfa\xbb\x88\x99\xbf\x0e\xbd\x8c\x1a\x56\x9f\xfc\xc4\x1c\x0f\xc8\xf3\xd8\xdd\x37\x21\x66\x19\x0e\x36\xe1\x66\x5d\x71\x18\x83\x50\xe3\xcc\x73\x45\xba\x82\xac\x1e\xee\x82\xc0\x9c\x04\x89\x0a\xf5\x1c\xc3\x71\x3e\x1a\x92\xa3\xc9\x84\xe0\x06\xf3\x25\xa1\x5a\xd3\xe7\xae\xcb\x41\x88\x11\x87\x09\xf9\xae\x88\xfd\xb6\xc1\xe4\x11\xe3\xf2\x1a\xd1\xa9\x56\xca\xc9\x89\x75\x72\x52\x3b\x99\x70\xc0\x09\xdf\x43\x3a\x66\x21\x75\xeb\xe6\x04\x9c\x30\xb5\x6d\xe6\x99\xd1\xb7\xfb\xb5\x83\x99\x64\x98\x79\x8a\xfe\x17\x1c\xd4\xd1\x8e\x6c\x62\x73\x25\x44\xf3\xd6\xe4\xff\xcd\x2c\x9d\xb2\x2c\xfc\x92\x7f\x7b\xf3\xaa\xf8\xfb\xf2\xf8\x99\x1b\x2c\xa1\x13\x8e\x0e\xe6\x7a\x30\xd7\xfc\xe0\xe7\x63\xae\xc5\x6b\x34\xc6\x9f\x8a\x11\xf6\x09\x80\x8a\x9f\x27\x45\xe0\xbb\xac\x9b\xc9\xcf\xdc\x16\x7e\x27\x07\xf8\xad\xc2\x2f\xeb\xff\xd7\x86\x3f\x05\xfe\xc4\x93\x02\x70\x30\x38\xdd\x39\x02\x4f\x5f\x10\x02\x77\x1a\xaf\xcc\x09\x97\x21\xf2\xe2\xff\x3e\x49\xa4\x32\xa7\x20\x37\x0d\x52\x50\xb4\x15\x4e\x80\x30\x94\xda\x4d\x32\x2a\xda\xb0\x8a\x60\x46\x0f\xee\xdb\x51\x92\xd7\x7b\x57\xbc\x43\x37\xb9\xb7\x05\xbe\xc1\x14\xe1\x98\x82\xdc\x2a\x6e\x8a\x48\xb4\xe6\x33\xb2\xa0\x4d\xe5\x3c\x19\xec\x17\x14\x83\x70\xec\x11\x3c\x1c\xc5\x66\x06\xf5\x60\xfc\x97\x65\xbf\xb3\xec\x7e\x9b\x60\x24\x81\xc5\x59\x28\x81\x6f\x8a\xc8\x94\x7b\x2f\x59\xfe\x0a\xe4\x8c\xb9\x8a\xa8\x23\x91\x24\x78\x5d\xdb\x19\xf0\x98\xe2\x36\x2c\x26\x9c\x70\xf6\x1e\x89\x1c\x8d\xad\xf4\xed\x31\xe4\xbe\x47\x1e\xa2\x18\x78\x13\x5d\xf7\xed\x02\x5d\xbb\x10\x00\x75\xc5\x67\x5a\x88\x78\xf3\xcf\xa4\xfe\x31\x74\x5f\x77\x1a\x6c\x79\xe7\xd8\xe8\xa4\x5b\xd0\x39\xba\x59\x97\xf7\xa6\xc5\x8d\xf6\xc6\x0f\xdc\xe8\x31\xc2\xb7\x40\xdd\xe4\x30\x60\xcc\xdb\xca\xda\x63\x72\xc5\xa6\x58\x60\x70\x05\x3e\x67\xc2\x75\xc1\xca\x1d\x8e\x2e\x18\x9d\x90\x69\xc8\xb5\xf0\x5b\xb1\x95\xd0\x6c\xcb\x0d\x05\x9c\xcc\x91\x84\x62\xfb\xb8\x5c\x50\xe4\xd7\x47\x23\x19\xb4\xd4\x2e\xaa\x27\x11\xbd\x42\xbb\x38\xcc\x3e\xc5\xde\xb2\xfc\x4b\xb3\x6d\x25\x51\xdc\xf3\x09\xc9\x14\x64\xd5\xc3\xd2\x8c\xbd\x68\xd8\x4f\x6b\x27\x74\xba\x75\x6e\x1f\x30\x2e\xad\xca\x4c\x62\x23\x74\xe4\x6d\x6a\xa3\xcd\xc5\x8c\x62\x24\x5f\x57\xef\xf1\x9a\xbb\x53\xfb\x9b\x7a\x80\xce\xd1\xb1\xd1\xe9\x15\xd8\x75\xf2\xae\x1e\x00\x35\xc8\x8d\xe9\xa8\x40\xd0\x3c\x33\xde\xd9\x35\xc3\x81\xa2\xb1\x07\x1f\x3c\x86\x24\xa1\xd3\xe1\xc8\x3c\x33\x26\xc8\x13\x50\x33\xad\xc4\x0f\x3c\xbe\x2a\xcb\x1c\x52\xfa\x61\x6b\x85\x26\x84\x9a\x6a\x94\xb8\x1e\x7c\x21\x3e\xb0\x50\x0e\xe9\x15\xa1\xa1\xd4\xb0\x1c\xd4\x4c\x53\x82\x5e\x12\x21\x39\x19\x87\xc9\x41\x73\x09\x13\x14\x7a\x55\xf1\x99\x11\xc3\x7f\x5c\x1e\x9a\x66\x98\x6b\x4f\xf5\x7a\x59\xd1\x8b\xed\x73\x6b\x3d\x67\xf3\xa4\x47\xcb\x9e\xb5\x00\xd5\x99\xe2\xc1\xc3\xe4\xe5\x8b\xed\x61\x30\xa8\x4b\x80\x0f\x2e\x66\x43\x17\x53\xaf\xd2\x5f\xd2\xc7\x0c\x06\xa7\x4f\xee\x64\x1a\x45\x53\x2c\x94\x8d\xe2\xa4\x48\xb8\x3d\x8a\x8d\x08\x95\xc0\xe7\xc8\x1b\x52\x07\x30\xa3\xae\x9a\xf2\xa6\x46\xa3\x34\xf4\xc7\xc0\x3f\x4f\x46\x89\x30\xb5\x15\xd9\x86\xe7\xe8\x8b\x39\x0f\x1e\x53\xab\xf5\xae\xa3\x7d\x0b\xd8\x93\x5a\x83\x23\x19\x47\x53\xe8\x89\xe8\xe7\x39\xc6\x2c\xa4\xb2\xbe\xda\xf0\xc6\xb2\xdf\x5a\xfd\x37\x6d\x56\x76\xd4\xf8\xe8\xac\xe8\x5e\xc3\x54\x39\xdc\x85\xb3\xc6\x55\xc1\xdc\xda\x32\x64\x34\xf3\x4b\x2c\xb7\x23\x11\x75\x11\x77\xbf\xfd\x7e\xed\xb4\xa1\xbd\x8b\xe8\xc2\x45\x0f\xcd\x11\xf1\xd0\x98\x78\x44\x2e\x1c\x68\xa0\xbe\xb7\x96\x3d\xb0\x4e\x6d\x2b\xe0\x30\x27\x90\x6d\x14\x6c\xa5\x46\x24\x1e\x7a\x0f\x24\xba\x9e\xa3\xce\x21\xc9\xc3\x82\x40\xc3\x0c\x3c\x24\x27\x8c\xfb\x1f\xd4\xd9\x77\xc9\x7c\x44\xe8\x85\x52\x6f\xb1\xa5\xa5\xc3\xbf\x06\x2e\x92\xb0\x3e\xfe\x74\x1f\xeb\x92\x98\x05\x8b\x62\xd5\xe0\x98\xed\xb5\xcd\x88\x77\xff\x22\xc6\x66\x81\x06\xd2\x5a\x16\x63\x41\x75\x75\x71\x9b\x1d\x4f\x82\x86\x0e\x09\x92\xad\xb7\x3a\xc7\x86\x92\x66\x48\x5d\xf8\xfe\xfa\xe8\xe8\xa6\xcd\x6a\x69\x52\x0d\xda\x13\x17\x96\xb9\xd9\x31\x54\x07\xc7\x04\xe1\x06\x30\x28\xe9\x95\x3c\x21\x0c\xb6\xab\xd6\x66\x7a\x45\x2a\x32\x9c\x53\x90\x9d\xdc\xf6\x37\x21\x56\x78\x51\x46\x91\x5c\xb9\x69\xf6\x30\xca\x45\xc5\xbc\x7a\x08\xef\xb2\xd6\x9c\x5a\x10\x25\xb8\x35\x13\x8a\x32\xb8\x73\x8c\xc1\x03\x8e\x24\xb8\xb1\xfc\x84\x4e\x75\x01\x57\x65\x73\x45\xda\x8b\xe6\x0d\x47\x1f\x18\xbf\x43\xdc\x8d\x86\x97\x78\x64\x12\xb4\x57\x51\x26\x01\xd6\xb4\x5a\xac\x28\xfb\x88\x2f\x4a\x99\xcf\x0e\xde\xba\xfc\x1c\xf7\xe5\x5a\x4f\xb7\x4a\xcd\x4a\xe5\x5a\x71\x3f\x31\xfe\xb9\x7d\xaa\x15\xd9\x46\xb0\xb3\x12\x7a\xb0\x9d\xd5\xe5\xe4\xd9\x49\x2e\x58\xe4\x7b\xca\xbb\xc8\x8d\xc4\x6f\xea\xce\x72\x4b\x2c\x5b\x8c\x13\x63\x24\x5d\x21\x3c\x23\xb4\xd9\x01\x55\x15\x26\x3e\xdb\x83\x2a\x77\x54\xaf\x60\xb1\xc2\x05\x6f\x74\xda\x94\xc5\xe6\x6a\xcf\x7f\xc6\xc8\x3b\x6d\x61\xa6\xc7\xca\xdc\x6f\xed\x54\xc9\x88\xf3\x40\xab\x78\xa0\x72\x0a\x90\xa0\xd7\x9a\x21\xee\xde\x21\x0e\x23\xce\x26\xc4\xab\xb8\xf0\x31\xf7\x1d\xf2\x03\x72\xba\xfc\xe3\x4a\xbd\x56\x9a\x6c\xb4\x5c\x0c\x9f\xda\xd5\xf2\x11\x61\xd5\x05\x93\xba\x82\xc4\x06\x5e\xe6\x01\xc0\xde\xd0\x8f\x36\xbb\xec\xc2\x44\xad\x8e\x90\xeb\x13\xfa\x55\x00\x4f\x40\x8b\x3d\x16\xba\x56\x28\x72\x8d\xf7\x74\x4a\x2c\x04\xdf\x02\xe6\x29\x2d\x8f\xd0\xf0\x7b\xf3\x52\xb5\xe9\x12\xa1\xa2\xa3\x11\x12\xe2\x8e\x71\xf7\x3c\x94\x33\xa0\x92\xa4\x36\x5a\x1d\x6a\x98\x42\xcc\x1a\x04\x2e\xfa\xa4\xfc\x2f\x2c\xaa\xf1\x92\x3c\xf5\x07\xb4\xa6\x7a\x0b\x8b\x4b\x24\x51\x8c\x7d\xc7\xf9\x38\x4a\x96\x39\x17\x8e\xe4\x84\x4e\xd3\xf2\xc7\xea\xc7\xbc\x7b\x29\xe7\x1b\x49\x25\x9d\xd9\x9b\x31\x1f\x7a\x3f\xf7\xb1\xd7\x15\x62\xd6\x43\xa1\x9c\x31\x4e\x7e\x80\xfb\xed\x56\x89\x56\x4b\xb3\xfc\x68\x37\x0a\x01\x58\x3d\x2f\xff\xb6\xf0\x76\x56\x54\xf0\xa9\x45\x2d\xf1\xd1\x14\xae\x61\x02\x1c\x68\xc5\x05\x33\xa3\xc4\x70\xd7\x9c\xcf\x30\xa2\x15\x7d\xd5\x71\x81\xb1\x5c\x2a\x47\x98\x77\x97\x7a\x59\xed\x24\xcb\x29\x7c\x42\xbe\xf2\x64\xa5\x91\x60\x49\x04\x68\x32\x71\x49\xc4\x6d\xb5\x28\x58\x07\x0e\x3a\x7d\xb8\x06\xe4\xfe\x8f\x93\x5c\x95\x67\x7d\x3c\x07\x24\xe1\x73\x7a\xc3\xf4\x03\x67\xbe\x66\xb6\xe9\x6d\x96\x1d\x06\x3f\x3d\xf8\x2e\x81\x0a\x9d\xb5\xfc\x2a\x89\x7a\x89\x2a\x56\x0e\x89\x72\x17\xfa\xb4\x81\x8a\x4a\x75\xb0\x28\x48\xf8\xeb\x03\x96\x50\xb2\xaf\xc1\x94\x23\x17\xae\x08\x65\xfc\xe7\xce\x96\x55\x19\x39\x93\x80\x25\xb8\x0e\x48\x49\xe8\xb4\x3c\x09\x32\xa3\xeb\xd3\xb1\xc4\xef\x91\x80\xb7\x03\xe3\xb5\x23\x11\x97\x61\x60\xfc\x15\x4b\xf2\xd7\x51\xe3\xc0\x42\x7b\x7e\x31\x03\xbe\x8e\xe0\x73\xfd\xc7\x8a\xff\x2e\xc3\x6b\xc4\xca\x0a\xb3\x45\xa4\x13\xbb\xb8\x08\x85\x64\xbe\x13\x31\x5e\x32\xee\x23\xa2\xae\x07\x7c\xd5\x06\xba\xf6\x0b\xae\x57\x3f\xe8\xcf\xa0\x0e\xd5\xea\xcd\xbc\xdf\x50\x29\x79\xaf\x6a\xd5\x7a\xdb\x0f\x95\xea\xc7\xab\x54\xb7\x07\x81\xe7\x52\xa7\xd6\x10\x7b\x18\xdd\xaa\x4b\x16\x3b\xa8\x7a\x17\x1a\xc3\xa3\xd5\xbc\xdb\x31\xc5\x5f\xbd\xe2\xbd\x0a\x99\xf7\x1b\xdc\x42\x6f\xc6\xf6\xda\x52\xfb\x72\xed\xcd\xa8\x4f\x14\xeb\x2b\xff\x87\x36\x41\xf6\xd9\xdb\x36\x41\x23\x37\x95\x93\x66\xeb\x5a\xc1\x3e\x34\x09\xd2\xb3\xe4\x31\xb3\xe4\x47\x6e\x11\xec\x4b\x84\x50\x57\x47\x2d\x83\x61\x6b\xed\x81\xe2\xcd\xde\x55\xce\xdd\xce\xf1\xfb\x28\xad\x81\x0a\x2b\x38\x34\x06\xda\x07\xf4\x86\x9e\x73\x5f\xdb\x02\x8d\x00\x9e\x52\x3a\x34\x05\x0e\x4d\x81\x15\x11\x0e\x4d\x81\x43\x53\xe0\xc5\xd6\x44\x1a\xb5\x04\x9e\xae\x38\x50\xb2\xf8\x73\x6c\x07\x68\x39\x5e\x7a\x33\x40\xff\x16\x23\xc1\x64\xa1\x0c\x42\x19\x71\xf1\x6a\xf9\xea\xef\x00\x00\x00\xff\xff\x42\xad\xb0\x92\xa0\x51\x00\x00")

func azuredeployJsonBytes() ([]byte, error) {
	return bindataRead(
		_azuredeployJson,
		"azuredeploy.json",
	)
}

func azuredeployJson() (*asset, error) {
	bytes, err := azuredeployJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "azuredeploy.json", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _startupSh = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x94\x55\x7f\x4f\xe3\x46\x10\xfd\x7f\x3f\xc5\x34\x54\x4d\x7b\x92\xbd\x70\x3d\xa1\x96\x0b\x48\x24\x18\x84\xf8\x11\x14\xd2\x93\x4e\x55\x85\xd6\xbb\x63\xb2\x8a\xbd\x6b\x66\xc7\x11\x69\x94\xef\x5e\x39\x0e\x5c\x12\x02\x85\xff\x92\x99\x37\xef\x3d\x3d\x27\xcf\x3b\x3f\xc9\xd4\x3a\x99\xaa\x30\x82\x08\x1f\x85\xd8\x81\x61\xff\xa4\x7f\x00\x12\x59\x4b\xe3\x42\xa1\xc2\x43\x6c\xa4\x27\x7b\x6f\x5d\x54\x95\x81\x09\x55\x11\x19\x17\x62\xed\x5d\x06\x36\x80\xae\x88\xd0\x71\x3e\x85\x91\x22\xa3\xbd\x41\xf3\x15\x2c\x8b\x1d\x28\xc9\xa7\x2a\xcd\xa7\x10\x46\xbe\xca\x8d\x6b\x33\xa4\x28\xc4\x6d\x32\xf8\x76\xde\x4b\xee\x86\xdf\x6f\x92\xc3\x86\x59\xd8\x0c\xfe\x86\x28\x83\xd6\x42\x38\x4c\x43\xcd\x6e\xef\xa5\x62\x5f\x58\x1d\xf9\x12\x5d\x18\xd9\x8c\x23\xe7\x0d\xb6\xe0\x9f\xaf\xc0\x23\x74\x02\x00\x60\x8d\x6e\x13\x2f\x32\x2b\x44\x98\x06\xc6\x42\x73\x0e\x81\x7d\x09\xc6\xeb\x31\x52\x1c\x90\x26\x56\xa3\x28\xc6\x59\x88\x1f\xb3\x50\xcb\x4b\x83\x13\x19\x4c\xba\x27\x50\x8f\x3c\xb4\x9f\xbf\x03\xc8\x89\x22\x99\xdb\x54\x36\xe7\x00\xf5\x09\xdc\x53\xf9\x50\x79\x56\x00\xbb\xb0\xdb\x86\xa3\xa3\x85\xff\x2c\xb0\x4a\x45\xe1\x2b\xc7\x9b\x77\x82\x30\xb0\x27\xd4\xde\x41\x34\x78\xb1\x5d\xb5\xaa\x88\x37\xbd\x8a\xd9\xcc\x66\x80\x0f\x10\x27\x8f\x4c\x2a\x1e\xf8\x1c\xa1\x65\x5d\x46\xaa\x35\x9f\x37\xa6\x5b\xdd\x7e\x7f\x78\x3b\x1c\x1c\xdf\xdc\xf5\xfa\xd7\xa7\xe7\x67\x77\xd7\xc7\x57\xc9\x61\x1d\x5c\xd4\xa4\x1a\x35\x07\x4f\x6e\x7f\xa4\xfd\xf3\x6c\x35\xcc\xf9\x22\x6c\x31\x9b\x61\x1e\xf0\x03\xec\xda\x17\x65\xc5\xf8\x01\x7e\x67\xe6\x73\x21\x02\x1a\x88\x2c\x44\x08\xad\xb0\x73\x92\x74\xff\x3a\xbb\xbb\xec\x9f\x5d\x26\xdf\x92\xcb\xc3\xcf\x9b\x83\x2f\x3b\x2d\x78\x0f\xbb\xa0\x02\x22\xca\x1a\x2c\xb2\x36\xf2\x53\xf3\xb9\xf9\xe1\xc9\x42\x05\x46\x92\x9f\x84\x48\x55\xc0\xfd\x2f\x10\x19\xe8\x74\x3a\x30\x9b\x41\xb7\x19\xfc\xfa\x5d\x15\xf9\x95\xa2\x30\x52\x39\xc4\xbd\x85\x56\x7c\xed\x0d\x76\xbd\xe7\xc0\xa4\xca\x8b\x2a\xc5\xc6\xc3\x6f\x30\x9f\xc3\xd1\x2a\x7f\x6d\x42\xa6\x4f\xc8\x78\xfc\x0c\x7d\x5d\xaf\x87\xc4\xc7\xa1\x3b\x65\x0c\xdb\xf5\x6a\xc0\x2b\x4a\x8b\xe7\xf0\x2c\x57\x22\xc5\x9a\xf8\x75\xa9\x1b\xb2\x13\xc5\x78\x81\xd3\x37\x05\x2f\x70\xfa\x6e\xbd\x31\x4e\x85\x1e\x15\xde\xc0\xee\xfe\xee\x2e\xbc\xef\xe2\x25\x6c\x6b\x64\x1f\xcb\xac\xa7\xde\x08\x4a\xab\x45\x32\xba\x7c\x29\xdd\xac\x9a\x79\x39\xb6\x52\xab\x88\xa9\x0a\x2c\x83\xaf\x48\xa3\x54\x4e\x8f\x3c\x05\xf9\xa3\x94\x96\x64\x55\x69\x14\x63\xf4\x84\x17\xcb\x12\x71\xaa\xc0\xfa\x0f\x8c\x04\x7b\xfb\x7f\xc4\xfb\xbf\xc7\x7b\x9f\xff\x8c\xf7\xf6\xdb\x5b\x6c\x11\x06\x9f\x4f\x16\xdd\x2a\x8a\xb1\xb1\x04\xd1\xba\x43\x9d\xfb\xca\x94\xe4\x27\xd6\x20\x09\xad\x78\x9d\x64\x6d\x2d\xd5\xbf\x15\x61\x53\xd4\x9d\x4e\x3b\xe9\x9f\xb6\x05\xa3\x53\x8e\xcf\xcd\x41\x9d\x5e\x7c\xa5\x9c\xcd\x30\x70\x3c\x6c\xc6\x27\x30\x9f\x8b\x50\xa5\x41\x93\x2d\xd9\x7a\xf7\x02\x78\xbb\xba\x5c\xc0\x95\x32\xbd\xdc\xe2\x16\xd2\xe5\x78\x1d\x75\x8b\x9a\x90\xb7\x22\x9b\xd5\x12\x3d\xfc\x7f\xa3\x75\x58\xf5\x13\x39\x23\x5f\x95\x1b\xb8\xc1\xea\xae\x06\xe7\x5e\xab\xda\xf4\x06\xee\x72\x39\xae\x21\x49\xff\xb4\x7e\x03\x3a\xcf\x78\x00\xdb\xda\x04\x34\xd5\x6f\xca\xdc\xfb\x32\x40\xe5\xd8\xe6\xd0\xf4\x47\xfd\x1e\xac\xca\x95\xf6\x46\xa7\xd2\x1c\xb7\x92\x3c\x97\xf9\x66\xd7\xbf\x05\x86\x5f\xc4\x7f\x01\x00\x00\xff\xff\xb6\x5f\xf1\xa7\xac\x07\x00\x00")

func startupShBytes() ([]byte, error) {
	return bindataRead(
		_startupSh,
		"startup.sh",
	)
}

func startupSh() (*asset, error) {
	bytes, err := startupShBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "startup.sh", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
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
	"azuredeploy.json": azuredeployJson,
	"startup.sh":       startupSh,
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
	"azuredeploy.json": {azuredeployJson, map[string]*bintree{}},
	"startup.sh":       {startupSh, map[string]*bintree{}},
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