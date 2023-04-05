package fileutil

import (
	"os"
	"path/filepath"
)

// WriteFile write bytes to file. will create all path dir if not exists
func WriteFile(filename string, data []byte) (err error) {
	EnsureDirExists(filename)
	return os.WriteFile(filename, data, 0755) // rwxr-xr-x
}

// OpenWrite open file for write. if the file doesn't exist, create it
func OpenWrite(filename string) (*os.File, error) {
	EnsureDirExists(filename)
	return os.OpenFile(filename, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0755) // rwxr-xr-x
}

// IsExist return true if file exists
func IsExist(filename string) bool {
	_, err := os.Stat(filename)
	return !os.IsNotExist(err)
}

// Delete removes the named file or directory. not mind for error
func Delete(name string) {
	_ = os.Remove(name)
}

// DeleteAll removes path and any children it contains. not mind for error
func DeleteAll(path string) {
	_ = os.RemoveAll(path)
}

// EnsureDirExists create all directory if not exists
// Example:
//
//	EnsureDirExists("/foo/bar/baz.js") // the following directory will be created: /foo/bar/
//	EnsureDirExists("/foo/bar/baz/") // the following directory will be created: /foo/bar/baz/
func EnsureDirExists(path string) {
	var dir string
	if filepath.Ext(path) == "" && filepath.Base(path) != "." { // is dir
		dir = path
	} else { // is file
		dir = filepath.Dir(path)
	}
	_, err := os.Stat(dir)
	if err == nil {
		return
	}
	if os.IsNotExist(err) {
		_ = os.MkdirAll(dir, 0755) // rwxr-xr-x
	}
}
