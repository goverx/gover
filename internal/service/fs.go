package service

import (
	"os"
	"fmt"
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

func LinkBinaries(srcDir, dstDir string) error {
	if _, err := os.Stat(srcDir); os.IsNotExist(err) {
		return fmt.Errorf("Go not found at %s", srcDir)
	}

	if err := os.MkdirAll(dstDir, 0o755); err != nil {
		return fmt.Errorf("failed to create binDir: %w", err)
	}

	entries, err := os.ReadDir(srcDir)
	if err != nil {
		return fmt.Errorf("failed to read %s: %w", srcDir, err)
	}

	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}

		src := filepath.Join(srcDir, entry.Name())
		dst := filepath.Join(dstDir, entry.Name())

		_ = os.Remove(dst)

		if err := os.Symlink(src, dst); err != nil {
			return fmt.Errorf("failed to symlink %s -> %s: %w", src, dst, err)
		}
	}

	return nil
}
