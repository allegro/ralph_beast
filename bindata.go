// Code generated by go-bindata.
// sources:
// bundled_scripts/idrac.py
// bundled_scripts/idrac.toml
// bundled_scripts/ilo.py
// bundled_scripts/ilo.toml
// DO NOT EDIT!

package main

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

var _idracPy = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xcc\x3a\x7b\x6f\xe2\xb8\xf6\xff\xf3\x29\xfc\x63\xb5\x4a\xd0\xb6\x01\xfa\x98\x9d\xe1\x37\xcc\x8a\xe1\xb1\x8b\x76\xa0\xbd\x85\xee\xce\x55\x6f\x15\xb9\x89\x29\xd9\x09\x09\x6b\x27\xa5\xdc\xaa\xdf\xfd\x1e\xdb\x79\x38\x2f\x28\xed\x1d\xe9\x52\xa9\x71\xe2\xf3\xf2\x79\xdb\xc9\x0f\xff\xd7\x0c\x19\x6d\xde\x39\x5e\x93\x78\x0f\x68\xbd\x0d\x96\xbe\x57\xab\x39\xab\xb5\x4f\x03\xf4\x17\x83\x9b\x68\xec\xb3\x78\x44\x49\x3c\x62\xdb\xe4\x61\x18\x3a\x76\x6d\x41\xfd\x15\x7a\x5c\xb9\x06\x09\x28\x21\x28\x9a\x1a\xba\x64\x45\xbc\x60\xce\x1f\x61\x86\x86\xf3\x5a\x4a\xe9\xef\x90\xb0\x80\x49\xc4\xf8\xce\x58\x63\xeb\x1b\xbe\x27\xcc\x08\xa9\xeb\x3a\x77\xa7\x06\x79\xb4\xc8\x3a\x70\x7c\x8f\xc5\x34\xc7\x1e\x23\x56\x48\xc9\x95\xc4\xf9\x13\x53\xcf\xf1\xee\x6b\xb5\x59\xff\xb7\xe1\xa4\x87\xba\xa8\xbe\x0c\x82\x75\xa7\xd9\x64\xd6\x92\xac\x30\x33\xec\x55\xb0\x30\x7c\x7a\xdf\xdc\xdc\x91\x55\x73\xc3\x2c\x67\xd5\x6c\x37\xe1\xff\xb1\x84\x68\x9e\xd4\x6b\x5f\x27\x5f\xa6\x33\x73\xc6\xb1\x9f\x22\xf4\xcd\x66\x63\x6c\x4e\x05\xe2\x49\xab\x75\xda\x6c\x9d\x37\x99\x8f\xd7\xc7\xa0\x2c\xe2\xfa\x6b\xf2\x1c\x63\xfd\x39\x1b\x4e\x55\xc4\x98\x2f\x68\x83\x23\x48\xd6\x8c\x13\x39\x6b\xb6\x3e\x80\xb2\xc3\x15\xa1\x98\xaf\x49\x21\x31\xe9\x95\xd2\xc8\xcb\xbe\xc2\x1e\xc8\x2e\xae\xc6\x23\xb3\x13\x02\xd3\xb6\xf9\xb9\x37\x1b\x96\x92\x20\xae\x6b\x58\xfe\xaa\x7a\xf9\xcd\x1f\x19\x10\xfa\x01\xfd\x4a\x3c\x42\x1d\x0b\x09\xf2\xc7\x16\xc5\x8b\x80\xd8\x88\xaf\x01\xad\x08\x63\x60\x97\xda\xec\xa2\x77\x69\x0e\xa7\xd7\x13\x29\xb2\x39\x1f\x4e\x2e\xbf\xf4\xe6\x9c\xb1\xa6\x69\x1f\x7f\x81\x35\xa3\x07\x42\x19\x2c\xae\x5b\x6f\x1b\xad\xfa\x2f\x9f\x6a\x1f\x59\x67\x18\xe9\x8c\x7b\x88\xc7\x3a\xac\x5b\x7f\xa1\x92\xeb\x11\xc6\x86\xe1\x6e\xde\xae\xa5\xfa\x7d\xdf\xc4\xb6\x4d\x41\x5a\xf0\x89\x14\x19\xd6\x53\x40\xdf\xa3\xda\x14\x99\x14\x71\xf7\x9b\xb6\xfe\xa9\x86\x10\xac\xfc\x37\x82\x6d\x42\xf9\x0d\xdc\xc2\x2a\x3a\x3d\x8b\x4f\x23\xd6\x59\x85\x2c\xb8\xf6\x60\x92\x05\xd8\xb3\xbb\xf5\x80\x86\xa4\xfe\xe9\x60\x46\xcd\x61\x34\x26\x1f\x9b\x29\x7d\x85\xe1\xdc\xaf\x64\xf6\x04\x6b\x05\xab\xf2\x10\x35\x21\xe0\x9e\x25\x85\xb9\x9f\x60\xc3\x7c\xe7\x8a\x30\x3f\xa4\x16\xb9\xbe\x1a\x57\x13\xa2\x11\x90\x20\x91\xc3\x52\x64\x99\x48\x37\x1a\x0f\x2a\x29\xf1\x64\xd2\x79\xe2\xff\x23\x69\x12\x14\x85\xcc\x15\x59\xbb\xdb\x58\xcc\x58\xaf\xd2\xec\x2f\x52\xa0\xea\x24\x4d\xea\xbb\xa4\x89\x3d\xdf\xdb\xae\xfc\x90\x45\x4a\x8c\x88\x49\x96\xcd\x02\xcf\x48\x35\x33\xe2\x12\x2b\xf0\xe9\x8c\x04\x8a\x2c\xea\x0c\x9a\xe2\x15\xe9\xd6\x4d\x13\xe2\xcd\x83\x21\x83\x04\xc7\xf5\xc5\xa2\xf9\x44\x5f\x31\x42\xca\xb2\x84\xc1\xc7\xa6\xea\x50\xe0\x5d\x9f\x7d\x7b\x9b\x48\x44\xbc\x4e\xe2\x0b\x39\x71\x2e\x20\x85\xae\x9c\x7f\x93\xa1\xe2\x37\x39\x90\x09\x7e\x8c\xb2\x35\xe3\x7e\xf1\x68\x92\xe8\x2e\x11\x51\x85\x48\xa4\xcc\x33\xe5\x22\x4a\xa9\xf8\x28\x8e\xfb\x4f\x35\xc8\x0e\xb5\x51\xdf\x1c\x4f\x47\x17\xe6\xf0\xeb\xe5\xd5\x70\x36\x1b\x5f\xf0\x94\x47\x09\x4f\x4e\x6b\xc7\x25\x3a\xd5\xf4\x9b\xd6\xf1\x87\xdb\x9f\x1a\xc7\xf2\xaa\x35\x6a\xb5\x49\xaf\x6f\x02\xf8\x68\xfc\xd5\xfc\xfc\xa5\xd7\xff\xfd\xcb\x78\x36\x07\xb4\x1b\x21\x81\x76\xde\x82\xbf\x33\xed\x08\x69\xa7\xa7\xe7\xad\x77\x23\x3e\x6a\xb5\x3e\xbc\xff\xf9\x9d\x1c\xf1\x5f\x32\xea\xf3\xd1\x49\xeb\xac\x7d\x7e\xca\x47\xed\xb3\x0f\xed\x13\x98\x95\x94\x5a\xe0\x1a\x92\xd2\x68\x38\x82\x9f\x80\xe8\x8d\x3e\x9c\x48\xfc\x93\xd6\x7b\x39\x1a\x0c\x7b\x83\x93\xbe\x84\xeb\x0d\xce\x06\x80\x7f\x5b\x9b\x0d\xaf\xc6\xbd\x2f\x25\x02\x4e\x7d\x8f\x00\x28\x07\x9f\xfa\x01\xea\x3d\x60\xc7\xc5\x77\x2e\xe1\x0f\xbe\x3e\x8a\x3f\x3e\x3c\xe6\x3f\x3e\xb8\xb9\xf6\xbe\x79\xfe\xc6\xbb\x4d\xc5\x97\x4b\x90\x42\x72\x1a\xb3\x35\xb1\x9c\x85\x43\x6c\x0e\xf2\xcf\xdf\xdb\xad\xfe\x40\xc8\x7a\x72\x7a\x76\xfe\xee\xe7\xf7\x1f\x5a\x92\x99\x27\x78\x40\xf0\x7f\x26\x68\xe4\xb8\x2e\x64\xf2\xcf\x5b\x74\x61\x0c\x8d\x89\x21\x64\xae\xd5\x6c\xb2\x40\x9e\x4f\x57\xd8\x05\xd7\x30\x57\xd8\x32\xa3\xb0\xd0\x95\x71\xa3\x23\x58\x2b\x4f\x60\x71\xca\x9d\x11\xae\xd7\x84\xea\x0d\x83\x42\x80\x80\x73\xeb\x9a\x58\x48\x07\x6c\xc7\xf1\x28\x09\x42\xea\xa9\x08\xc0\xb8\xba\xdc\xdb\x0e\xe3\xea\x31\x37\xb2\xaa\x33\xbd\xbc\xda\x83\x5f\xd4\x2c\x17\x83\x2c\x63\x9b\x62\x6b\x48\xa9\x4f\xf5\x61\xdc\x2a\x44\x22\xaf\xb1\x60\x16\xc1\x0d\xae\x7a\x7d\xdd\xbf\xfb\x0b\xe2\x09\xe6\x05\x00\x5f\xbf\x69\x3a\x9e\x13\x98\xa6\x0e\x11\xb9\x38\x42\x4b\x9f\x05\x47\x28\x64\x84\x1e\x09\xfc\x8d\x4f\xed\x88\x1c\xff\x71\x20\x83\xc3\x80\x0e\xf8\x25\x3b\xc1\xd1\x60\x82\x5f\xb2\x13\x31\x25\x98\x8c\x87\xa9\x04\x34\xf4\x4c\x88\x00\x88\x30\x3b\x12\x42\x48\x6c\xf2\x74\x71\x84\xe2\x44\xd1\xd5\xa8\xef\x07\x4d\x1b\xf2\x88\xa6\x08\x94\x4d\xe0\x71\xf3\xc3\x20\x01\x3e\x3d\xcb\xb0\xad\x1b\x0b\x6e\xe2\x40\x4f\x64\x6f\x24\xd8\xf7\xbc\xe0\x43\xe0\xda\x26\x4f\xb8\x5c\x74\xb8\x18\xfc\x5f\x5b\x4f\xa1\xa2\xda\x0f\xd3\x55\xd5\x3f\x66\x91\xa0\x48\xcb\xcb\x0a\xd0\x95\x8d\x99\x01\x79\x9e\x3a\x6b\x5d\x6b\x6a\x0d\xf4\x13\x82\x0b\xfc\x57\x56\x9a\xc1\xcd\x2e\xab\x9b\xbd\xcd\x82\x72\x69\xbb\xd9\x85\x64\x01\x12\x0d\xc6\x83\x3c\xab\x34\xd5\x75\x4f\xce\xcf\xd3\xd9\x54\x03\x91\x13\x0f\xe7\x06\x74\x5b\x52\x91\x26\x64\x3e\xdb\xe4\xa5\x25\xbb\x6a\x2e\xae\x96\xb7\x81\x56\xb4\x41\x4e\x08\xa9\xe2\x6e\x74\x55\x64\x68\x28\xae\x9a\xb2\x94\x7e\xc2\x75\x11\xa3\x2a\x3e\x51\xaf\xd7\xe7\x74\x8b\x02\xa8\xfb\x80\x90\xe9\xdf\xf8\xc3\x35\xb0\x17\xce\x12\xf2\xf2\x87\xb8\xac\xe8\x0e\x33\xc7\x4a\x08\xe0\x30\x58\x82\x3a\x1c\x4b\x14\x0a\x03\x92\x58\x00\x9e\x18\x2c\x71\x80\x36\x04\xd9\xbe\xa7\x41\xff\x0f\x8a\x84\x9e\xde\xdb\x02\x13\xc6\xdb\x3d\xe4\x78\x72\x91\x30\x4e\xe5\x87\xf4\x82\x1e\x20\xbf\xd8\x60\x1c\x34\x9b\x7d\x41\x16\xa1\x01\x64\x2f\x20\x4d\x0c\xd4\x03\xf4\x85\xef\xba\xfe\x86\x4b\x12\x67\x05\xb4\x81\x74\x05\x77\xc7\x5c\xfc\x84\x92\x10\x51\x88\x86\x96\xa2\x04\x22\x7c\x8f\x1d\xcf\x50\x97\x9d\x1a\x4c\x14\x96\x38\xc9\xc0\x82\x0b\x46\xca\xea\x1f\xc4\xc3\x45\xe5\xc7\xba\xe8\xea\x49\x78\x1f\x65\x03\x3a\x67\x46\xe8\x7c\x9d\xc5\xb6\x3b\xc2\x2e\xcb\x91\x91\x22\xb3\xee\x53\xe6\x29\xff\x69\x7d\xdf\x0b\x40\xdb\xc7\xf3\xed\x9a\x68\x1d\xa4\xe1\xf5\xda\x8d\x54\x2f\x7a\xe1\x9f\xa0\x85\xf9\x7f\x6b\x89\x29\x23\x41\xf7\x7a\x3e\x3a\x7e\xaf\x65\x69\x3f\x97\xb9\xac\xc3\x53\x3b\x6c\xb3\x0c\xff\x5b\x27\x03\x0d\x13\xd4\x80\x96\x2b\x08\x19\xa4\x1d\x1b\xc2\xba\x8b\xce\x5a\xed\x4e\x41\x30\x8a\x1d\x46\xd4\xfc\x5a\xef\x71\xe5\x13\x3e\x36\x60\x17\x26\xcc\x2a\x92\x1d\x8f\x5f\x04\x86\x8e\xb5\x62\xd4\x1b\xb5\x9d\x94\x0a\xbc\xea\x33\x58\xa8\x98\xeb\x20\x68\xb4\x60\xb7\x47\x3a\xe8\xe9\xf9\x5f\x5e\x94\xf6\xf9\x4d\xbd\x34\xcb\x24\x3c\x8c\x80\x3c\x06\x49\x38\x1c\x15\x80\x72\xa6\x4a\x25\x14\x0b\x62\xe6\x1a\xc3\xe2\x60\x17\xf3\xc4\x9e\x79\xff\xd2\x84\xeb\x08\x87\x6e\x90\xc6\x6e\x37\xda\x24\x16\x50\x5d\x47\x14\x84\x9b\xdb\xfc\x84\x27\xf4\x1b\xe7\x0d\x29\x62\xc3\x58\x38\x90\xe9\x15\xae\x19\xa3\x29\x98\x59\x93\xe4\x78\x71\x00\x93\xd3\x83\x00\xa2\x28\xbd\x73\x3c\x95\x84\xe1\x04\x10\x6f\xf0\x5c\x6f\xdc\x1e\x68\x12\x2d\x63\x92\xc4\x0a\x47\x70\xc3\x84\x81\x22\x3e\xfc\xa1\xb6\xd3\x34\xb1\x49\x90\x76\xa4\x19\x7f\xf9\x8e\xa7\x2b\x6b\x69\x94\x58\xaa\xc2\x50\x51\x16\x96\x6a\x8c\xda\x17\xf3\x9e\x04\x26\xa4\x06\x62\xf2\xf4\xa3\x3b\x7c\x41\xa6\x2c\x17\x34\x4a\x89\xe2\x3c\xa2\x8b\x32\x53\x86\x5a\x74\xb5\x41\x7f\x3c\x31\x67\x5b\x16\x90\xd5\x1f\x0e\xd9\x44\x9d\x8b\xd8\xfe\x99\x5e\x1b\x70\xb3\x7b\xec\x1f\x51\x3d\x87\x21\x13\xcf\xdf\x62\xf7\x1d\x79\xcf\x73\xd2\x15\xc7\xfa\x82\x67\x63\x80\x67\x70\xcd\xa3\x17\xd4\x17\x79\xda\x51\xee\x01\x3f\x69\x28\x3e\x83\x4a\x9c\x3e\x8c\xa5\x96\x4f\xe2\x1e\xcc\x82\xa8\xe4\x7d\x1b\xd7\x85\xf0\x3f\xec\xba\xfa\xdf\x72\x36\xce\x14\x12\x28\x75\xba\x62\xfc\x8f\x3d\x00\x01\xb8\x00\xf2\x3e\xdb\x40\x06\x06\x67\x83\x4a\x91\x33\x42\x1c\xfd\x50\xff\x21\x7c\x80\x69\x9a\xf3\xb4\x15\xf8\xa4\x2b\xea\x3d\x64\x3a\xd0\x55\x65\x58\x47\xd2\xdc\xb4\x6e\x65\xb8\x14\x33\xc6\xd3\xb3\x82\x9b\xac\x1a\x69\x13\xec\x85\x0b\x6c\x81\xab\x10\xaa\xe5\x83\x5e\xb8\x8e\x21\x1b\x91\xb4\x63\xad\x43\x36\xb3\x8c\xfa\x11\x94\x90\x1c\xc2\x1b\xc4\xe0\x2b\xdd\xcd\x5f\x49\xdb\x72\xf8\x2c\xfe\x43\x42\x75\x30\x28\x29\x5c\xdd\x11\x59\xc7\x2a\x64\xa8\xe4\xdd\x5f\x42\x22\x76\xd8\x8c\xd0\x07\xc7\x22\x73\x7c\x1f\xcb\x91\xe5\x1f\x5b\x3f\xcb\x90\xfb\x02\xd8\x35\xbf\xa3\x51\xfc\x42\xd8\xf5\x46\xcb\xa0\x69\xb7\x20\x69\xe6\x89\xda\xfd\x4b\x14\x35\x64\x95\xfd\x00\x61\x6f\x09\xdb\xe9\xb8\x7f\x48\xcc\x46\xe0\xaf\x0a\xd8\x18\xf7\x3b\x47\x6b\x46\x35\x69\x59\xe1\x69\x5e\xba\x82\x08\xbb\x6c\x20\xa7\xd6\x09\xe8\xb6\x93\xeb\x6f\xad\xc4\x89\x0e\x75\xe2\x7e\x08\xe1\xee\x05\xb0\x05\x8f\x0e\x44\xca\x1d\x3a\x2d\x7b\x62\x07\x86\x7a\x01\x38\xd8\x5d\x18\x10\x59\x3e\x32\x18\x16\x74\x3b\x8e\x17\x92\x7c\xa3\x02\x62\xee\x01\x94\x0b\xa9\xdc\xaf\x66\xaa\x28\xdc\xdf\x74\xde\xdd\x72\x4d\x95\x9d\x1f\xec\xe7\x94\x5a\xc0\x80\x7e\x0c\xfa\xd0\x94\x45\x71\x47\x4b\x98\xea\xda\x6b\xea\x5b\xf0\x10\x0a\xdc\x5b\xfc\xba\x7f\x79\x7d\x88\x5f\x47\xe0\xaf\xf2\xeb\x18\xf7\xbb\x57\x21\x9e\x04\x0e\xf7\x68\x51\x33\x72\x2e\xbc\x2f\xed\x96\x24\x3a\xa9\xfb\x5c\x70\x00\x2d\xf2\xc8\xad\xe1\x95\xf4\x2e\x3b\x63\x86\xff\x2a\xa5\x18\x7b\xfc\x4c\xd3\x22\xe3\x41\x3e\x60\x84\x3a\xb2\x55\x88\x41\xbb\x1f\xe8\x9a\xa1\x35\x6e\x8e\xdb\xb7\x95\x5d\xaa\x0c\x2d\xfd\x0f\xec\x86\x32\xae\x8e\xa0\x6a\x81\xf4\x62\xdc\xd8\xe3\xd1\x91\xee\x63\x5f\xce\xee\x42\x34\x58\x26\x61\x50\x90\xb9\x16\x5e\x97\x28\xa6\x22\xe5\x5f\x2c\x2e\x63\xdf\xef\x0b\x92\x3b\xcb\x5f\x6e\x32\xdb\x1b\x88\x9b\x1c\x00\x5b\x13\x62\xbf\x49\xcc\x09\x7e\xec\xbb\xbe\xf5\x6d\x26\x28\x1d\x24\x9d\xf0\x14\xc1\x1c\xae\xb9\xb9\x05\x5e\x39\xee\x16\x26\x5f\x99\x65\x2f\xaf\x47\x92\xc2\x0b\xdb\x05\xc1\xd4\xc5\x77\xe0\xe9\x05\x4d\x3d\x37\x8a\x85\x37\x93\x9e\x56\x64\xe5\xd3\xed\x5b\x52\xd3\x44\x50\x38\x24\x3b\xa5\x18\xaf\x4a\x50\x0a\xfa\x77\xcf\x51\x42\x6d\x37\x09\xd0\x53\xb9\xd2\x35\xd1\xba\x56\x6f\x7b\x5e\x9f\x3a\x76\xb5\xaf\xfb\x7c\xe2\x8d\x9c\xcb\x3a\xd6\x7d\x2c\xf3\x5e\xc9\xa0\x32\xbf\x29\x44\x67\x9c\xc0\x41\x91\xf9\xf6\xb4\xf0\xdf\x4b\x07\xcf\xa2\xa8\xc9\x87\x4a\x6d\x8b\x5f\xd7\x11\x3d\x57\xe5\x8e\x10\x14\x0a\x1a\x74\xdb\xd2\xfd\x6e\xd5\x40\xb5\x1d\xf6\xed\x4d\x2d\xc4\xe5\x72\xcb\x1c\x0b\xbb\x03\x20\x74\x48\xb4\xe6\xf1\x5e\x15\xb3\x05\x22\xff\xab\xdd\x85\x12\x6f\xbb\xfa\xe4\x83\x03\xb6\xa2\x05\x49\xab\x9c\x38\x6c\xaa\x4e\x23\xaa\x60\x65\x9b\xd2\x57\x6e\x48\x5f\xb6\x1f\x4d\x46\x3c\x9e\x61\x57\x6f\xde\x6d\x03\xb1\x19\xa9\xe4\xad\x71\xbe\x5a\x45\x40\x8f\xbd\xcf\x1c\xbf\xa8\x9e\x94\x4f\xe9\xa6\xf7\x20\x3e\x82\x82\x6c\x44\x76\x30\xda\xdd\x09\x29\xe9\xab\xa0\xd9\xd2\x87\x45\x1d\x61\x7e\xe4\xae\x3e\x89\xd5\xcb\x4f\x48\x5b\xc5\xe4\x8a\x9a\xa8\xdd\x3a\x39\xcb\x5e\x76\xe7\xd8\xcc\xc6\xbb\x93\xd3\x9d\xe0\xaf\x3e\x51\xf9\x6b\xda\xae\x2e\xa2\xe4\x45\x4c\x49\x57\x56\x06\x95\x74\x40\xe5\x5e\xbb\xb7\x29\x59\x38\x77\x60\xf1\x25\xf6\x3c\xa0\x6f\x61\x6a\xbf\x2d\xf1\xf5\xc7\x03\xc2\x8f\x3e\x0e\xca\x7a\x2a\xd2\xeb\x52\x5e\x86\xc2\x77\xce\x77\x21\x23\xb6\xe9\x88\x43\x3d\x46\x02\xfd\x6d\x49\x50\x78\xc1\x6b\xb2\x5f\x7d\x40\x98\x05\xde\x25\xbe\x78\xa9\x0e\x3a\xd8\x91\x6b\xc2\xc4\x28\xb2\xb1\x16\x1f\x33\x09\xce\x86\xeb\x6f\xf8\x2b\xe4\xbd\xdb\xf2\xc0\xe2\x47\xf4\xc5\xcf\x08\x0c\x46\x30\xb5\x96\x65\x47\x88\x07\xa6\xc9\xfa\xe8\x1f\x83\x41\xfe\x14\x50\xae\x66\xc7\x5b\x16\x21\xd9\x1e\xe9\xd7\x51\x41\x34\xc5\x8b\x56\x81\x61\xdc\x53\x3f\x5c\xeb\xed\x0c\x3d\x15\x0e\x14\x14\xdb\x79\x0f\xf5\x18\xcc\xc0\xb6\xad\x2b\x24\x1a\x2f\x4c\x7c\x0a\x0a\x44\xb1\x72\x57\x91\x2f\xc4\xf5\x65\x01\x2e\xa3\xd5\x16\x91\x51\x7d\x4a\xaf\xcc\x83\x7a\x76\x1e\xeb\xd7\xa4\x33\x64\x4f\xc9\xf6\x9e\x2a\xd6\x22\xfd\x66\x60\x52\xad\x2a\xfc\x6f\xb4\x0c\x8c\x38\xda\xcc\x9e\xf6\x08\x7b\x26\xc7\x3c\x31\xf7\xca\x83\x9f\x98\x75\x0a\x50\xc1\x37\x05\x10\x4c\xd3\x5b\xb9\x64\xb1\x0d\x4a\xd6\x5a\xb6\x8f\x4b\x16\x29\x26\xab\x56\x27\x26\xe5\xb2\xc4\x50\x1a\x80\xb7\x9b\x31\xf1\xb2\xde\x33\xa6\x2d\xe6\x2a\x48\x8b\x39\x41\x59\x8c\x64\x0e\x2a\xe4\xf6\x98\xcb\xde\xac\x1f\xb3\x2c\x02\x56\xf0\x2f\x02\x0a\x61\x8a\x8f\x55\x77\x55\x28\x44\x2e\xcb\x2c\xec\xe9\x3b\x3e\x05\x01\x91\xe4\x47\x20\x50\x1f\xea\xbb\xde\x98\x4c\x7d\x34\xbe\x44\xf1\x77\x33\xfc\x8d\x3c\x50\x46\x4b\xcc\xd0\x1d\x21\x1e\x37\xf0\x83\x63\x93\xe4\x85\x29\xd0\x95\xdf\x90\xbc\x80\x6e\xfa\x4d\x44\xfa\x0e\x76\x17\xe1\x17\x0b\xac\x10\x4e\xbe\x5b\xa9\x26\xac\xda\x0b\x34\x2d\x3f\xb3\x29\x55\x5d\x49\x98\xef\x4b\x0d\x51\xa0\xf1\xce\x8b\x7f\x8f\x6c\xd8\xe1\x6a\xcd\x74\x05\x9e\x7f\x1e\x51\x73\xf8\x57\x3c\x7c\xfd\xa6\xc9\x17\xa8\x99\x80\x0f\xcd\x97\xa9\xc9\x85\x46\x9f\xeb\xf8\xcc\x20\xde\x83\x43\x81\x0c\xf8\x9e\xae\x8d\x2f\xcd\xf9\x85\x39\xeb\xf7\xa6\x9a\x78\xf1\x13\x57\x55\x5a\x02\x0b\xd5\xb8\xf7\xeb\x70\x32\x9c\xce\xcd\xeb\xd9\xf0\xca\x9c\xf6\x26\x43\x05\x4b\xf9\xbc\x67\x1f\xe6\x65\x6f\x36\xfb\xf3\xe2\x6a\xa0\x60\x67\x0e\x25\xab\x3d\x4f\x80\x44\xa7\x80\xa9\xc5\xf8\x57\xd5\xca\x8b\x62\xa9\x2b\x62\x60\x7a\xcf\x5f\x18\x29\x7d\xfc\x16\x04\x7b\x74\x02\x5e\x6c\xfe\x13\x00\x00\xff\xff\xbb\x94\x05\xb5\xf9\x2d\x00\x00")

