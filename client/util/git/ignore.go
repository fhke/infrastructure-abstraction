package git

import (
	"os"
	"path/filepath"
)

// IgnoreDir creates a catch-all .gitignore file in dir given by path
func IgnoreDir(path string) error {
	ignoreFile := filepath.Join(path, ".gitignore")
	return os.WriteFile(ignoreFile, []byte("*"), 0640)
}
