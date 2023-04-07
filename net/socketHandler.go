package net

import (
	"atogmer/driver"
	"atogmer/driver/event"
	"time"
)

type IOHandler struct {
	ctx    driver.ExecutorContext
	name   string
	events []event.Proto
}

func NewIOHandle(ctx driver.ExecutorContext, events []event.Proto) *IOHandler {
	res := &IOHandler{
		events: events,
		ctx:    ctx,
		name:   event.FormatName(events),
	}
	return res
}

func (h *IOHandler) Execute() {
	println(h.Name() + "happen")
	time.Sleep(10 * time.Second)
	println(h.Name() + "happen finish")
}

func (h *IOHandler) Events() []event.Proto {
	return h.events
}

func (h *IOHandler) Name() string {
	return h.name
}

func (h *IOHandler) Context() driver.ExecutorContext {
	return h.ctx
}
