package io

import (
	"litecluster/rpc"
	"net"
)

type NIODriver struct {
	Handler func(err error)
	conn    net.Conn
}

func EmptyDriver() *NIODriver {
	return &NIODriver{}
}

func (d *NIODriver) ExceptionHandle() func(err error) {
	return d.Handler
}
func (d *NIODriver) Start(conf *rpc.Conf) *net.Listener {
	return nil
}
func (d *NIODriver) Conn() net.Conn {
	return d.conn
}
