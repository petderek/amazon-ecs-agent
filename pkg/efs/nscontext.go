// +build go1.10
package efs

import (
	"fmt"
	"runtime"

	"github.com/vishvananda/netns"
)

// Inspired by weave
// https://github.com/weaveworks/weave/blob/master/net/netns.go
func WithNetNS(pid int, work func() error) error {
	runtime.LockOSThread()
	defer runtime.UnlockOSThread()

	path := fmt.Sprintf("/proc/%d/ns/net", pid)
	ns, err := netns.GetFromPath(path)
	if err != nil {
		return err
	}

	oldNs, err := netns.Get()
	if err == nil {
		defer oldNs.Close()

		err = netns.Set(ns)
		if err == nil {
			defer netns.Set(oldNs)

			err = work()
		}
	}

	return err
}
