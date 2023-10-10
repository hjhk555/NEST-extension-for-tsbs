package arcadetsdb

import (
	"github.com/timescale/tsbs/pkg/data"
	"github.com/timescale/tsbs/pkg/data/usecases/common"
	"github.com/timescale/tsbs/pkg/targets"
)

type simulationDataSource struct {
	simulator common.Simulator
}

func newSimulationDataSource(simulator common.Simulator) targets.DataSource {
	return &simulationDataSource{simulator: simulator}
}

func (s *simulationDataSource) NextItem() data.LoadedPoint {
	point := data.NewPoint()
	var success bool
	for !s.simulator.Finished() {
		if success = s.simulator.Next(point); success {
			break
		}
		point.Reset()
	}
	if !success && s.simulator.Finished() {
		// no more points
		return data.LoadedPoint{}
	}

	return data.LoadedPoint{
		Data: GetInsertDetail(point),
	}
}

func (s *simulationDataSource) Headers() *common.GeneratedDataHeaders {
	panic("arcadetsdb contains no headers")
}
