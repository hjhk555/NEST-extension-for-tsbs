package arcadetsdb

import "github.com/timescale/tsbs/pkg/targets"

type processor struct {
	dbName string
	client *Client
}

func newProcessor(dbName string, host string, port int) targets.Processor {
	return &processor{
		dbName: dbName,
		client: NewClient(host, port),
	}
}

func (p *processor) Init(_ int, _, _ bool) {
	err := p.client.Connect()
	if err != nil {
		panic(err)
	}
}

func (p *processor) ProcessBatch(b targets.Batch, _ bool) (metricCount, rowCount uint64) {
	insertBatch := b.(*batch)
	ret, err := p.client.SendMsgAndWaitRet(InsertMessage{
		clientMessage: clientMessage{
			Database: p.dbName,
		},
		Strategies: nil,
		Inserts:    insertBatch.inserts,
	})
	if err != nil {
		panic(err)
	}
	if _, err := ret.GetResult(); err != nil {
		panic(err)
	}

	rowCnt := uint64(len(insertBatch.inserts))
	var mtcCnt uint64
	mtcCnt = 0
	for _, insert := range insertBatch.inserts {
		mtcCnt += uint64(len(*(insert.Metrics)))
	}
	return mtcCnt, rowCnt
}

func (p *processor) Close(_ bool) {
	p.client.Close()
}
