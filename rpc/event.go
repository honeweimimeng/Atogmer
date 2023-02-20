package rpc

type Event struct {
	topic  string
	Driver Driver
}

func Use(topic string) *Event {
	return &Event{topic: topic}
}

const ListenerTopic = "listener"
