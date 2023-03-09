package data

import (
	"io"
	"net"
)

type NetDataBuf struct {
	listen net.Listener
	conn   net.Conn
	bytes  []byte
}

func (n *NetDataBuf) Create() {
	conn, err := n.listen.Accept()
	if err != nil {
		panic(err)
	}
	n.conn = conn
}

func (n *NetDataBuf) Writer0() io.Writer {
	return n.conn
}

func (n *NetDataBuf) Reader0() io.Reader {
	return n.conn
}

func (n *NetDataBuf) Size() int64 {
	return int64(len(n.bytes))
}
