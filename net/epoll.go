package net

import "litecluster/driver/event"

type EpollTrigger struct {
	next event.Trigger
}

func (e *EpollTrigger) AcceptEvents(ch chan []event.Proto) {
	println("===>start Epoll")
}

func (e *EpollTrigger) Next() event.Trigger {
	return e.next
}

func (e *EpollTrigger) Child(trigger event.Trigger) {
	e.next = trigger
}
