package core

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

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
	url := fmt.Sprintf("https://go.dev/dl/go%s.linux-amd64.tar.gz", version)
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
