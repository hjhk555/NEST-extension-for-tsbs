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
	// trans point to InsertDetail
	tagMap := make(map[string]string)
	fieldMap := make(map[string]*TimeValue)

	tagKeys := point.TagKeys()
	tagValues := point.TagValues()
	fieldKeys := point.FieldKeys()
	fieldValues := point.FieldValues()

	timestamp := point.Timestamp().UnixNano()
	timeList := make([]int64, 1)
	timeList[0] = timestamp

	for i := range tagKeys {
		tagMap[string(tagKeys[i])] = tagValues[i].(string)
	}
	for i := range fieldKeys {
		valueList := make([]interface{}, 1)
		valueList[0] = fieldValues[i]
		fieldMap[string(fieldKeys[i])] = &TimeValue{
			Timestamps: &timeList,
			Values:     &valueList,
		}
	}
	return data.LoadedPoint{
		Data: &InsertDetail{
			ObjectType: string(point.MeasurementName()),
			Tags:       &tagMap,
			Metrics:    &fieldMap,
		},
	}
}

func (s *simulationDataSource) Headers() *common.GeneratedDataHeaders {
	return s.simulator.Headers()
}
