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
	//ctx, cancel := context.WithCancel(context.Background())
	//defer cancel()
	//c, _ := cid.Decode(hash)
	//j.cid = c
	//ctx, cancel := context.WithCancel(context.Background())
	//defer cancel()

	// Bootstrappers are using 1024 keys. See:
	// https://github.com/ipfs/infra/issues/378
	crypto.MinRsaKeyBits = 1024

	//root := j.store + string(os.PathSeparator) + repo.Root

	ds, err := ipfslite.BadgerDatastore("/tmp/stest")
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

	//c, _ := cid.Decode("QmWATWQ7fVPP2EFGu71UkfnqhYXDYH566qy47CnJDgvs8u")
	//rsc, err := lite.GetFile(ctx, c)
	//if err != nil {
	//	panic(err)
	//}
	//defer rsc.Close()
	//content, err := ioutil.ReadAll(rsc)
	//if err != nil {
	//	panic(err)
	//}
	//
	//
	j.peer = peer
	return j
}

func (j *JavazacDB) ReadList(hash string) (itms items.I) {
	c, _ := cid.Decode(hash)
	rsc, err := j.peer.Get(j.ctx, c)
	checkError(err)
	for _, item := range rsc.Links() {
		//pss, err := rsc.Stat()
		//checkError(err)
		//nonono, err := item.GetNode(j.ctx, j.peer)
		//checkError(err)
		//nns, err := nonono.Stat()
		//checkError(err)

		itms = append(itms, &items.FolderListItem{
			Name: item.Name,
			Cid:  item.Cid,
			Size: item.Size,
			//Type:  uint8,
			Btn:   new(widget.Clickable),
			Check: new(widget.Bool),
		})
	}
	return
}
