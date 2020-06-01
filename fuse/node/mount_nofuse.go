// +build !windows,nofuse

package node

import (
	"errors"

	core "github.com/IPFS-eX/go-ipfs-ex/core"
)

func Mount(node *core.IpfsNode, fsdir, nsdir string) error {
	return errors.New("not compiled in")
}
