package driver

import (
	"context"
	"litecluster/utils"
)

type EventContext struct {
	ctx        context.Context
	group      ExecutorGroup
	boot       Executor
	config     *ExecutorConfig
	cancel     context.CancelFunc
	channelMap *utils.SafeMap[AcceptableEvent, chan Executor]
}

func (c *EventContext) Config() *ExecutorConfig {
	return c.config
}

func (c *EventContext) Boot() Executor {
	c.boot = UseBootExecutor(c)
	return c.boot
}

func UseEventContext() *EventContext {
	ctx, cancel := context.WithCancel(context.Background())
	return &EventContext{
		group:      UseExecutorBus(UseTriggerManager()),
		channelMap: utils.UseInterSafeMap[AcceptableEvent, chan Executor](),
		config:     DefaultConfig(),
		ctx:        ctx,
		cancel:     cancel,
	}
}

func (c *EventContext) Interrupt() context.CancelFunc {
	return c.cancel
}

func (c *EventContext) Context() context.Context {
	return c.ctx
}

func (c *EventContext) Channel(executor Executor) chan Executor {
	evx := executor.(*EventExecutor)
	ch := c.channelMap.Get(evx.event)
	if ch == nil {
		ch = make(chan Executor)
		c.channelMap.Put(evx.event, ch)
	}
	return ch
}

func (c *EventContext) Group() ExecutorGroup {
	return c.group
}
