package rpc

import "net"

type Conf struct {
	Ip       string
	Port     string
	WaitTime int64
	LongConn bool
	Stream   bool
}

type Driver interface {
	ExceptionHandle() func(err error)
	Start(conf *Conf) *net.Listener
	Conn() net.Conn
}

type Reader interface {
}

type Writer interface {
}
