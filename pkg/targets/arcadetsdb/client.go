package arcadetsdb

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net"
	"strconv"
)

const (
	NETWORK         = "tcp"
	DEFAULT_HOST_IP = "localhost"
	DEFAULT_PORT    = 8809
)

type Client struct {
	hostIP string
	port   int
	conn   net.Conn
	reader *bufio.Reader
}

func NewClientDefault() *Client {
	return NewClient(DEFAULT_HOST_IP, DEFAULT_PORT)
}

func NewClient(hostIP string, port int) *Client {
	return &Client{
		hostIP: hostIP,
		port:   port,
		conn:   nil,
		reader: nil,
	}
}

func (c *Client) Connect() error {
	conn, err := net.Dial(NETWORK, c.hostIP+":"+strconv.Itoa(c.port))
	if err != nil {
		return err
	}
	c.conn = conn
	c.reader = bufio.NewReader(conn)
	return nil
}

func (c *Client) SendMsgAndWaitRet(msg IClientMessage) (ret *ServerMessage, err error) {
	if c.conn == nil {
		return nil, fmt.Errorf("null connection")
	}

	jsonMsg, err := msg.GetJsonMessage()
	if err != nil {
		return nil, err
	}

	return c.SendBytesAndWaitRet(jsonMsg)
}

func (c *Client) SendBytesAndWaitRet(bytes []byte) (ret *ServerMessage, err error) {
	_, err = c.conn.Write(append(bytes, '\n'))
	if err != nil {
		return nil, err
	}

	line, err := c.reader.ReadBytes('\n')
	if err != nil {
		return nil, err
	}

	var retMsg ServerMessage
	err = json.Unmarshal(line, &retMsg)
	if err != nil {
		return nil, err
	}

	return &retMsg, nil
}

func (c *Client) Close() {
	if c.conn == nil {
		return
	}
	c.conn.Write([]byte("close\n"))
	c.conn.Close()
}
