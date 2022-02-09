package command

import (
	"github.com/hanwen/go-fuse/v2/fuse"
)

func osSpecificMountOptions() []fuse.MountOption {
	return []fuse.MountOption{}
}

func checkMountPointAvailable(dir string) bool {
	return true
}
