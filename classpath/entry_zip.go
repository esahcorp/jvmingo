package classpath

import (
	"archive/zip"
	"errors"
	"io/ioutil"
	"path/filepath"
)

// Zip file classpath

type ZipEntry struct {
	absPath string
}

func newZipEntry(path string) *ZipEntry {
	absPath, err := filepath.Abs(path)
	if err != nil {
		panic(err)
	}
	return &ZipEntry{absPath}
}

func (zipEntry *ZipEntry) readClass(className string) ([]byte, Entry, error) {
	r, err := zip.OpenReader(zipEntry.absPath)
	if err != nil {
		return nil, nil, err
	}

	defer r.Close()
	for _, f := range r.File {
		if f.Name == className {
			rc, err := f.Open()
			if err != nil {
				return nil, nil, err
			}

			data, err := ioutil.ReadAll(rc)
			if err != nil {
				return nil, nil, err
			}

			e := rc.Close()
			if e != nil {
				return nil, nil, e
			}

			return data, zipEntry, nil
		}
	}

	return nil, nil, errors.New("class not found: " + className)
}

func (zipEntry *ZipEntry) String() string {
	return zipEntry.absPath
}
