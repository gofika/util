package fileutil

import (
	"os"
	"path/filepath"
)

// CurrentDir return the directory where the program is located.
func CurrentDir() (currentDir string) {
	currentDir, _ = filepath.Abs(filepath.Dir(os.Args[0]))
	return
}
