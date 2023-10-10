package arcadetsdb

import (
	"bufio"
	"encoding/json"
	"github.com/timescale/tsbs/load"
	"github.com/timescale/tsbs/pkg/data"
	"github.com/timescale/tsbs/pkg/data/usecases/common"
	"github.com/timescale/tsbs/pkg/targets"
)

type fileDataSource struct {
	scanner *bufio.Scanner
}

func (f fileDataSource) NextItem() data.LoadedPoint {
	f.scanner.Scan()
	if err := f.scanner.Err(); err != nil {
		panic(err)
	}

	ins := &InsertDetail{}
	err := json.Unmarshal(f.scanner.Bytes(), ins)
	if err != nil {
		panic(err)
	}
	return data.LoadedPoint{Data: ins}
}

func (f fileDataSource) Headers() *common.GeneratedDataHeaders {
	panic("arcadetsdb contains no headers")
}

func newFileDataSource(path string) targets.DataSource {
	return &fileDataSource{scanner: bufio.NewScanner(load.GetBufferedReader(path))}
}
