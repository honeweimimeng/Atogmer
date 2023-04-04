package event

import (
	"context"
	"litecluster/driver"
	"litecluster/utils"
)

type Context struct {
	process    driver.ExecutorCtxProcesses
	ctx        context.Context
	group      driver.ExecutorGroup
	boot       driver.Executor
	config     *driver.ExecutorConfig
	cancel     context.CancelFunc
	channelMap *utils.SafeMap[AcceptableEvent, chan driver.Executor]
}

func (c *Context) Config() *driver.ExecutorConfig {
	return c.config
}

func (c *Context) Boot() driver.Executor {
	if c.process != nil {
		c.process.Process(c)
	}
	c.boot = driver.UseBootExecutor(c)
	return c.boot
}

func UseTriggerEventLoop(trigger Trigger) *Context {
	eventLoop := UseEventLoopExecutor(trigger)
	return &Context{process: eventLoop}
}

func (c *Context) Init() *Context {
	ctx, cancel := context.WithCancel(context.Background())
	c.group = UseExecutorBus(UseTriggerManager())
	c.channelMap = utils.UseInterSafeMap[AcceptableEvent, chan driver.Executor]()
	c.config = driver.DefaultConfig()
	c.ctx = ctx
	c.cancel = cancel
	return c
}

func (c *Context) Interrupt() context.CancelFunc {
	return c.cancel
}

func (c *Context) Context() context.Context {
	return c.ctx
}

func (c *Context) Channel(executor driver.Executor) chan driver.Executor {
	evx := executor.(*Executor)
	ch := c.channelMap.Get(evx.event)
	if ch == nil {
		ch = make(chan driver.Executor)
		c.channelMap.Put(evx.event, ch)
	}
	return ch
}

func (c *Context) Group() driver.ExecutorGroup {
	return c.group
}
