package net

import (
	"litecluster/driver"
	"litecluster/driver/event"
)

type IOHandler struct {
	ctx    driver.ExecutorContext
	name   string
	events []event.Proto
}

func NewIOEventLoop(ctx driver.ExecutorContext, handlers []event.Handler) *event.LoopExecutor {
	return event.NewEventLoop(ctx, handlers)
}

func NewIOHandle(ctx driver.ExecutorContext, events []event.Proto) *IOHandler {
	res := &IOHandler{
		events: events,
		ctx:    ctx,
	}
	for _, e := range res.events {
		res.name = res.name + e.Name() + "@" + string(rune(e.Id()))
	}
	return res
}

func (h *IOHandler) Events() []event.Proto {
	return h.events
}

func (h *IOHandler) Name() string {
	return h.name
}

func (h *IOHandler) Execute() {
	println(h.Name() + "happen")
}

func (h *IOHandler) Context() driver.ExecutorContext {
	return h.ctx
}

type Registry struct {
}

func (r *Registry) Context(ctx driver.ExecutorContext) {
	handlers := []event.Handler{NewIOHandle(ctx, Events(0))}
	ctx.Group().Join(NewIOEventLoop(ctx, handlers).AddTrigger(&EpollTrigger{}))
}

func (r *Registry) Process(executor driver.Executor) {

}

func (r *Registry) Group(group driver.ExecutorGroup) {
}
