package pathutil

import (
	"os"
	"path/filepath"
)

// CurrentDir return current path
func CurrentDir() (currentDir string) {
	currentDir, _ = filepath.Abs(filepath.Dir(os.Args[0]))
	return
}
