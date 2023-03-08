package eventdriver

type EventBus interface {
	Publish(event Event)
	Registry(proto EventsProto)
	RegistryHandle(event Event, handle EventHandle)
}

type CycleEventBus struct {
	EventLooper
}

func (c *CycleEventBus) Publish(event Event) {

}
func (c *CycleEventBus) Registry(proto EventsProto) {

}
func (c *CycleEventBus) RegistryHandle(event Event, handle EventHandle) {

}
