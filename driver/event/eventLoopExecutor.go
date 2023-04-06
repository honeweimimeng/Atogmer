package event

import (
	"litecluster/driver"
	"litecluster/utils"
)

type Proto interface {
	Name() string
	Id() int
}

type Handler interface {
	driver.Executor
	Events() []Proto
}

type LoopExecutor struct {
	name    string
	events  []Handler
	context driver.ExecutorContext
	sel     *utils.MultiCaseSel[driver.Executor]
}

func NewEventLoop(ctx driver.ExecutorContext, events []Handler) *LoopExecutor {
	return &LoopExecutor{
		name:    ctx.Config().Name,
		events:  events,
		context: ctx,
		sel:     utils.NewMulti[driver.Executor](ctx.Config().Name, ctx.Context(), ctx.Config().Logger),
	}
}

func (l *LoopExecutor) AddTrigger(trigger Trigger) *LoopExecutor {
	bus := l.context.Group().(*Bus)
	bus.AddTrigger(trigger)
	return l
}

func (l *LoopExecutor) Name() string {
	return l.name
}

func (l *LoopExecutor) Execute() {
	for _, item := range l.events {
		l.sel.ChannelHandler(l.context.Group().Channel(item), func(ex driver.Executor) {
			println("===>happen")
		})
	}
	l.sel.Start()
}

func (l *LoopExecutor) Context() driver.ExecutorContext {
	return l.context
}
