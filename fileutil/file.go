package fileutil

import (
	"io/ioutil"
	"os"
	"path/filepath"
)

// WriteFile write bytes to file. will create all path dir if not exists
func WriteFile(filename string, data []byte) (err error) {
	EnsureDirExists(filename)
	return ioutil.WriteFile(filename, data, 0755)
}

// OpenWrite open file for write. if the file doesn't exist, create it
func OpenWrite(filename string) (*os.File, error) {
	EnsureDirExists(filename)
	return os.OpenFile(filename, os.O_WRONLY|os.O_CREATE, 0755)
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

// EnsureDirExists create all parent paths if not exists
// Example:
//     EnsureDirExists("/foo/bar/baz.js") // will create path /foo/bar
//     EnsureDirExists("/foo/bar/baz/") // will create path /foo/bar
func EnsureDirExists(path string) {
	dir := filepath.Dir(path)
	_, err := os.Stat(dir)
	if err == nil {
		return
	}
	if os.IsNotExist(err) {
		_ = os.MkdirAll(dir, os.ModePerm)
	}
}
