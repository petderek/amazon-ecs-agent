package efs

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

const (
	efsBin = "mount.efs"
)

func HasEFSSupport() bool {
	return findEFSBin() != ""
}

func findEFSBin() string {
	path, _ := exec.LookPath(efsBin)
	return path
}

type EFSMounter struct {
	Filesystem        string
	TransitEncryption bool
	ReadOnly          bool
	RootDirectory     string
	LocalTarget       string
}

func (m *EFSMounter) Mount() error {
	optstring := "hard"
	if m.TransitEncryption {
		optstring += ",tls"
	}
	if m.ReadOnly {
		optstring += ",ro"
	}

	remoteTarget := fmt.Sprintf("%s:%s", m.Filesystem, filepath.Join("/", m.RootDirectory))
	mountcmd := exec.Command("mount.efs", remoteTarget, m.LocalTarget, "-o", optstring)
	mountcmd.Stderr = os.Stderr
	return mountcmd.Run()
}

func (m *EFSMounter) Unmount() error {
	path, err := exec.LookPath("umount")
	if err != nil {
		return err
	}

	umountCmd := exec.Command(path, m.LocalTarget)
	return umountCmd.Run()
}
