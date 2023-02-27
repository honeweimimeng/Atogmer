package eventdriver

type Event interface {
	name() string
	proto() EventsProto
	attr() map[string]interface{}
}
