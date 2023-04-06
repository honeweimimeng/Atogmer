package event

import (
	"context"
	"litecluster/driver"
)

type Context struct {
	process   driver.ExecutorProcess
	config    *driver.ExecutorConfig
	ctx       context.Context
	group     driver.ExecutorGroup
	cancel    context.CancelFunc
	groupRule driver.GroupRule
}

func UseEventContext() *Context {
	ctx, cancel := context.WithCancel(context.Background())
	res := &Context{}
	return res.loadProperty(driver.DefaultConfig(), ctx, cancel)
}

func Process(process driver.ExecutorProcess) *Context {
	ctx, cancel := context.WithCancel(context.Background())
	res := &Context{process: process, ctx: ctx, cancel: cancel}
	return res
}

func ConfigEventContext(config *driver.ExecutorConfig) *Context {
	ctx, cancel := context.WithCancel(context.Background())
	res := &Context{}
	return res.loadProperty(config, ctx, cancel)
}

func Group(group driver.ExecutorGroup) *Context {
	ctx, cancel := context.WithCancel(context.Background())
	res := &Context{group: group, ctx: ctx, cancel: cancel}
	return res
}

func (c *Context) loadProperty(config *driver.ExecutorConfig,
	ctx context.Context, cancel context.CancelFunc) *Context {
	res := &Context{ctx: ctx, cancel: cancel}
	res.LoadProperty(config)
	return res.LoadProperty(config)
}

func (c *Context) LoadProperty(config *driver.ExecutorConfig) *Context {
	c.config = config
	if c.groupRule == nil {
		c.groupRule = &DefaultLoopStrategy{}
	}
	return c
}

func (c *Context) GroupRule() driver.GroupRule {
	return c.groupRule
}

func (c *Context) GroupRul(rule driver.GroupRule) {
	c.groupRule = rule
}

func (c *Context) SetGroup(group driver.ExecutorGroup) {
	c.group = group
}

func (c *Context) Process() driver.ExecutorProcess {
	return c.process
}

func (c *Context) Config() *driver.ExecutorConfig {
	return c.config
}

func (c *Context) Group() driver.ExecutorGroup {
	return c.group
}

func (c *Context) Context() context.Context {
	return c.ctx
}

func (c *Context) Interrupt() context.CancelFunc {
	return c.cancel
}
