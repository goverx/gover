package service

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
)

func DownloadAndExtract(url, target string) error {
	if _, err := os.Stat(target); err == nil {
		fmt.Printf("Already installed: %s\n", target)
		return nil
	}

	if err := os.MkdirAll(target, 0755); err != nil {
		return err
	}

	fmt.Printf("Downloading %s...\n", url)
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return fmt.Errorf("failed to download: %s", resp.Status)
	}

	return Untar(resp.Body, filepath.Join(target, "go"))
}
