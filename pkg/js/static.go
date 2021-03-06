// Code generated by "esc"; DO NOT EDIT.

package js

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path"
	"sync"
	"time"
)

type _escLocalFS struct{}

var _escLocal _escLocalFS

type _escStaticFS struct{}

var _escStatic _escStaticFS

type _escDirectory struct {
	fs   http.FileSystem
	name string
}

type _escFile struct {
	compressed string
	size       int64
	modtime    int64
	local      string
	isDir      bool

	once sync.Once
	data []byte
	name string
}

func (_escLocalFS) Open(name string) (http.File, error) {
	f, present := _escData[path.Clean(name)]
	if !present {
		return nil, os.ErrNotExist
	}
	return os.Open(f.local)
}

func (_escStaticFS) prepare(name string) (*_escFile, error) {
	f, present := _escData[path.Clean(name)]
	if !present {
		return nil, os.ErrNotExist
	}
	var err error
	f.once.Do(func() {
		f.name = path.Base(name)
		if f.size == 0 {
			return
		}
		var gr *gzip.Reader
		b64 := base64.NewDecoder(base64.StdEncoding, bytes.NewBufferString(f.compressed))
		gr, err = gzip.NewReader(b64)
		if err != nil {
			return
		}
		f.data, err = ioutil.ReadAll(gr)
	})
	if err != nil {
		return nil, err
	}
	return f, nil
}

func (fs _escStaticFS) Open(name string) (http.File, error) {
	f, err := fs.prepare(name)
	if err != nil {
		return nil, err
	}
	return f.File()
}

func (dir _escDirectory) Open(name string) (http.File, error) {
	return dir.fs.Open(dir.name + name)
}

func (f *_escFile) File() (http.File, error) {
	type httpFile struct {
		*bytes.Reader
		*_escFile
	}
	return &httpFile{
		Reader:   bytes.NewReader(f.data),
		_escFile: f,
	}, nil
}

func (f *_escFile) Close() error {
	return nil
}

func (f *_escFile) Readdir(count int) ([]os.FileInfo, error) {
	if !f.isDir {
		return nil, fmt.Errorf(" escFile.Readdir: '%s' is not directory", f.name)
	}

	fis, ok := _escDirs[f.local]
	if !ok {
		return nil, fmt.Errorf(" escFile.Readdir: '%s' is directory, but we have no info about content of this dir, local=%s", f.name, f.local)
	}
	limit := count
	if count <= 0 || limit > len(fis) {
		limit = len(fis)
	}

	if len(fis) == 0 && count > 0 {
		return nil, io.EOF
	}

	return fis[0:limit], nil
}

func (f *_escFile) Stat() (os.FileInfo, error) {
	return f, nil
}

func (f *_escFile) Name() string {
	return f.name
}

func (f *_escFile) Size() int64 {
	return f.size
}

func (f *_escFile) Mode() os.FileMode {
	return 0
}

func (f *_escFile) ModTime() time.Time {
	return time.Unix(f.modtime, 0)
}

func (f *_escFile) IsDir() bool {
	return f.isDir
}

func (f *_escFile) Sys() interface{} {
	return f
}

// _escFS returns a http.Filesystem for the embedded assets. If useLocal is true,
// the filesystem's contents are instead used.
func _escFS(useLocal bool) http.FileSystem {
	if useLocal {
		return _escLocal
	}
	return _escStatic
}

// _escDir returns a http.Filesystem for the embedded assets on a given prefix dir.
// If useLocal is true, the filesystem's contents are instead used.
func _escDir(useLocal bool, name string) http.FileSystem {
	if useLocal {
		return _escDirectory{fs: _escLocal, name: name}
	}
	return _escDirectory{fs: _escStatic, name: name}
}

// _escFSByte returns the named file from the embedded assets. If useLocal is
// true, the filesystem's contents are instead used.
func _escFSByte(useLocal bool, name string) ([]byte, error) {
	if useLocal {
		f, err := _escLocal.Open(name)
		if err != nil {
			return nil, err
		}
		b, err := ioutil.ReadAll(f)
		_ = f.Close()
		return b, err
	}
	f, err := _escStatic.prepare(name)
	if err != nil {
		return nil, err
	}
	return f.data, nil
}

// _escFSMustByte is the same as _escFSByte, but panics if name is not present.
func _escFSMustByte(useLocal bool, name string) []byte {
	b, err := _escFSByte(useLocal, name)
	if err != nil {
		panic(err)
	}
	return b
}

// _escFSString is the string version of _escFSByte.
func _escFSString(useLocal bool, name string) (string, error) {
	b, err := _escFSByte(useLocal, name)
	return string(b), err
}

// _escFSMustString is the string version of _escFSMustByte.
func _escFSMustString(useLocal bool, name string) string {
	return string(_escFSMustByte(useLocal, name))
}

var _escData = map[string]*_escFile{

	"/helpers.js": {
		name:    "helpers.js",
		local:   "pkg/js/helpers.js",
		size:    22817,
		modtime: 0,
		compressed: `
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
Ef83AAD//yJmQYUhWQAA
`,
	},
}

var _escDirs = map[string][]os.FileInfo{}
