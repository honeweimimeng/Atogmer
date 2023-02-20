package io

import (
	"litecluster/rpc"
	"litecluster/rpc/io/pool"
)

type EventBootStrap interface {
	Run(driver rpc.Driver)
	Close()
}

type EventLooper struct {
	pool.EventPool
	executor *rpc.Executor
}

func (l *EventLooper) Run(driver rpc.Driver) {
	e := &rpc.Event{Driver: driver}
	l.InitPool()
	l.PublishEvent(e)
}

func (l *EventLooper) StartUp() {
	l.config.ParseConfig()
	l.Consumer(l.executor, rpc.Use(rpc.ListenerTopic))
}

func (l *EventLooper) PublishEvent(e *rpc.Event) {

}

func (l *EventLooper) Close() {

}
