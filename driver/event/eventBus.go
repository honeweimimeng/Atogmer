package event

import (
	"litecluster/driver"
	"litecluster/utils"
	"reflect"
)

type Bus struct {
	group            driver.ExecutorGroup
	trigger          Trigger
	adviceCh         chan []Proto
	ctx              driver.ExecutorContext
	registeredHandle []driver.Executor
	handleChannel    *utils.SafeMap[Handler, chan driver.Executor]
	eventHandle      *utils.SafeMap[reflect.Type, Handler]
}

func UseEventBus(ctx driver.ExecutorContext) *Bus {
	res := &Bus{
		registeredHandle: make([]driver.Executor, 0),
		group:            driver.NewExecutorPoll(ctx),
		trigger:          NewTriggerManager(ctx),
		adviceCh:         make(chan []Proto),
		ctx:              ctx,
		handleChannel:    utils.NewSafeMap[Handler, chan driver.Executor](),
		eventHandle:      utils.NewSafeMap[reflect.Type, Handler](),
	}
	ctx.SetGroup(res)
	p := ctx.Process()
	if p != nil {
		p.Context(ctx)
	}
	return res
}

func (b *Bus) Join(executor driver.Executor) driver.ExecutorGroup {
	return b.group.Join(executor)
}

func (b *Bus) Execute() {
	b.group.Execute()
	b.startEventLoop(b.ctx.GroupRule().Strategy(b.registeredHandle))
}

func (b *Bus) startEventLoop(executors []driver.Executor) {
	b.trigger.AcceptEvents(b.adviceCh)
	b.ctx.GroupRule().Provide(b.ctx.Group(), executors)
	go func() {
		for {
			select {
			case acceptedEvents := <-b.adviceCh:
				b.processEvents(acceptedEvents)
				break
			}
		}
	}()
}

func (b *Bus) registerEventHandler(handler Handler) *Bus {
	for _, item := range handler.Events() {
		b.eventHandle.Put(reflect.TypeOf(item), handler)
	}
	b.registeredHandle = append(b.registeredHandle, handler)
	return b
}

func (b *Bus) Channel(executor driver.Executor) chan driver.Executor {
	handler, ok := executor.(Handler)
	if ok {
		ch := b.handleChannel.Get(handler)
		if ch == nil {
			ch = make(chan driver.Executor)
			b.handleChannel.Put(handler, ch)
		}
		b.registerEventHandler(handler)
		return ch
	}
	b.ctx.Config().Logger.Panicln("cannot init channel,because executor not type of handler")
	return nil
}

func (b *Bus) AddTrigger(trigger Trigger) *Bus {
	b.trigger.Child(trigger)
	return b
}

func (b *Bus) processEvents(events []Proto) {
	for _, item := range events {
		ex := b.eventHandle.Get(reflect.TypeOf(item))
		b.ctx.Group().Channel(ex) <- ex
	}
}
