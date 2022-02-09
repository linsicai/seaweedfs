package mount

import (
	"context"
	"github.com/hanwen/go-fuse/fuse"
	"github.com/hanwen/go-fuse/v2/fs"
	"syscall"
)

type myNode struct {
	fs.Inode
}

// Node types must be InodeEmbedders
var _ = fs.InodeEmbedder(&myNode{})

// Node types should implement some file system operations, eg. Lookup
var _ = (fs.NodeLookuper)((*myNode)(nil))

func (n *myNode) Lookup(ctx context.Context, name string, out *fuse.EntryOut) (*fs.Inode, syscall.Errno) {
	ops := myNode{}
	return n.NewInode(ctx, &ops, fs.StableAttr{Mode: syscall.S_IFDIR}), 0
}
