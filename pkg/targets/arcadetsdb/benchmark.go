package arcadetsdb

import (
	"github.com/timescale/tsbs/pkg/targets"
)

type benchmark struct {
}

func (b benchmark) GetDataSource() targets.DataSource {
	//TODO implement me
	panic("implement me")
}

func (b benchmark) GetBatchFactory() targets.BatchFactory {
	//TODO implement me
	panic("implement me")
}

func (b benchmark) GetPointIndexer(maxPartitions uint) targets.PointIndexer {
	//TODO implement me
	panic("implement me")
}

func (b benchmark) GetProcessor() targets.Processor {
	//TODO implement me
	panic("implement me")
}

func (b benchmark) GetDBCreator() targets.DBCreator {
	//TODO implement me
	return &dbCreator{}
}
