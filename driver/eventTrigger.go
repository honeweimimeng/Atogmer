package driver

type EventTrigger interface {
	AcceptEvent() []AcceptableEvent
	NextTrigger(trigger EventTrigger)
	Trig(trigger EventTrigger)
	Await()
}

type EventTriggerManager struct {
	resChannel  chan []AcceptableEvent
	trigChannel chan EventTrigger
	next        EventTrigger
}

func UseTriggerManager() *EventTriggerManager {
	return &EventTriggerManager{
		resChannel:  make(chan []AcceptableEvent),
		trigChannel: make(chan EventTrigger),
	}
}

func (m *EventTriggerManager) NextTrigger(trigger EventTrigger) {
	m.next = trigger
}

func (m *EventTriggerManager) AcceptEvent() []AcceptableEvent {
	m.Await()
	return <-m.resChannel
}

func (m *EventTriggerManager) Trig(trigger EventTrigger) {
	m.trigChannel <- trigger
}

func (m *EventTriggerManager) Await() {
	go func() {
		for {
			select {
			case r := <-m.trigChannel:
				m.resChannel <- r.AcceptEvent()
			}
		}
	}()
}
