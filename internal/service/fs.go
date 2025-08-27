package service

import (
	"os"
	"path/filepath"
)

func ListDirs(path string) ([]string, error) {
	entries, err := os.ReadDir(path)
	if err != nil {
		return nil, err
	}
	var dirs []string
	for _, e := range entries {
		if e.IsDir() {
			dirs = append(dirs, e.Name())
		}
	}
	return dirs, nil
}

func LinkBinaries(src, dst string) error {
	if err := os.MkdirAll(dst, 0755); err != nil {
		return err
	}

	links := []string{"go", "gofmt"}
	for _, bin := range links {
		linkPath := filepath.Join(dst, bin)
		_ = os.Remove(linkPath)
		if err := os.Symlink(filepath.Join(src, bin), linkPath); err != nil {
			return err
		}
	}
	return nil
}
