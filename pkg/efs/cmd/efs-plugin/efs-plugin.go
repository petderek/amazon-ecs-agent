package main

import (
	"os/user"
	"strconv"

	"github.com/aws/amazon-ecs-agent/pkg/efs"
	"github.com/docker/go-plugins-helpers/volume"
)

func main() {
	driver := &efs.EFSVolumeDriver{}
	handler := volume.NewHandler(driver)
	rootUser, _ := user.Lookup("root")
	gid, _ := strconv.Atoi(rootUser.Gid)
	handler.ServeUnix("efs-plugin", gid)
}
