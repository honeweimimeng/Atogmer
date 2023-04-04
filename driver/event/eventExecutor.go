package event

import (
	"fmt"
	"litecluster/driver"
)

type AcceptableEvent interface {
	Name() string
	Id() int
	Msg() Channel
}

type RegisterEventExecutor struct {
	id int
}

func (e *RegisterEventExecutor) Name() string {
	return "RegisterEventExecutor"
}

func (e *RegisterEventExecutor) Id() int {
	return e.id
}

func (e *RegisterEventExecutor) Msg() Channel {
	return nil
}

type Executor struct {
	event AcceptableEvent
	child driver.Executor
	ctx   *Context
}

func (e *Executor) Execute() {
	for {
		select {
		case ex := <-e.ctx.Channel(e):
			e.Context().Group().Join(ex)
			break
		}
	}
}

func (e *Executor) Context() driver.ExecutorContext {
	return e.ctx
}

func (e *Executor) Name() string {
	str := fmt.Sprintf("eventLoop At %s-%d", e.event.Name(), e.event.Id())
	return str
}
