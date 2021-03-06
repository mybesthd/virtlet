// Code generated by go-bindata. DO NOT EDIT.
// sources:
// deploy/data/virtlet-ds.yaml
package tools

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

var _deployDataVirtletDsYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xd4\x5a\x5b\x6f\xdb\x38\x16\x7e\xcf\xaf\x38\x68\x80\xed\x0c\xb0\x8a\x93\x62\x67\xa7\x63\xec\x3e\xa4\xb1\x27\x63\x34\xb1\x0d\xe7\xd2\x79\x33\x68\xea\x58\xe6\x9a\x22\xb5\x24\xa5\xc4\xfb\xeb\x17\x24\x25\x45\x37\x3b\x4e\x9a\x18\x9d\xbc\xd4\xe5\xe5\xe3\xb9\x5f\x28\x06\x41\x70\x44\x12\x76\x8f\x4a\x33\x29\xfa\x40\x92\x44\xf7\xb2\xb3\xa3\x35\x13\x61\x1f\x06\x04\x63\x29\x6e\xd0\x1c\xc5\x68\x48\x48\x0c\xe9\x1f\x01\x08\x12\x63\x1f\x32\xa6\x0c\x47\x93\xff\x5f\x27\x84\x62\x1f\xd6\xe9\x02\x03\xbd\xd1\x06\xe3\x23\x9d\x20\xb5\xcb\x35\x72\xa4\x46\x2a\xfb\x1b\x20\x26\x86\xae\xae\xc8\x02\xb9\xf6\x03\x00\x2a\x15\x86\xd5\x21\x0d\xc6\x09\x27\x06\xf3\x3d\x95\xc3\xed\x5f\x93\x00\xfb\xc7\x6b\x90\x9d\xa0\x00\x05\x49\xf6\x6f\x25\xb5\x19\xa3\x79\x90\x6a\xdd\x07\xa3\x52\xcc\xc7\x43\xa1\xa7\x92\x33\xba\xe9\xc3\x05\x4f\xb5\x41\xf5\x3b\x53\xda\x7c\x63\x66\xf5\x87\xdf\x92\x2f\x3c\x76\x10\xd3\xd1\x00\x98\x76\x00\x60\x24\xfc\x74\xf6\x33\xa0\x20\x0b\x8e\x70\x7f\xad\xed\x88\x4e\x55\xc6\x32\x2c\xe8\x00\x2a\x85\x21\x4c\xa0\x02\x85\xda\x10\xf5\x04\xf7\x93\x91\xb0\x40\xa0\x2b\xa4\x6b\x0c\x7f\x06\x22\x42\xf8\xe9\xd3\xcf\x16\x24\x87\x34\x2b\x84\x54\x23\xc8\x25\x08\x8d\xc2\xa0\x02\x26\x80\x09\x56\x81\xad\xb0\x37\x1d\x0d\x6a\xac\x1d\xc3\x42\x4a\xa3\x8d\x22\x09\x24\x4a\x52\x0c\x53\x85\x20\x10\x43\x47\x29\x55\x48\x0c\x02\xb1\x58\x4b\x16\xc5\x24\xb1\xe8\x15\x95\x3e\x69\x3a\x07\xd4\xa8\x32\x46\xf1\x9c\x52\x99\x0a\x33\xae\xa9\xa5\x3c\x53\x0a\xbe\xb1\xea\x80\xfb\x5c\x02\x89\x0c\x35\x48\xe1\xb8\x11\x32\x44\x0d\x0f\xcc\xac\x00\x1f\x8d\x22\x33\xaf\xb6\x7f\x17\xd2\x72\x6a\xcd\xa1\xc8\x72\x69\x59\xdd\x3c\x29\xd9\xee\x3e\x6f\x8d\x02\x28\xfc\x6f\xca\x14\x86\x83\x54\x31\x11\xdd\xd0\x15\x86\x29\x67\x22\x1a\x45\x42\x96\xc3\xc3\x47\xa4\xa9\xb1\x56\x5f\xd9\xe9\x31\x6f\x72\x93\xbd\x45\x15\xeb\xfa\x74\xe0\x2d\x78\xf8\x98\x28\xd4\xd6\x67\x1a\xf3\x76\xc5\x1a\x37\xfd\x1a\x3b\x8d\x15\x00\x32\x41\x45\xac\x4f\xc0\x48\xb4\x26\x33\xc2\x53\x6c\xc1\x5a\xe0\x86\x6c\x2d\xdf\x17\x85\xde\xcb\x0d\xc7\x70\xbb\xc2\x86\x51\x00\x95\x09\x43\x5d\x00\x7c\xd4\xb0\xe4\xf8\x98\x49\x9e\xc6\x08\xa1\x62\x59\x69\x37\xc7\xd6\x12\xac\x66\x42\x5c\x92\x94\x1b\xa7\x7f\xa7\x35\x9e\x46\x4c\x40\xc8\x94\x33\x4c\x14\x3a\x55\xa8\xc1\xac\xc8\x93\x05\xbb\x7d\x4c\x39\xd9\xd9\xe3\xac\x69\x61\x08\x8b\x0d\x70\xb6\xb0\x67\xc3\xdf\x4a\x3f\xc0\x47\xa6\x4d\x61\x06\xd6\x5a\x8f\x0a\x2e\xbd\x7b\x27\x0a\x13\xa2\x30\xb0\xfa\x28\x45\xc1\x62\x12\x61\x1f\x62\xa6\x88\x30\x4c\xf7\xea\x31\x20\x9f\x9f\xa6\x9c\x17\x2e\x3c\x5a\x8e\xa5\x99\x2a\xb4\xde\x52\xae\xa2\x32\x8e\x89\x08\x9f\x24\x1c\x40\xaf\x7a\xdc\x89\x5e\x95\x53\x5e\x46\xd7\xd6\xbe\x75\x75\x83\x27\x72\xfd\x59\x07\x4f\x92\x0c\xbc\x8c\x74\x10\x32\x55\xd1\x5e\x6c\x37\x4f\x89\x59\xf5\xa1\x97\x4b\x33\xa8\x6f\x68\xe1\xaa\xb4\x6a\x16\xc7\x30\x90\xe2\xa3\x01\x12\x86\xf0\xc1\xa3\x29\x99\x90\x88\x38\xeb\x85\x2f\xcc\xcb\x9c\x49\x41\xf8\x87\xbf\x03\x33\xf0\xc0\x38\x07\x4e\xe8\xda\x1f\x0e\x28\x8c\xda\x6c\x21\xa9\x7a\x56\x71\x7e\x28\xe9\x1a\x95\x96\x74\xbd\x65\x53\x46\x94\xdd\xd8\xf3\x0b\x4f\x6a\x2b\x0b\x10\x2e\xa3\x2d\xbb\xad\xba\xab\xb3\xc7\xb0\x94\xca\x9b\x14\x13\x91\xb3\x29\x7f\x04\x67\x8b\x5e\x6e\x3a\x3d\xa7\x5b\xed\xed\xc6\xc5\x8f\x9a\x65\x14\x87\x66\x44\x05\x9c\x2d\x76\x1c\x1c\x34\x97\x94\x4c\x63\xb6\x65\x5b\x75\x26\x68\xc9\xc1\x12\xd9\x34\xc4\xee\x24\x65\x23\x26\x4d\x15\x33\x1b\xeb\xb6\xf8\x68\xaa\x4e\x9e\x28\x96\x31\x8e\x11\x86\xb5\xa0\x0d\x80\x22\x6b\x5b\xde\xd7\xbb\x2f\xc3\xf9\x78\x32\x18\xce\xc7\xe7\xd7\xc3\x0a\x8c\x8b\x1e\xbf\x2b\x19\xd7\x03\xc8\x92\x21\x0f\x67\xb8\x6c\x86\x95\x6a\xf2\xcf\xce\x1a\x93\x6e\x93\xe7\xd4\xa6\xce\x13\x2b\x71\x1b\xe5\x5b\xd4\xdc\x8f\x66\xb7\x57\xc3\xdb\xf9\x60\x74\x73\xfe\xe5\x6a\x38\xff\x7a\x7f\xfd\x3c\x49\x3e\xcd\x5c\x93\xe4\x2b\x6e\x3a\x28\xab\x09\x30\xf0\x8b\x1b\x4b\x5c\xa0\x0d\x99\xb6\xc9\x71\xbe\xce\xe2\xc6\xb4\x4c\xbc\x4f\x34\xe4\xd9\x24\xfa\x66\x36\x9a\xdc\xcf\x6f\xee\xa6\xd3\xc9\xec\xf6\x60\x64\x6b\xc5\x64\x36\xd7\x69\x92\x48\x65\x5e\x47\xf8\x60\xf2\x6d\x7c\x35\x39\x1f\xcc\xa7\xb3\xc9\xed\xe4\x62\x72\x75\x38\x99\xcb\x07\xc1\x25\x09\xe7\x89\x92\x46\x52\xc9\x5f\xc7\xc0\xd5\xe4\xf2\x6a\x78\x3f\x3c\x1c\xdd\x5c\x46\x1c\x33\x7c\x25\xb9\x17\xe7\x57\xa3\x8b\xc9\xfc\xe6\xee\xcb\x78\x78\x38\x43\xa1\x84\x33\x2a\x03\x9d\x2e\x04\xbe\xd0\x50\x46\xd7\xe7\x97\xc3\xf9\x6c\x78\x39\xfc\x73\x3a\xbf\x9d\x9d\x8f\x6f\xae\xce\x6f\x47\x93\xf1\xc1\x68\x77\x31\x7b\xae\x30\xc2\xc7\x64\x6e\x14\x11\x9a\xbb\xa4\xf5\x3a\xf9\xcf\xce\xbf\xcd\x07\xc3\xfb\xd1\xc5\xf0\xe6\x60\x1c\x28\xf2\x30\x0f\xd1\x56\xb9\xfa\x95\x4e\x9a\x87\xc4\xab\xc9\xe5\xe5\x68\x7c\x79\xf0\xb0\xc8\x65\x14\x31\xd1\x5c\xb2\xaf\xc5\x4f\xef\xe6\xd7\x93\xc1\x01\x3d\x94\x26\x69\x10\xcb\xf0\xa5\x2e\x6a\xd3\xa1\x33\x91\xc9\xc4\x8a\x7c\x76\x30\x7a\xf3\x82\x6e\xae\xa4\x34\xf3\x7a\xdd\xf7\x02\x39\x7b\x47\xad\x78\xe8\x4d\x17\x13\x7d\xe8\xa1\xa1\x45\xad\x91\x17\x44\x45\x33\x40\x5b\x8d\x40\x59\x87\xf9\x02\x6a\xef\x22\xfa\x18\x46\x02\x28\xd1\x08\x0f\xb6\x8f\xf8\x0f\x52\x03\x5c\x52\xc2\xcb\xda\xdd\x21\xd8\xd9\x07\x22\x8c\x6d\x18\x6c\x53\xca\x0c\x08\x69\x40\x2e\x97\x8c\x32\xc2\xf9\x06\x48\x46\x18\x77\x8d\xab\x14\xf8\x06\x35\x7a\xce\xc8\x3e\xe5\x79\xb5\x46\xd3\x1b\xdd\x5b\xea\x1e\x8d\x94\x4c\x93\x56\x85\xd6\x18\xae\x6f\xb5\xa5\x5d\x2c\xc3\x94\xd7\xbc\xdf\x6f\x6c\x8f\x2b\x24\xe1\x44\xf0\x4d\x4b\xd9\x55\x48\xdb\x82\xb7\xb0\x1a\x83\x7b\x01\xbd\x77\x8f\xd0\xee\x44\xbe\xaf\xf4\xed\xde\xdd\x34\x4e\xd8\x62\xb4\xed\xdd\xb6\xfd\x78\x66\x77\x60\xfb\x12\x34\xba\x62\xda\xb6\xdb\xe4\x32\x72\x7d\x2c\x2b\x3b\xd4\x15\x2a\x84\x05\x52\xe2\x6e\x57\xcc\x0a\xd5\x03\xd3\x58\x76\xad\x4e\x54\x89\x92\x61\x4a\x11\x50\x29\xa9\xaa\x90\x9c\xad\x11\xcc\x8a\x55\x0c\xf0\x18\xee\xf2\x1b\x1b\x69\x1b\xd9\x20\xbf\x5a\xa1\x2b\xa2\x42\xcc\x60\xc9\x38\xc2\x47\x2f\x03\x19\xf5\xb2\x58\xf7\xc8\x32\xfc\xf5\x97\xc5\x62\x11\x7c\xc6\xdf\x7e\x0d\xce\xce\xf0\xd7\xe0\xb7\x5f\xfe\x79\x16\x9c\x7e\xfa\xc7\xa7\x53\x42\x4f\x4f\x4f\x4f\x3f\xf5\x28\x53\x4a\xea\x20\x8b\xe7\xa7\x27\x5c\x46\x1f\xfb\x30\x96\xa0\x53\xba\xf2\x88\x52\x95\xdd\xf7\xa6\xdd\x18\xc5\x3a\xd8\xde\x91\x55\x48\x69\xf7\x71\xb9\x30\x9f\xdf\xdd\x56\xda\x4b\x3a\xab\xd7\xf4\x46\xd6\x53\x98\x40\xad\xa7\x4a\x2e\xb0\xba\x05\x1f\x9f\xee\xfa\xfc\x5f\x2b\xa4\x78\x12\x7b\x0b\x26\x7a\x95\x90\xe2\x47\x03\xda\x18\xd0\x92\x12\x03\x01\xdc\x8d\x47\x7f\xf6\x9b\x06\xd8\xab\x1a\x5c\xa0\x24\xfc\xcb\x72\xd6\x13\x29\xe7\x8d\x60\xdc\x79\x63\xf1\xa3\x07\xe3\x7d\xa2\xec\xe1\xc3\xd1\xb1\x0f\xa6\xee\x3a\xaa\x1a\xa9\x81\x28\x2c\xaf\x00\x61\xb1\x01\x9d\x26\xa8\x62\xb6\x2d\x90\xfd\x68\x41\xfe\x70\xb7\x11\x05\xee\x56\xd5\xfc\x50\xc1\xbb\x8e\x92\x6a\x47\x83\x75\x73\x77\xab\xa6\x04\x1a\xd4\xe5\x05\x5b\x7e\xb3\xd6\xf3\xa6\xdb\xb3\xcb\x5a\x07\xed\x71\x7b\xd7\xcd\x77\x7e\x48\x2f\x91\x61\xdb\x64\x2c\xaa\x9d\xe8\xbc\x05\xdc\x47\xd2\xaf\x8f\xd7\xd5\x15\x1d\x95\x62\x93\x52\x37\x1c\xd8\xdf\x41\xa5\x37\x6b\x27\x00\xc7\xcd\xf3\xb4\xd4\xa4\x71\x5c\xa4\xd6\xa5\xcb\x4a\x24\x12\x52\x1b\x46\x21\x49\x55\x22\x35\xbe\x47\x96\x11\x68\x76\xde\xbd\x16\x76\xe7\xd6\xbd\x42\x33\x7f\xa5\xfc\x64\xf9\xcc\xb5\xef\x2e\x67\x77\x24\xa4\x8a\x05\xbd\xdf\xf5\x7a\x16\xeb\x97\x96\xed\x6f\x11\xcc\xfe\x90\xda\xdc\xca\x8b\xc6\xb7\xb1\xef\x0f\x67\x6f\xe3\xa2\x6f\x1b\x4e\xb6\xf3\xfa\xb2\x84\xb2\x2d\xf1\xed\x4e\x99\x5e\xa3\x95\x8f\x50\x16\xb5\x52\x61\xdb\x30\xb0\x92\xda\x80\x22\x0f\xe0\x2f\x54\x80\x50\x8a\x5a\x97\xf6\xe8\x3e\x59\x5a\xfc\xaa\x63\xb5\x29\x6c\x72\xb3\x73\x63\x77\xef\xd7\xd1\xf9\xed\x44\xe9\xaa\x10\xba\xc4\xb4\x13\xa4\x96\xfe\x5b\x15\xc1\xce\xad\xd5\xaa\xa7\x59\x07\x1d\xc3\xed\x64\x30\xe9\x43\xe8\xea\x2d\xdb\x60\x50\x19\x62\xfe\x05\x07\x7c\x0e\x75\x15\xa3\xb5\x12\xd7\xe8\x3c\x6d\x5c\x31\xed\xeb\xae\xbc\x5a\x82\x8b\xd9\xc8\xf6\x39\x8f\x1b\x60\x42\x1b\xc2\x7d\x46\xb0\x45\x65\xf5\x40\x26\xbc\x2a\x7d\xa1\x56\x7e\x14\x3e\xd9\x87\x95\x5d\x1f\x8e\xb6\x7c\x7b\x7a\x16\xaf\x2b\x4a\x74\xc5\x88\xbd\x80\x9a\xce\xde\x15\x02\x9e\x07\xaa\x44\x85\xe6\xc7\xb0\x9d\x9b\xbf\xa3\xaa\xd9\xb3\xa6\xd9\x4b\x08\x9d\x11\x69\x6b\x3c\xda\x07\xb2\xa9\x98\xda\x37\xb8\x7d\xe4\x59\x16\x33\xd5\x78\xda\x15\x87\xf7\x02\xdb\xa9\xe5\x97\x80\x75\x15\xb2\xbb\xca\xd8\xbd\xa8\xeb\x10\x7b\xa3\x06\xdb\x8b\xae\x7a\xa1\xd3\x5d\x24\x05\x4f\x17\xa1\xfd\x6d\x29\x36\xf0\x85\x62\x67\x8d\xb8\xbb\x92\x6c\x3e\x2f\x52\x0b\x42\x4f\x48\x6a\x56\x52\xb1\xff\xb9\x35\x27\xeb\xcf\xfa\x84\xc9\x5e\x76\xb6\x40\x43\x8a\x87\x47\xf9\xcb\x9b\x99\xe4\xf8\x85\x89\x90\x89\x68\xc7\x0b\x24\x25\x39\xe6\x37\xb8\x24\x61\x97\x36\xa8\xef\x38\xe9\x08\xa0\x75\x46\x0b\x52\xa7\x0b\xdb\x69\xeb\xfe\x51\x90\xaf\xbe\xa9\x3d\x75\xd9\xff\x15\x94\x95\x40\xfb\xbc\x97\xc9\xe4\x15\x8f\xaf\x94\xcd\x4a\x76\x7d\x50\xca\x24\xcf\xcd\x01\x7c\xf8\xe0\x7e\x28\xd4\x32\x55\x14\xcb\xf1\xf2\xd9\x8f\xce\x07\xdc\xe3\x1c\xf7\x3b\x43\xb5\x78\x5a\xe7\x2e\xb3\xf2\xff\x44\x68\xde\x42\xcb\x1d\x3c\x96\xe4\x04\xb6\x88\x46\x55\xf0\xd4\xe0\x28\xe7\xa7\xc6\x4d\x83\x97\x92\x7a\x4f\xae\xfd\x97\x33\xed\x7f\x3c\x10\x43\x57\xef\xc4\x41\xe1\x3e\xa9\x46\x65\x67\xbe\x9b\x91\xc0\xf6\x20\xca\x07\x93\x06\x53\xef\xea\x69\x45\xfa\xb1\x06\x11\x2c\xf2\x65\x6f\xe8\x76\x2d\x55\x57\xfd\xef\x25\xe0\x97\x79\x45\xe7\x61\xbd\x2f\xf4\xbd\x19\xbf\x6f\x28\x8a\x9f\x94\xfc\x0e\xf2\xd9\x66\x48\x7f\x91\x30\x15\x50\x15\x6e\x37\x7a\x92\x30\x7c\x34\x28\xdc\x63\xba\x1c\xb3\xcb\x11\x52\x6d\x64\x5c\x0c\x86\xe8\x5e\xfd\xe5\xa9\xa8\xe2\x0b\x79\x70\x6a\x1f\x53\x74\xc3\xeb\xcf\xba\x03\x3d\x9f\x75\x79\x2c\x26\x49\xc2\x44\xa4\xab\x13\xa5\x85\x16\x33\x95\x23\xcb\x58\xf2\xee\x7e\x58\x93\xe7\xdb\x9b\x97\x85\x7d\x5b\x93\x6a\xbc\x2e\xea\x04\x7c\x45\x76\xfb\x7f\x00\x00\x00\xff\xff\x84\x6e\x8c\xa4\xb6\x2c\x00\x00")

func deployDataVirtletDsYamlBytes() ([]byte, error) {
	return bindataRead(
		_deployDataVirtletDsYaml,
		"deploy/data/virtlet-ds.yaml",
	)
}

func deployDataVirtletDsYaml() (*asset, error) {
	bytes, err := deployDataVirtletDsYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "deploy/data/virtlet-ds.yaml", size: 11446, mode: os.FileMode(420), modTime: time.Unix(1522279343, 0)}
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
	"deploy/data/virtlet-ds.yaml": deployDataVirtletDsYaml,
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
	"deploy": &bintree{nil, map[string]*bintree{
		"data": &bintree{nil, map[string]*bintree{
			"virtlet-ds.yaml": &bintree{deployDataVirtletDsYaml, map[string]*bintree{}},
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

