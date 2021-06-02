package classpath

import (
	"os"
	"path/filepath"
)

// Composite pattern

type Classpath struct {
	bootClasspath Entry
	extClasspath  Entry
	userClasspath Entry
}

func Parse(jreOption, cpOption string) *Classpath {
	cp := &Classpath{}
	cp.parseBootAndExtClasspath(jreOption)
	cp.parseUserClasspath(cpOption)

	return cp
}

func (classpath *Classpath) ReadClass(className string) ([]byte, Entry, error) {
	className = className + ".class"
	if data, entry, err := classpath.bootClasspath.readClass(className); err == nil {
		return data, entry, nil
	}

	if data, entry, err := classpath.extClasspath.readClass(className); err == nil {
		return data, entry, nil
	}

	return classpath.userClasspath.readClass(className)
}

func (classpath *Classpath) String() string {
	return classpath.userClasspath.String()
}

func (classpath *Classpath) parseBootAndExtClasspath(jreOption string) {
	jreDir := getJreDir(jreOption)

	// jre/lib/*
	jreLibPath := filepath.Join(jreDir, "lib", "*")
	classpath.bootClasspath = newWildcardEntry(jreLibPath)

	// jre/lib/ext/*
	jreExtPath := filepath.Join(jreDir, "lib", "ext", "*")
	classpath.extClasspath = newWildcardEntry(jreExtPath)
}

func (classpath *Classpath) parseUserClasspath(cpOption string) {
	if cpOption == "" {
		cpOption = "."
	}
	classpath.userClasspath = newEntry(cpOption)
}

func getJreDir(jreOption string) string {
	if jreOption != "" && exists(jreOption) {
		return jreOption
	}

	if exists("./jre") {
		return "./jre"
	}

	if jh := os.Getenv("JAVA_HOME"); jh != "" {
		return filepath.Join(jh, "jre")
	}

	panic("Can not find jre directory!")
}

func exists(path string) bool {
	if _, err := os.Stat(path); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}
