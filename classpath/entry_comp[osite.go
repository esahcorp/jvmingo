package classpath

import (
	"errors"
	"strings"
)

// Composite classpath spilt by pathListSeparator

type CompositeEntry []Entry

func newCompositeEntry(pathList string) CompositeEntry {
	var compositeEntry []Entry
	for _, path := range strings.Split(pathList, pathListSeparator) {
		entry := newEntry(path)
		// slice
		compositeEntry = append(compositeEntry, entry)
	}
	return compositeEntry
}

func (compositeEntry CompositeEntry) readClass(className string) ([]byte, Entry, error) {
	for _, entry := range compositeEntry {
		data, from, err := entry.readClass(className)
		if err == nil {
			return data, from, nil
		}
	}
	return nil, nil, errors.New("class not found: " + className)
}

func (compositeEntry CompositeEntry) String() string {
	strs := make([]string, len(compositeEntry))
	for i, entry := range compositeEntry {
		strs[i] = entry.String()
	}
	return strings.Join(strs, pathListSeparator)
}
