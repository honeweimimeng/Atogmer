package event

type Trigger interface {
	AcceptEvent() []AcceptableEvent
	Channel() chan Trigger
	Next() Trigger
	Child(trigger Trigger)
	Await()
}

type TriggerManager struct {
	directEvent []AcceptableEvent
	resChannel  chan []AcceptableEvent
	next        Trigger
}

func UseTriggerManager() *TriggerManager {
	return &TriggerManager{
		resChannel: make(chan []AcceptableEvent),
	}
}

func (m *TriggerManager) AcceptEvent() []AcceptableEvent {
	if m.directEvent != nil {
		return m.directEvent
	}
	m.Await()
	return <-m.resChannel
}

func (m *TriggerManager) Await() {
	go func() {
		for {
			select {
			case r := <-m.Channel():
				m.resChannel <- r.AcceptEvent()
			}
		}
	}()
}

func (m *TriggerManager) Channel() chan Trigger {
	return m.next.Channel()
}

func (m *TriggerManager) Next() Trigger {
	return m.next
}

func (m *TriggerManager) Child(trigger Trigger) {
	m.next = trigger
}
