package call

import "net"

type IOChannel interface {
	Write([]byte) int64
	ReadBytes() []byte
	ReadString() string
}

type NIOChannel struct {
	conn    net.Conn
	adapter *DataAdapter
}

func (c *NIOChannel) Write([]byte) int64 {

}

func (c *NIOChannel) ReadBytes() []byte {

}

func (c *NIOChannel) ReadString() string {

}
