package event

import "litecluster/driver"

type StartEventLoop struct {
	id int
}

func (e *StartEventLoop) Name() string {
	return "StartEventLoop"
}

func (e *StartEventLoop) Id() int {
	return e.id
}

func (e *StartEventLoop) Msg() Channel {
	return nil
}

type LoopExecutorRegistry struct {
	child driver.ExecutorCtxProcesses
}

func UseEventLoopExecutor(trigger Trigger) driver.ExecutorCtxProcesses {
	registry := &TriggerRegistry{eventTrigger: trigger}
	return &LoopExecutorRegistry{child: registry}
}

func (r *LoopExecutorRegistry) Process(ctx driver.ExecutorContext) {
	bus, okBus := ctx.Group().(*ExecutorBus)
	if eventCtx, ok := ctx.(*Context); ok && okBus {
		eventLoop := &Executor{event: &RegisterEventExecutor{}, ctx: eventCtx}
		ctx.Group().Join(eventLoop)
		r.child.Process(ctx)
		r.StartEventLoop(bus)
		return
	}
	panic("cannot startUp eventLoop,because ctx is not type of *event.Context")
}

func (r *LoopExecutorRegistry) StartEventLoop(bus *ExecutorBus) {
	go func() {
		bus.eventTrigger.Channel() <- &TriggerManager{directEvent: []AcceptableEvent{&StartEventLoop{id: 0}}}
	}()
}
