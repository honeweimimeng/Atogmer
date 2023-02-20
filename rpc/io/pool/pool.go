package pool

import (
	"litecluster/rpc"
	"litecluster/rpc/meta"
	"net"
)

type EventPool struct {
	listen  net.Listener
	connMap map[string]meta.ConnMeta
	bus     rpc.EventBus
	config  *EventPoolConfig
}

func (p *EventPool) InitPool() bool {
	p.connMap = make(map[string]meta.ConnMeta)

	return true
}

func (p *EventPool) Product(e rpc.Event) int {
	return 0
}

func (p *EventPool) Consumer(handler *rpc.Executor, e *rpc.Event) {

}

type BossEventPool struct {

}

type