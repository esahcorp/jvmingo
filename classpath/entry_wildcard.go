package classpath

import (
	"io/fs"
	"path/filepath"
	"strings"
)

// Classpath with wildcard '*'

func newWildcardEntry(path string) CompositeEntry {
	// Remove '*'
	baseDir := path[:len(path)-1]
	var compositeEntry []Entry

	walkFn := func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() && path != baseDir {
			return filepath.SkipDir
		}

		if strings.HasSuffix(path, ".jar") || strings.HasSuffix(path, ".JAR") {
			jarEntry := newZipEntry(path)
			compositeEntry = append(compositeEntry, jarEntry)
		}

		return nil
	}
	err := filepath.Walk(baseDir, walkFn)
	if err != nil {
		return nil
	}

	return compositeEntry
}
