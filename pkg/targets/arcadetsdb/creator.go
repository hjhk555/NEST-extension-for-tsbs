package arcadetsdb

import "github.com/timescale/tsbs/pkg/targets"

type dbCreator struct {
	client *Client
}

func newDBCreator(host string, port int) targets.DBCreator {
	return &dbCreator{client: NewClient(host, port)}
}

func (d *dbCreator) Init() {
	err := d.client.Connect()
	if err != nil {
		panic(err)
	}
}

func (d *dbCreator) DBExists(dbName string) bool {
	ret, err := d.client.SendMsgAndWaitRet(ManageMessage{
		clientMessage: clientMessage{
			Database: dbName,
		},
		ManageType: "exist",
	})
	if err != nil {
		panic(err)
	}
	result, err := ret.GetResult()
	if err != nil {
		panic(err)
	}
	return result.(bool)
}

func (d *dbCreator) CreateDB(dbName string) error {
	ret, err := d.client.SendMsgAndWaitRet(ManageMessage{
		clientMessage: clientMessage{
			Database: dbName,
		},
		ManageType: "create",
	})
	if err != nil {
		return err
	}
	_, err = ret.GetResult()
	return err
}

func (d *dbCreator) RemoveOldDB(dbName string) error {
	ret, err := d.client.SendMsgAndWaitRet(ManageMessage{
		clientMessage: clientMessage{
			Database: dbName,
		},
		ManageType: "drop",
	})
	if err != nil {
		return err
	}
	_, err = ret.GetResult()
	return err
}

func (d *dbCreator) Close() {
	d.client.Close()
}
