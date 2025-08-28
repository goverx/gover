package core

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"

	"gover/internal/service"
)

type goManager struct {
	baseDir     string
	versionsDir string
	binDir      string
}

var Manager = goManager{
	baseDir:     filepath.Join(os.Getenv("HOME"), ".gover"),
	versionsDir: filepath.Join(os.Getenv("HOME"), ".gover", "versions"),
	binDir:      filepath.Join(os.Getenv("HOME"), ".gover", "bin"),
}

func (m goManager) Install(version string) error {
	operationSystem := runtime.GOOS
	arch := runtime.GOARCH

	if operationSystem == "darwin" && arch == "arm64" {
		arch = "arm64"
	} else if operationSystem == "darwin" && arch == "amd64" {
		arch = "amd64"
	}

	url := fmt.Sprintf("https://go.dev/dl/go%s.%s-%s.tar.gz", version, operationSystem, arch)
	target := filepath.Join(m.versionsDir, version)

	return service.DownloadAndExtract(url, target)
}

func (m goManager) Use(version string) error {
	return service.LinkBinaries(filepath.Join(m.versionsDir, version, "go", "bin"), m.binDir)
}

func (m goManager) List() ([]string, error) {
	return service.ListDirs(m.versionsDir)
}

func (m goManager) Current() (string, error) {
	goBin := filepath.Join(m.binDir, "go")
	out, err := exec.Command(goBin, "version").CombinedOutput()
	return string(out), err
}
