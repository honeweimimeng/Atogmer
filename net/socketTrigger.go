package net

import (
	"github.com/honeweimimeng/atogmer/driver"
	"github.com/honeweimimeng/atogmer/driver/event"
	"sync"
)

type Registry struct {
}

func (r *Registry) Context(ctx driver.ExecutorContext) {
	handlers := []event.Handler{NewIOHandle(ctx, Events(0))}
	ctx.Group().Join(event.NewEventLoop(ctx, handlers).AddTrigger(&EpollTrigger{ctx: ctx}))
}

func (r *Registry) Process(executor driver.Executor) {}

func (r *Registry) Group(group driver.ExecutorGroup) {}

type EpollTrigger struct {
	ctx  driver.ExecutorContext
	next event.Trigger
	once sync.Once
}

func (e *EpollTrigger) AcceptEvents(ch chan []event.Proto) {
	e.once.Do(e.Listener)
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

func (e *EpollTrigger) Listener() {
	//fd, err := syscall.Socket(syscall.AF_INET, syscall.O_NONBLOCK|syscall.SOCK_STREAM, 0)
	//if err != nil {
	//	e.ctx.Config().Logger.Panicln("listener socket err,", err.Error())
	//}
}
