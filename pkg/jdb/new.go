package jdb

import (
	"context"
	_ "gioui.org/app/permission/storage"
	"gioui.org/widget"
	"github.com/gioapp/cms/pkg/items"
	ipfslite "github.com/hsanjuan/ipfs-lite"
	"github.com/ipfs/go-cid"
	"github.com/libp2p/go-libp2p-core/crypto"
	"github.com/multiformats/go-multiaddr"
)

// JavazacDb Structure
type JavazacDB struct {
	ctx   context.Context
	peer  *ipfslite.Peer
	index map[string]string
	store string
}

func New(ctx context.Context, store string) *JavazacDB {
	j := &JavazacDB{
		ctx:   ctx,
		index: make(map[string]string),
		store: store,
	}
	//log.SetLogLevel("*", "warn")

	crypto.MinRsaKeyBits = 1024
	ds, err := ipfslite.BadgerDatastore("datastore")
	if err != nil {
		panic(err)
	}
	priv, _, err := crypto.GenerateKeyPair(crypto.RSA, 2048)
	if err != nil {
		panic(err)
	}
	listen, _ := multiaddr.NewMultiaddr("/ip4/0.0.0.0/tcp/4005")

	h, dht, err := ipfslite.SetupLibp2p(
		ctx,
		priv,
		nil,
		[]multiaddr.Multiaddr{listen},
		ds,
		ipfslite.Libp2pOptionsExtra...,
	)

	if err != nil {
		panic(err)
	}
	peer, err := ipfslite.New(ctx, ds, h, dht, nil)
	if err != nil {
		panic(err)
	}
	peer.Bootstrap(ipfslite.DefaultBootstrapPeers())

	j.peer = peer
	return j
}

func (j *JavazacDB) ReadList(hash string) (itms items.I) {
	c, _ := cid.Decode(hash)

	//fmt.Println("nnnnnnnnnnhashhashhashhash--------><", hash)

	node, err := j.peer.Get(j.ctx, c)
	if err != nil {
		panic(err)
	}
	//fmt.Println("nnnnnnnnnn--------><", node)

	n := node.Links()
	for i := 0; i < len(n); i++ {
		itms = append(itms, &items.FolderListItem{
			Name: n[i].Name,
			Cid:  n[i].Cid,
			Size: n[i].Size,
			//Type:  uint8,
			Btn:   new(widget.Clickable),
			Check: new(widget.Bool),
		})
	}
	return
}