func idracPyBytes() ([]byte, error) {
	return bindataRead(
		_idracPy,
		"idrac.py",
	)
}

func idracPy() (*asset, error) {
	bytes, err := idracPyBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "idrac.py", size: 11769, mode: os.FileMode(493), modTime: time.Unix(1466599199, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _idracToml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xf2\x49\xcc\x4b\x2f\x4d\x4c\x4f\x55\xb0\x55\x50\x2a\xa8\x2c\xc9\xc8\xcf\x53\xe2\x82\x89\x85\xa5\x16\x15\x67\xe6\xe7\x01\xa5\x8c\xb9\xb8\xa2\xa3\x8b\x52\x0b\x4b\x33\x8b\x52\x73\x53\xf3\x4a\x62\x63\xb9\xf2\x12\x73\xc1\x9a\x40\xa2\xa9\xc5\x25\xc5\x4a\x5c\x65\x70\xe5\x4a\x46\x7a\x86\x06\x7a\x06\x4a\x80\x00\x00\x00\xff\xff\x20\xd2\xea\x2d\x5d\x00\x00\x00")

func idracTomlBytes() ([]byte, error) {
	return bindataRead(
		_idracToml,
		"idrac.toml",
	)
}

func idracToml() (*asset, error) {
	bytes, err := idracTomlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "idrac.toml", size: 93, mode: os.FileMode(420), modTime: time.Unix(1466505779, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _iloPy = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xcc\x59\x7b\x6f\xe3\xb8\x11\xff\xdf\x9f\x82\xe7\x45\x21\x09\xf5\x7a\x1d\x27\xbb\xdd\x4d\x91\x02\x3e\xc7\x69\x8d\xc6\x4e\x10\xe7\xfa\x40\x10\x08\xb2\x44\xc7\xbc\x95\x28\x9d\x28\x25\xf1\x15\xf7\xdd\x3b\x43\x52\x12\x65\x49\x89\xb3\xe8\x61\x2b\xff\x23\x93\x9c\xe1\x3c\x7e\xf3\x20\xf5\xee\x87\x0f\xb9\x48\x3f\xac\x19\xff\x40\xf9\x23\x49\x76\xd9\x36\xe6\xbd\x1e\x8b\x92\x38\xcd\xc8\xcf\x02\xfe\xe8\xf7\x58\x14\x6f\x62\x27\x7a\x9b\x34\x8e\x88\x1f\x27\x3b\xa2\x07\x03\x4a\x13\xfc\x5f\xd2\x6e\x13\x16\xc6\xbd\x5e\x6f\x31\x99\xba\xd7\x37\xb3\x8b\xf9\xbf\xdc\x1f\x2f\x27\xd3\xbf\x5f\xce\x57\xb7\xe4\x8c\xdc\xf5\x08\x3c\xd6\xc7\x11\xfc\x4e\xac\x01\xb1\x8e\x8f\x3f\x8e\x3e\x5d\xe0\xdb\x68\xf4\xe5\xf3\x9f\x3e\xa9\x37\x7c\xca\xb7\x29\xbe\x8d\x47\x27\x47\x1f\x8f\xf1\xed\xe8\xe4\xcb\xd1\x18\x66\x15\xa7\xd1\x78\xa4\x39\x5d\xcc\x2e\xe0\x91\x2b\x26\x17\x5f\xc6\x8a\x7e\x3c\xfa\xac\xde\xce\x67\x93\xf3\xf1\x54\xad\x9b\x9c\x9f\x9c\x03\xfd\x7d\xaf\x77\x3e\xfb\xc7\x7c\x3a\x73\xe7\xcb\x8b\x2b\xf7\x76\xb6\xb8\xbe\x9c\xdc\xce\x40\xca\xff\x48\xde\xfd\x28\x0e\x68\xe8\x72\x2f\xa2\xfd\x53\xd2\xef\xab\x1d\xfb\x49\x1a\xfb\x54\x88\x38\x15\x30\x7a\x77\xaf\x47\x23\xcf\x77\xbd\x20\x48\x61\x86\xd6\x26\x02\x26\xbe\xea\x01\x52\x3e\xef\x48\xce\x73\x41\x03\x62\x4b\x7b\x91\x20\xa6\x82\x5b\x19\x01\xde\x8f\x2c\xa0\x44\xe4\xfe\x96\x30\xbe\x89\x1d\xc5\x44\xd0\x94\x79\x20\x4a\x1e\xad\x69\x6a\x0a\x13\xd1\x28\x4e\x77\x7a\xbf\xdf\x7a\xbd\xeb\x9b\xab\xe9\x6c\xb5\xba\xba\x39\x40\x9d\xb7\x8b\xb1\xf1\x22\x16\xee\x0a\x6a\xa5\x88\xa2\x4c\x69\x96\xa7\x5c\xc0\x62\x80\x00\x4d\xe9\x80\xac\x73\x00\x0f\x0f\x77\x64\x13\xa7\x84\x5d\x5e\x8d\x89\xad\xdf\x8e\x49\xb6\x65\x82\x6c\x18\x0d\x03\x02\x2f\x34\x4a\xb2\x9d\x83\xbc\x78\xfc\x8b\xa7\x76\x0a\xbd\x35\x0d\x4d\x45\x19\x0f\xe8\x33\x0c\x2c\x63\x4e\x07\x35\x1b\xaa\xad\x04\x03\xd1\xbc\x54\x4a\x8b\x4c\xbd\x47\x0f\xfe\xaf\x43\x4a\x3c\x51\xb0\xd3\xb6\x4c\x28\x0d\x0a\x4e\x6a\xc8\x8f\x53\xe9\x34\x35\x04\x66\x5c\xcc\x16\x57\x37\xff\x6e\xb1\x61\x43\x2e\xc1\x7e\xa5\x75\x66\x2d\xfc\xf7\x84\xff\x06\xc9\x41\xa6\x9e\x1f\x7a\x42\x90\x79\x18\xcf\xd2\x34\x4e\xed\xd9\xb3\x4f\x93\x8c\xc5\xdc\x39\x95\xbb\x24\x30\x0b\xab\x02\xba\x01\x3b\xa6\x91\x17\x82\x64\xae\x01\x4b\xdb\x78\xd7\x24\xc6\x08\xa8\x68\xfc\x1b\xe6\x49\x42\x53\xdb\x19\xa6\x34\x09\x3d\x9f\xda\xd6\x7b\x0c\x9d\x53\x4b\x01\x41\x79\xdb\x24\xd0\x1b\x3f\xd0\xcc\x05\x38\xb8\x8c\x8b\xcc\xe3\x40\xb7\x8d\x45\x36\x20\xa0\x6c\x3a\x90\x02\x3e\xc5\x69\xa0\x37\x47\xd8\x9c\x29\xf8\x0c\x41\x29\xb9\x14\xc1\x79\xa6\x68\xc2\xf8\x81\xf1\xb3\x3a\xe5\x59\xc9\xc2\x14\x43\x25\x1c\xdc\xde\xc5\xfd\x41\x2a\x61\xa7\xde\x93\x7c\x19\xe0\xac\xfb\x48\x53\x51\x19\xea\x1d\xb9\xdd\x52\x12\x78\x99\x47\x44\x96\xe6\x3e\x30\xa1\x12\xa5\x90\xb3\x48\x19\xc3\x9a\x3b\x04\x87\x4c\x7a\x0a\xe6\xe0\x9e\x04\xc6\xb3\x1d\xe1\x9e\xc8\x76\x03\xcd\x8f\x82\xcf\x7d\x08\xd0\x0a\xed\xc7\xc4\xe6\x31\xf1\x43\x0a\x8e\x85\x0c\x90\x31\xee\xa3\xab\xc8\x9a\x66\x4f\x94\x72\xc0\xfc\x9a\x06\x01\x30\x5f\xce\xa7\xe0\x72\x1e\x68\x4e\x6c\x35\x5d\xcd\x09\x66\x52\xe1\x0c\x95\x99\x36\xa6\x0a\xe4\xec\x8c\x1c\x2b\x35\xf0\x01\x2b\xa7\x60\xf1\xe0\x19\x2c\x39\x92\xa3\x34\x14\xb4\x7d\xfe\x68\xdf\xe5\x14\x9d\x7e\x77\x2f\x87\x51\xec\x08\x20\x48\x0a\xc3\x55\x3c\x64\x9c\x4a\x7c\x0c\xc1\xbc\xb6\xa5\xfe\x03\x1c\xee\xee\x9d\x6a\x11\xaa\xad\xe8\xf9\x03\xb5\xcb\x7d\xc1\x8f\x94\xdb\x8a\xc4\x19\x90\xb1\x53\xf1\xd5\xca\xd9\xb5\x81\x6a\xc3\x3b\x76\x7f\x67\x21\x1e\xac\x7b\x54\xda\xba\x06\x9b\x58\xa5\xa5\x3a\xd6\x3f\x7a\x61\x8e\x04\x3f\x00\x01\x78\xc1\x42\x93\x42\xfc\xc4\xfc\x41\x90\x2c\x26\xd1\x43\x94\x15\x1e\xae\xb1\xd9\x13\x4b\x1b\x0a\x74\x6e\x8f\xa4\x62\x47\xf2\x47\x72\x54\xed\xea\x34\x78\x80\x7a\x40\x76\x77\xfa\xe9\x1e\x18\x65\x68\x9e\xb6\xb2\xd8\xdc\xbb\xe1\xa8\xa1\x07\xf1\xc8\x03\x0c\xe1\xae\x08\xa4\xc2\x0c\x02\x99\x84\x6c\x51\x22\x5e\x78\x51\x02\x19\x45\x93\x49\x81\x0d\x5c\x43\x3e\x1b\x7f\x3c\x3e\x26\x8b\xbf\xfd\xda\x2f\x40\x27\x10\xec\x28\x35\x26\x2d\x03\x51\x60\x14\x48\xf2\xb6\x18\x8a\x24\x64\x99\xdd\x27\x7d\xe7\x6e\x74\x5f\x93\xaa\x26\x49\x55\x2e\x65\x50\xe2\x5f\x14\x4b\xae\x2f\x52\x86\xcc\xbf\xb6\x6f\x78\xe1\x00\x89\x4f\x48\xbc\x21\x27\x44\xd2\xfe\x99\x7c\x86\xa2\x92\x52\x2f\x10\xfd\x9e\x61\x7d\xbf\x5d\x09\x7c\x7c\xad\x88\x5f\x2a\x02\xec\x0c\x5d\x0c\x7d\x7c\x25\x6d\xa5\x49\x3d\x6e\x92\x22\x6e\xa4\x6e\xd5\x2e\xf8\x17\x56\x16\x2d\x92\xdd\xac\xcd\x4e\x6d\xed\x9d\xa5\xea\x2b\x02\x1e\x33\x93\x9d\xa8\x80\xbb\x50\xa3\x03\x28\x3a\xce\x3e\x85\x2c\x0f\x92\x40\x2f\xbe\x94\x03\x72\xed\xde\x52\x09\x09\xb9\xd4\x40\x88\xa6\x5a\xc9\xb9\x06\x77\x69\x5b\x49\x52\x79\x49\x53\xcc\x9e\xa9\x9f\xcb\x94\x76\x4b\xfd\x2d\x8f\x21\x5d\xef\xf6\x19\x28\x63\x15\xd8\xc5\x91\x1a\x4c\xaa\x25\xb5\xf4\x2d\x3b\x1a\x95\xc0\xe5\xeb\x3e\x58\xb0\xe2\x96\xc8\x3e\x14\x2b\xa3\x2f\x9f\xc8\xe2\xc7\x1a\x36\x3a\x00\xfe\x2a\xc8\xeb\x40\x97\xa1\x2a\xe5\xec\x4a\xa6\x72\xb2\xe2\x0f\xff\x4d\x50\xec\x75\x1a\x8e\xb9\xce\x74\x6f\xd4\xed\x5e\xb9\x12\xad\x52\x7a\x4a\x9a\x48\x53\xac\x70\xc2\xd9\x5f\xdd\x86\x85\xa8\x03\x0b\x4a\x83\x32\x05\xd1\xa8\x9e\x82\xe4\x2c\x78\xb0\xb5\xac\xd6\x8b\xa8\x3a\x67\xbc\x2f\x6b\xe9\x2f\x39\xcb\x28\x18\xca\x8f\xf9\x23\xe5\x8c\x42\xf3\x88\xb6\x8b\xf3\x14\xb8\x41\xe9\x7f\xef\x7b\x02\x1a\xa6\x0c\x3b\x4a\x98\xa0\xe4\x89\x12\x0e\xd2\x61\x2e\x07\x2c\x6e\xbd\x04\xa8\x21\xa3\x13\x70\x52\x06\x00\x58\xb3\x6c\xa8\x90\x04\x55\x3a\xf1\x52\xea\x62\x2f\xe1\xa2\x4c\x12\x50\xe5\xbf\xb6\xb6\xa0\x9c\x2c\xbb\x3d\x7c\xfa\x70\xec\x71\xb1\x35\x33\x7a\x7a\x39\xde\x7a\x0c\x90\x33\xb5\x9e\xbc\x1a\xed\x38\x20\xfc\xd6\x51\xe5\xc7\xa7\xb5\xe2\x0a\xda\x64\x05\xa4\x4a\x49\x1b\xc5\x14\x57\x69\x2f\xe6\xeb\x9f\xa9\x9f\x59\x8e\x2c\x9e\xab\x9d\xc8\x00\x77\x73\xd0\x03\xea\x19\x86\xad\xd5\xac\x3a\x25\x5b\xc0\x87\x56\xda\xba\x2f\xa3\x17\x38\x37\x0b\x1c\x38\x0e\x5a\x9b\x9c\x1e\x2a\xc7\x75\x61\xb4\xc3\x45\xa9\xec\xfc\xad\xc2\x34\x3b\x0c\x7c\xba\x44\x5c\xa8\x60\x3e\xa7\x8f\xcc\xa7\xed\x0d\xc7\x1e\xb9\x0c\x30\xd9\x72\xa8\x22\x0f\xcd\x4f\x18\x42\x08\xbd\xd6\x61\x18\x4a\x2a\xc8\xfc\x2f\xad\x0d\xd1\x85\x79\x0d\xe5\x37\xe6\x75\xef\xe6\x74\xa7\x3e\x7c\xca\x8e\xaf\x49\xd8\x5c\x0b\xbe\xd4\x27\x39\xae\x09\xdb\xfb\x99\x4e\x47\x94\x0b\x44\x79\x6a\x90\x7c\x06\xd0\x38\xfb\x99\xd3\xe9\x81\x9a\xb4\x55\xeb\x87\x3e\x9c\x19\xdd\xb5\xec\xea\x27\x42\xb0\x07\x1e\x41\x82\xb1\x5a\x79\xb5\x78\xa7\x78\x4c\x2f\x99\x21\xfc\x9a\xb3\x8a\x67\x0d\x5d\xc9\xd7\x97\x5d\x49\xc3\x17\x9b\xfc\x6f\x09\x7f\x08\xb4\x00\x52\x30\x59\x62\x03\xfd\x8a\xbb\x7f\x97\xb8\x6f\xef\x0f\x0e\x16\xe4\xf7\x88\xfa\x4a\x38\x55\x45\x6b\xd2\xb4\xc2\xac\x11\xe4\x6f\x20\x50\x55\xd4\xa4\xf8\x0e\xf9\xe0\xc0\x78\x37\x63\xb8\x49\xdc\x1e\x1b\xaa\xdc\x83\x08\x01\x93\x6e\xc6\x1a\x5d\x6c\x03\x38\xde\x41\x81\x0e\x19\x5e\x63\xa8\x5b\x8e\x6c\xeb\x65\x96\xc0\x52\xde\xc1\x4d\xde\x15\x3d\x79\x3b\xac\xed\xea\xc4\xfc\x90\x33\xb1\x2d\x4f\xcc\x85\x38\x50\xeb\x45\x26\x55\xf7\x18\x87\x45\x1d\xec\xea\x07\xec\x47\x31\x54\x47\x6b\xa1\xce\xd6\xc3\xef\x92\xa2\x0e\x3a\xcd\xb6\xd0\x99\xa9\x0d\x4f\xb5\xff\x97\x29\xcc\xbc\x76\x48\x3d\x26\x68\x75\x53\xd5\xff\x89\x7f\xe5\xf1\x13\x27\x45\x7e\x83\x93\x16\xe8\x71\x4a\xfe\x10\xf4\x87\xaa\x0d\xb0\xcd\x7e\xcc\x29\x5a\x22\xbc\x3b\x68\xcd\x4e\x0e\xf9\x0b\x39\xea\xdc\xaf\x26\x6a\xff\x56\x22\x53\x6c\xe3\x1c\x00\xbe\xa6\x0a\x68\x18\xc0\x6d\x5d\x91\x74\x28\xe9\xd7\x39\x40\x54\x64\x45\x6b\x5b\x36\xb4\xeb\x5d\xad\x9d\x1d\x56\x34\xb5\xf6\xb8\x14\x7f\xef\x86\x4c\xeb\x2a\xf5\x8e\x3c\xee\x3d\xd0\x54\x7b\x70\xf3\x54\xd5\x01\x62\x4c\x63\x50\xba\xd5\xa4\x5d\x5a\xa9\x1a\x53\x71\xab\xd6\x63\xb1\xab\x0e\xdf\xaa\xb9\xe9\xe3\xcd\x54\xbf\xb2\x5b\xad\xe6\x90\xe3\xaa\x18\xbd\x8d\xe5\xb8\x93\xe5\xb8\x05\x1c\xf5\x15\x65\x62\xac\x2e\xf3\x8a\x59\x6d\x31\x1c\x09\x64\x37\x26\x9d\x6f\x5a\xac\xad\x8d\xaf\x55\xc8\x16\x0b\x56\x27\x02\xa7\xd1\xf6\xbf\xed\xd8\xa0\x4f\xa5\xa5\x64\xc0\xa0\xed\x7b\xc2\xfe\xb2\x7a\x71\x2b\x8e\x60\xc6\x35\x49\x47\x19\x6c\xec\xd7\x88\x66\xe0\x55\x61\xbf\xba\x0b\xed\x0e\xff\x16\x7d\x5a\x76\xa9\x7d\x7c\xd8\xdb\xa5\x35\x3a\xe1\xa8\xac\x8b\xa0\xa4\x24\x4b\x45\x29\x4f\xad\x43\x38\x1b\xb2\xc4\xee\xdc\xac\xfa\x48\xf1\xa6\x9d\x6a\xfd\x4e\x75\x3c\x6e\xdb\x40\x97\xd6\xc2\xf0\xfa\xbe\xa1\xa5\xf8\xd6\xe2\xd8\xe0\xa1\x71\x29\x7c\x8f\xbf\x78\xbf\xbd\x91\x32\xcb\x20\xe9\x77\xe7\xc6\x65\x4c\xe6\xd7\xc5\xa5\x24\x16\x3f\xe4\x4b\xb6\x9e\x80\x5c\x45\x79\xf1\x29\x26\x18\xf6\xcb\x78\xc7\xdd\x0e\xe0\x5a\xc5\xac\xa4\x40\x93\xbe\xc8\xf6\x40\x61\x0d\xb6\x85\xc2\x2f\xb0\xad\x82\x4f\x5f\x4d\xbc\xfe\x75\xa0\xd7\x4c\x12\x2f\x25\xcd\x96\x28\x3c\x38\x63\xe8\x6b\x3d\xbc\xea\xc1\x6f\x9e\xc3\x20\x8f\x12\x61\x1b\xa4\x50\x8a\x7a\x3d\xb0\x8d\x2b\x11\xe9\xba\xb2\x04\xbb\xc0\x8c\x71\xd7\xb5\xaa\x7b\x03\xd8\x34\x16\x43\xca\x1f\x59\x5a\xa4\xcc\xf9\xb5\x7b\x7b\xe5\xae\xa6\x93\xa5\x81\x47\xe5\xba\xc6\xda\xc5\x64\x39\xf9\xeb\x6c\x31\x5b\xde\xba\x3f\xad\x66\x37\xee\x72\xb2\x98\x19\x54\xa5\x9d\x5f\xa7\xbc\x9e\xac\x56\xff\xbc\xba\x39\x37\xa8\x33\xf3\xf6\xa9\x1b\xb4\x2a\x53\xcb\x6f\x4a\xc4\x2e\x3c\x3e\xa8\xbe\xcf\x4c\xe3\x28\xca\x39\xf3\x65\x9d\x94\x93\x0e\x7e\xa0\xa2\xe6\x7d\x27\x1a\x92\x0e\xbd\xf4\x41\xd4\x6e\xcb\x20\x5e\x87\xf4\x99\x65\xf6\x91\xd3\xfb\x6f\x00\x00\x00\xff\xff\xb1\x7a\xf4\x6c\x7a\x1e\x00\x00")

func iloPyBytes() ([]byte, error) {
	return bindataRead(
		_iloPy,
		"ilo.py",
	)
}

func iloPy() (*asset, error) {
	bytes, err := iloPyBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "ilo.py", size: 7802, mode: os.FileMode(493), modTime: time.Unix(1466158983, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _iloToml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xf2\x49\xcc\x4b\x2f\x4d\x4c\x4f\x55\xb0\x55\x50\x2a\xa8\x2c\xc9\xc8\xcf\x53\xe2\x82\x89\x85\xa5\x16\x15\x67\xe6\xe7\x01\xa5\x8c\xb9\xb8\xa2\xa3\x8b\x52\x0b\x4b\x33\x8b\x52\x73\x53\xf3\x4a\x62\x63\xb9\xf2\x12\x73\x91\x34\xe9\x66\x14\x64\xe6\xe4\x2b\x71\x95\xc1\xb5\x28\x19\xeb\x59\x28\x01\x02\x00\x00\xff\xff\x5e\xdd\x05\xf1\x5e\x00\x00\x00")

func iloTomlBytes() ([]byte, error) {
	return bindataRead(
		_iloToml,
		"ilo.toml",
	)
}

func iloToml() (*asset, error) {
	bytes, err := iloTomlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "ilo.toml", size: 94, mode: os.FileMode(420), modTime: time.Unix(1466505779, 0)}
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
	"idrac.py": idracPy,
	"idrac.toml": idracToml,
	"ilo.py": iloPy,
	"ilo.toml": iloToml,
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
	"idrac.py": &bintree{idracPy, map[string]*bintree{}},
	"idrac.toml": &bintree{idracToml, map[string]*bintree{}},
	"ilo.py": &bintree{iloPy, map[string]*bintree{}},
	"ilo.toml": &bintree{iloToml, map[string]*bintree{}},
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
