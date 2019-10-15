package efs

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
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
	args := []string{}
	if m.MountType != "" {
		args = append(args, "-t", m.MountType)
	}
	if m.Options != "" {
		args = append(args, "-o", m.Options)
	}
	args = append(args, m.Device, m.Target)
	mountcmd := exec.Command("mount", args...)
	mountcmd.Stderr = os.Stderr

	if err := m.Validate(); err != nil {
		return err
	}

	if m.NetNSPid != 0 {
		return WithNetNS(m.NetNSPid, mountcmd.Run)
	}

	return mountcmd.Run()
}

func (m *EFSMounter) Validate() error {
	requiredFields := []string{}
	if m.MountType == "" {
		requiredFields = append(requiredFields, "mountType")
	}
	if m.Device == "" {
		requiredFields = append(requiredFields, "device")
	}
	if m.Target == "" {
		requiredFields = append(requiredFields, "target")
	}
	if len(requiredFields) > 0 {
		return fmt.Errorf("missing required fields: [%s]", strings.Join(requiredFields, ","))
	}
	return nil
}

func (m *EFSMounter) Unmount() error {
	path, err := exec.LookPath("umount")
	if err != nil {
		return err
	}

	umountCmd := exec.Command(path, m.Target)
	return umountCmd.Run()
}
