package efs

import (
	"os"
	"os/exec"
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
	MountType string
	Device    string
	Target    string
	Options   string
	NetNSPid  int
}

func (m *EFSMounter) Mount() error {
	optstring := "hard"
	args := make([]string, 0)
	if m.MountType != "" {
		args = append(args, "-t", m.MountType)
	}
	if m.Options != "" {
		args = append(args, "-o", m.Options)
	}
	args = append(args, m.Device, m.Target)

	mountcmd := exec.Command("mount")
	mountcmd.Args = args
	mountcmd.Stderr = os.Stderr

	if m.NetNSPid != 0 {
		return WithNetNS(m.NetNSPid, mountcmd.Run)
	}

	return mountcmd.Run()
}

func (m *EFSMounter) Unmount() error {
	path, err := exec.LookPath("umount")
	if err != nil {
		return err
	}

	umountCmd := exec.Command(path, m.Target)
	return umountCmd.Run()
}
