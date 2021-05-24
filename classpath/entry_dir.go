package classpath

import (
	"io/ioutil"
	"path/filepath"
)

// Directory classpath

type DirEntry struct {
	absDir string
}

func newDirEntry(path string) *DirEntry {
	absDir, err := filepath.Abs(path)
	if err != nil {
		panic(err)
	}
	return &DirEntry{absDir}
}

func (dirEntry *DirEntry) readClass(className string) ([]byte, Entry, error) {
	fileName := filepath.Join(dirEntry.absDir, className)
	data, err := ioutil.ReadFile(fileName)
	return data, dirEntry, err
}

func (dirEntry *DirEntry) String() string {
	return dirEntry.absDir
}
