package worker

import (
	"testing"

	blocks "github.com/ipfs/fs-repo-migrations/ipfs-2-to-3/Godeps/_workspace/src/github.com/ipfs/go-ipfs/blocks"
	blockstore "github.com/ipfs/fs-repo-migrations/ipfs-2-to-3/Godeps/_workspace/src/github.com/ipfs/go-ipfs/blocks/blockstore"
	"github.com/ipfs/fs-repo-migrations/ipfs-2-to-3/Godeps/_workspace/src/github.com/ipfs/go-ipfs/exchange/offline"
	ds "github.com/ipfs/fs-repo-migrations/ipfs-2-to-3/Godeps/_workspace/src/github.com/jbenet/go-datastore"
	dssync "github.com/ipfs/fs-repo-migrations/ipfs-2-to-3/Godeps/_workspace/src/github.com/jbenet/go-datastore/sync"
)

func BenchmarkHandle10KBlocks(b *testing.B) {
	bstore := blockstore.NewBlockstore(dssync.MutexWrap(ds.NewMapDatastore()))
	var testdata []*blocks.Block
	for i := 0; i < 10000; i++ {
		testdata = append(testdata, blocks.NewBlock([]byte(string(i))))
	}
	b.ResetTimer()
	b.SetBytes(10000)
	for i := 0; i < b.N; i++ {

		b.StopTimer()
		w := NewWorker(offline.Exchange(bstore), Config{
			NumWorkers:       1,
			ClientBufferSize: 0,
			WorkerBufferSize: 0,
		})
		b.StartTimer()

		for _, block := range testdata {
			if err := w.HasBlock(block); err != nil {
				b.Fatal(err)
			}
		}

		b.StopTimer()
		w.Close()
		b.StartTimer()

	}
}
