package driver

import "litecluster/event"

type EventLooper interface {
	LooperGroup
	Group()
	AcceptedEvent() []event.Proto
	ProcessEvent()
}

type AbstractLooper struct {
}

func (l *AbstractLooper) Next() LooperGroup {
	return nil
}

func (l *AbstractLooper) Registry(channel Channel) {

}

func (l *AbstractLooper) AcceptedEvent() []event.Proto {
	return nil
}

func (l *AbstractLooper) ProcessEvent() {

}
