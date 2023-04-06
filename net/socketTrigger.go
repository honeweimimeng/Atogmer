package net

import (
	"litecluster/driver"
	"litecluster/driver/event"
	"net"
	"time"
)

type Registry struct{}

func (r *Registry) Context(ctx driver.ExecutorContext) {
	handlers := []event.Handler{NewIOHandle(ctx, Events(0))}
	ctx.Group().Join(event.NewEventLoop(ctx, handlers).AddTrigger(&EpollTrigger{ctx: ctx}))
}

func (r *Registry) Process(executor driver.Executor) {}

func (r *Registry) Group(group driver.ExecutorGroup) {}

type EpollTrigger struct {
	ctx      driver.ExecutorContext
	next     event.Trigger
	listener net.Listener
	socketFd uintptr
	epollFd  uintptr
}

func (e *EpollTrigger) AcceptEvents(ch chan []event.Proto) {
	println("===>start Epoll")
	time.Sleep(2 * time.Second)
	ch <- []event.Proto{&READ{id: 0}}
	time.Sleep(2 * time.Second)
	ch <- []event.Proto{&WRITE{id: 0}}
}

func (e *EpollTrigger) Next() event.Trigger {
	return e.next
}

func (e *EpollTrigger) Child(trigger event.Trigger) {
	e.next = trigger
}

func (e *EpollTrigger) Exception(err error) {
	if err != nil {
		e.ctx.Config().Logger.Panicln("load socket err:", err.Error())
	}
}
