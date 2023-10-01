package arcadetsdb

import (
	"github.com/timescale/tsbs/pkg/data"
	"github.com/timescale/tsbs/pkg/targets"
	"sync"
)

type batch struct {
	inserts []*InsertDetail
}

func (b *batch) Len() uint {
	return uint(len(b.inserts))
}

func (b *batch) Append(point data.LoadedPoint) {
	insert := point.Data.(*InsertDetail)
	b.inserts = append(b.inserts, insert)
}

func (b *batch) Clear() {
	b.inserts = b.inserts[:0]
}

func newBatch() targets.Batch {
	return &batch{}
}

type batchFactory struct {
	batchPool *sync.Pool
}

func (b *batchFactory) New() targets.Batch {
	batch := b.batchPool.Get().(*batch)
	batch.Clear()
	return batch
}

func newBatchFactory() targets.BatchFactory {
	return &batchFactory{
		batchPool: &sync.Pool{
			New: func() interface{} {
				return newBatch()
			},
		},
	}
}
