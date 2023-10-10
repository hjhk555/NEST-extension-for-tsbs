package arcadetsdb

import (
	"encoding/json"
	"fmt"
	"github.com/timescale/tsbs/pkg/data"
	"io"
)

type serializer struct {
}

func GetInsertDetail(p *data.Point) *InsertDetail {
	// trans p to InsertDetail
	tagMap := make(map[string]string)
	fieldMap := make(map[string]*TimeValue)

	tagKeys := p.TagKeys()
	tagValues := p.TagValues()
	fieldKeys := p.FieldKeys()
	fieldValues := p.FieldValues()

	timestamp := p.Timestamp().UnixNano()
	timeList := make([]int64, 1)
	timeList[0] = timestamp

	for i := range tagKeys {
		if tagValues[i] == nil {
			// null tag skip
			continue
		}
		// string tags only
		tagMap[string(tagKeys[i])] = fmt.Sprint(tagValues[i])
	}

	for i := range fieldKeys {
		value := fieldValues[i]
		if value == nil {
			// null value skip
			continue
		}
		valueList := make([]interface{}, 1)
		valueList[0] = value
		fieldMap[string(fieldKeys[i])] = &TimeValue{
			Timestamps: &timeList,
			Values:     &valueList,
		}
	}

	return &InsertDetail{
		ObjectType: string(p.MeasurementName()),
		Tags:       &tagMap,
		Metrics:    &fieldMap,
	}
}

// write point in format of insert detail
func (s *serializer) Serialize(p *data.Point, w io.Writer) error {
	jsonIns, err := json.Marshal(GetInsertDetail(p))
	if err != nil {
		return err
	}
	_, err = w.Write(append(jsonIns, '\n'))
	return err
}
