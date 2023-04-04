package net

import (
	"litecluster/driver/event"
)

type SocketTrigger struct {
	adviceChannel chan event.Trigger
	child         event.Trigger
}

func (t *SocketTrigger) AcceptEvent() []event.AcceptableEvent {
	res := make([]event.AcceptableEvent, 0)
	return res
}

func (t *SocketTrigger) Channel() chan event.Trigger {
	if t.adviceChannel == nil {
		t.adviceChannel = make(chan event.Trigger)
	}
	return t.adviceChannel
}

func (t *SocketTrigger) Await() {

}

func (t *SocketTrigger) Next() event.Trigger {
	return t.child
}

func (t *SocketTrigger) Child(trigger event.Trigger) {
	t.child = trigger
}
