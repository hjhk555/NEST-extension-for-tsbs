package arcadetsdb

import (
	"encoding/json"
	"fmt"
)

type IClientMessage interface {
	GetJsonMessage() ([]byte, error)
}

type clientMessage struct {
	Database string `json:"db,omitempty"`
	Action   string `json:"act"`
}

type ManageMessage struct {
	clientMessage
	ManageType string `json:"typ"`
}

type InsertMessage struct {
	clientMessage
	Strategies []*InsertStrategy `json:"stg,omitempty"`
	Inserts    []*InsertDetail   `json:"ins"`
}

type InsertStrategy struct {
	StrategyType string `json:"typ"`
	Separator    string `json:"sep"`
}

type InsertDetail struct {
	ObjectType string                 `json:"obj"`
	Tags       *map[string]string     `json:"tag,omitempty"`
	Metrics    *map[string]*TimeValue `json:"ts"`
}

type TimeValue struct {
	Timestamps *[]int64       `json:"t"`
	Values     *[]interface{} `json:"v"`
}

type QueryMessage struct {
	clientMessage
	Query *QueryDetail `json:"qry"`
}

type QueryDetail struct {
	ObjectType string            `json:"obj"`
	Tags       map[string]string `json:"tag,omitempty"`
	QueryType  string            `json:"typ"`
	Multiple   bool              `json:"mlt,omitempty"`
	MetricName string            `json:"mtc,omitempty"`
	BeginTime  int64             `json:"bgn,omitempty"`
	EndTime    int64             `json:"end,omitempty"`
	Limit      int               `json:"lmt,omitempty"`
}

type ServerMessage struct {
	Success bool        `json:"suc"`
	Result  interface{} `json:"res"`
	Error   *Error      `json:"err"`
}

type Error struct {
	ClassName string `json:"c"`
	Message   string `json:"m"`
}

func (m ManageMessage) GetJsonMessage() ([]byte, error) {
	m.Action = "manage"
	return json.Marshal(m)
}

func (i InsertMessage) GetJsonMessage() ([]byte, error) {
	i.Action = "insert"
	return json.Marshal(i)
}

func (q QueryMessage) GetJsonMessage() ([]byte, error) {
	q.Action = "query"
	return json.Marshal(q)
}

func (s *ServerMessage) GetResult() (interface{}, error) {
	if s.Success {
		res := s.Result
		if res == nil {
			return true, nil
		}
		return res, nil
	} else {
		return nil, fmt.Errorf("%s:%s", s.Error.ClassName, s.Error.Message)
	}
}
