/**
 * Description:
 * Author: Yihen.Liu
 * Create: 2020-05-09
 */

package commands

import (
"context"

iface "github.com/ipfs/interface-go-ipfs-core"

"github.com/ETHFSx/go-ipfs/core/commands/cmdenv"

cmds "github.com/ipfs/go-ipfs-cmds"
"github.com/ipfs/interface-go-ipfs-core/path"
)

const (
	countOptionName         = "count"
	pushRecursiveOptionName = "recursive"
	defaultCopyNum          = 3
	maxCopyNum              = 256
)

var PushCmd = &cmds.Command{
	Helptext: cmds.HelpText{
		Tagline: "Push a file or directory to ipfs.",

		ShortDescription: `
Push contents of <path> to ipfs. Use -r to add directories (recursively).`,

		LongDescription: `
Pushs contents of <path> to ipfs. Use -r to add directories.
Note that directories are added recursively, to form the ipfs
MerkleDAG.`,
	},

	Arguments: []cmds.Argument{
		cmds.StringArg("ipfs-path", true, true, "The cid of a file to be distributed to ipfs.").EnableStdin(),
	},
	Options: []cmds.Option{
		cmds.IntOption(countOptionName, "n", "How many copy number will be distributed in IPFS system.").WithDefault(defaultCopyNum),
		cmds.BoolOption(pushRecursiveOptionName, "r", "Recursively push block context.").WithDefault(false),
	},
	PreRun: func(req *cmds.Request, env cmds.Environment) error {
		cnt, _ := req.Options[countOptionName].(int)
		switch {
		case cnt < defaultCopyNum:
			req.Options[countOptionName] = defaultCopyNum
		case cnt > maxCopyNum:
			req.Options[countOptionName] = maxCopyNum
		}
		return nil
	},

	Run: func(req *cmds.Request, res cmds.ResponseEmitter, env cmds.Environment) error {
		api, err := cmdenv.GetApi(env, req)
		if err != nil {
			return err
		}

		root := req.Arguments[0]
		ctx := req.Context
		peers, err := api.Swarm().Peers(ctx)
		if err != nil {
			return err
		}

		node, err := api.ResolveNode(ctx, path.New(root))
		if err != nil {
			return err
		}

		copyNum, _ := req.Options[countOptionName].(uint32)
		for _, peer := range peers {
			log.Info("distribute to peer: ", peer.ID().String())
			if err := api.Block().Push(ctx, copyNum, peer.ID(), node.Cid()); err != nil {
				log.Errorf("push block err, push to node:%s, cid:%s", peer.ID().String(), node.Cid().String())
			}

		}

		if rec, _ := req.Options[recursiveOptionName].(bool); rec == false {
			return nil
		}
		if err := recursivePush(ctx, api, root, peers); err != nil {
			return err
		}

		return nil
	},
}

func recursivePush(ctx context.Context, api iface.CoreAPI, root string, peers []iface.ConnectionInfo) error {
	node, err := api.ResolveNode(ctx, path.New(root))
	if err != nil {
		return err
	}

	for _, link := range node.Links() {
		for _, peer := range peers {
			if err := api.Block().Push(ctx, 3, peer.ID(), link.Cid); err != nil {
				log.Errorf("push block err, push to node:%s, cid:%s", peer.ID().String(), link.Cid.String())
			}
		}

		if err := recursivePush(ctx, api, link.Cid.String(), peers); err != nil {
			log.Errorf("push block err in recursive push handle, cid:%s", link.Cid.String())
			return err
		}
	}

	return nil
}
