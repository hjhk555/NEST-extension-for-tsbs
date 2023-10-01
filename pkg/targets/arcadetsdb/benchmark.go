package arcadetsdb

import (
	"github.com/timescale/tsbs/internal/inputs"
	"github.com/timescale/tsbs/pkg/data/source"
	"github.com/timescale/tsbs/pkg/targets"
)

type benchmark struct {
	dbName     string
	host       string
	port       int
	dataSource targets.DataSource
}

func newBenchmark(targetDB string, dataSourceConfig *source.DataSourceConfig, specificConfig *SpecificConfig) (targets.Benchmark, error) {
	var ds targets.DataSource
	if dataSourceConfig.Type == source.FileDataSourceType {
		panic("FILE data source not supported yet")
	} else {
		dataGenerator := &inputs.DataGenerator{}
		simulator, err := dataGenerator.CreateSimulator(dataSourceConfig.Simulator)
		if err != nil {
			return nil, err
		}
		ds = newSimulationDataSource(simulator)
	}
	return benchmark{
		dbName:     targetDB,
		host:       specificConfig.host,
		port:       specificConfig.port,
		dataSource: ds,
	}, nil
}

func (b benchmark) GetDataSource() targets.DataSource {
	return b.dataSource
}

func (b benchmark) GetBatchFactory() targets.BatchFactory {
	return newBatchFactory()
}

func (b benchmark) GetPointIndexer(maxPartitions uint) targets.PointIndexer {
	//TODO look into point indexer
	return &targets.ConstantIndexer{}
}

func (b benchmark) GetProcessor() targets.Processor {
	return newProcessor(b.dbName, b.host, b.port)
}

func (b benchmark) GetDBCreator() targets.DBCreator {
	return newDBCreator(b.host, b.port)
}
