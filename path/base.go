package path

import (
	"os"
	"path/filepath"
)

func Join(p ...string) string {
	execpath, _ := os.Executable()
	baseDir := filepath.Dir(execpath)
	rootDir := filepath.Dir(baseDir)

	paths := []string{
		rootDir,
	}
	paths = append(paths, p...)

	path_ := filepath.Join(paths...)

	return path_
}
