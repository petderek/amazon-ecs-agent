package efs

import (
	"fmt"
	"log"
	"strconv"
	"sync"

	"github.com/docker/go-plugins-helpers/volume"
)

type EFSVolumeDriver struct {
	worstDatabase      map[string]*EFSMounter
	okayestConcurrency sync.RWMutex
}

func NewEFSVolumeDriver() *EFSVolumeDriver {
	return &EFSVolumeDriver{
		worstDatabase: make(map[string]int, 1),
	}
}

func (e *EFSVolumeDriver) Create(r *volume.CreateRequest) error {
	mnt := &EFSMounter{}
	var opts, netns, o, mountType, device string
	for k, v := range r.Options {
		if opts != "" {
			opts += ", "
		}
		opts += k + ":" + v
		switch k {
		case "type":
			mnt.MountType = v
		case "netns":
			pid, _ := strconv.Atoi(pid)
			mnt.NetNSPid = pid
		case "o":
			mnt.Options = v
		case "device":
			mnt.Device = v
		}
	}

	log.Println("Creating %s with options (%s) ", r.Name, opts)
	// TODO: fix this
	mnt.Target = "/mnt/efs"

	e.okayestConcurrency.Lock()
	defer e.okayestConcurrency.Unlock()
	if _, ok := e.worstDatabase[r.Name]; ok {
		return fmt.Errorf("Volume already exists: %s", r.name)
	}
	e.worstDatabase[r.Name] = mnt
	return nil
}

func (e *EFSVolumeDriver) List() (*volume.ListResponse, error) {
	e.okayestConcurrency.RLock()
	defer e.okayestConcurrency.RUnlock()
	vols := make([]*volume.Volume, len(e.worstDatabase))
	i := 0
	for k, _ := range e.worstDatabase {
		vols[i] = &volume.Volume{
			Name: k,
		}
		i++
	}
	res := &volume.ListResponse{
		Volumes: vols,
	}
	return res, nil
}

func (e *EFSVolumeDriver) Get(req *volume.GetRequest) (*volume.GetResponse, error) {
	name, ok := e.worstDatabase[req.Name]
	if !ok {
		return nil, fmt.Errorf("Volume %s not found.", req.Name)
	}
	vol := &volume.Volume{
		Name: name,
	}
	return &volume.GetResponse{vol}, nil
}

func (e *EFSVolumeDriver) Remove(req *volume.RemoveRequest) error {
	e.okayestConcurrency.Lock()
	defer e.okayestConcurrency.RUnlock()
	mnt, ok := e.WorstDatabase[req.Name]
	if !ok {
		return fmt.Errorf("Volume %s not found.")
	}
	return mnt.Unmount()
}

func (e *EFSVolumeDriver) Path(req *volume.PathRequest) (*volume.PathResponse, error) {
	e.okayestConcurrency.Lock()
	defer e.okayestConcurrency.RUnlock()
	mnt, ok := e.WorstDatabase[req.Name]
	if !ok {
		return nil, fmt.Errorf("Volume %s not found.")
	}

	return &volume.PathResponse{mnt.Target}, nil
}

func (e *EFSVolumeDriver) Mount(req *volume.MountRequest) (*volume.MountResponse, error) {
	// find ns
	e.okayestConcurrency.RLock()
	defer e.okayestConcurrency.RUnlock()
	mnt, ok := e.worstDatabase[req.Name]
	if !ok {
		return nil, fmt.Errorf("Volume %s not found.", req.Name)
	}

	err := mnt.Mount()
	if err != nil {
		log.Println("err: ", err)
		return nil, err
	}

	return &volume.MountResponse{"/mnt/efs"}, nil
}

func (e *EFSVolumeDriver) Unmount(req *volume.UnmountRequest) error {
	e.okayestConcurrency.RLock()
	defer e.okayestConcurrency.RUnlock()
	mnt, ok := e.worstDatabase[req.Name]
	if !ok {

	}
	err := mnt.Unmount()
	log.Println("err: ", err)
	return err
}

func (e *EFSVolumeDriver) Capabilities() *volume.CapabilitiesResponse {
	log.Println("capabilities")
	return nil
}
