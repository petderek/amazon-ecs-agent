package efs

import (
	"log"

	"github.com/docker/go-plugins-helpers/volume"
)

type EFSVolumeDriver struct {
}

func (e *EFSVolumeDriver) Create(r *volume.CreateRequest) error {
	log.Println("implement me")
	log.Println("name: ", r.Name)
	var opts string
	for k, v := range r.Options {
		if opts != "" {
			opts += ", "
		}
		opts += k + ":" + v
	}
	log.Println("opts: ", opts)
	return nil
}

func (e *EFSVolumeDriver) List() (*volume.ListResponse, error) {
	log.Println("implement me")
	res := &volume.ListResponse{
		Volumes: []*volume.Volume{
			{
				Name: "PotatoSalad",
			},
		},
	}
	return res, nil
}

func (e *EFSVolumeDriver) Get(*volume.GetRequest) (*volume.GetResponse, error) {
	vol := &volume.Volume{
		Name:       "PotatoSalad",
		Mountpoint: "/mnt/efs",
	}
	return &volume.GetResponse{vol}, nil
}

func (e *EFSVolumeDriver) Remove(*volume.RemoveRequest) error {
	log.Println("implement me")
	return nil
}

func (e *EFSVolumeDriver) Path(*volume.PathRequest) (*volume.PathResponse, error) {
	return &volume.PathResponse{"/mnt/efs"}, nil
}

func (e *EFSVolumeDriver) Mount(*volume.MountRequest) (*volume.MountResponse, error) {
	mnt := &EFSMounter{
		Filesystem:  "fs-990ae132",
		LocalTarget: "/mnt/efs",
	}

	err := mnt.Mount()
	if err != nil {
		log.Println("err: ", err)
		return nil, err
	}

	return &volume.MountResponse{"/mnt/efs"}, nil
}

func (e *EFSVolumeDriver) Unmount(*volume.UnmountRequest) error {
	mnt := &EFSMounter{
		LocalTarget: "/mnt/efs",
	}
	err := mnt.Unmount()
	log.Println("err: ", err)
	return err
}

func (e *EFSVolumeDriver) Capabilities() *volume.CapabilitiesResponse {
	log.Println("capabilities")
	return nil
}
