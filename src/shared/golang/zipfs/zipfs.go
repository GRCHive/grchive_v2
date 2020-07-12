package zipfs

import (
	"archive/zip"
	"golang.org/x/tools/godoc/vfs"
	"golang.org/x/tools/godoc/vfs/zipfs"
	"net/http"
	"os"
)

type ZipFile struct {
	rsc  vfs.ReadSeekCloser
	name string
	zip  vfs.FileSystem
}

func (f ZipFile) Read(p []byte) (n int, err error) {
	return f.rsc.Read(p)
}

func (f ZipFile) Close() error {
	return f.rsc.Close()
}

func (f ZipFile) Seek(offset int64, whence int) (int64, error) {
	return f.rsc.Seek(offset, whence)
}

func (f ZipFile) Readdir(count int) ([]os.FileInfo, error) {
	files, err := f.zip.ReadDir(f.name)
	if err != nil {
		return nil, err
	}
	return files[:count], nil
}

func (f ZipFile) Stat() (os.FileInfo, error) {
	return f.zip.Stat(f.name)
}

type ZipFS struct {
	Prefix string
	Zip    vfs.FileSystem

	z *zip.ReadCloser
}

func (z ZipFS) Open(name string) (http.File, error) {
	fullName := z.Prefix + name
	rsc, err := z.Zip.Open(fullName)
	if err != nil {
		return nil, err
	}

	return ZipFile{
		rsc:  rsc,
		name: fullName,
		zip:  z.Zip,
	}, nil
}

func (z *ZipFS) Close() {
	z.z.Close()
}

func CreateZipFSFromZipFile(prefix string, name string, fname string) (*ZipFS, error) {
	z, err := zip.OpenReader(fname)
	if err != nil {
		return nil, err
	}

	return &ZipFS{
		Prefix: prefix,
		Zip:    zipfs.New(z, name),
		z:      z,
	}, nil
}
