// +build linux darwin freebsd
// +build !nofuse

package readonly

import (
	core "github.com/IPFS-eX/go-ipfs-ex/core"
	mount "github.com/IPFS-eX/go-ipfs-ex/fuse/mount"
)

// Mount mounts IPFS at a given location, and returns a mount.Mount instance.
func Mount(ipfs *core.IpfsNode, mountpoint string) (mount.Mount, error) {
	cfg, err := ipfs.Repo.Config()
	if err != nil {
		return nil, err
	}
	allow_other := cfg.Mounts.FuseAllowOther
	fsys := NewFileSystem(ipfs)
	return mount.NewMount(ipfs.Process, fsys, mountpoint, allow_other)
}
