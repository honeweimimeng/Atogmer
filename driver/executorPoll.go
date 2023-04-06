package driver

import (
	"litecluster/utils/pool"
)

type ExecutorPoll struct {
	exJoin chan Executor
	ctx    ExecutorContext
	pool   pool.Pool
}

func NewExecutorPoll(c ExecutorContext) *ExecutorPoll {
	res := &ExecutorPoll{
		ctx:    c,
		exJoin: make(chan Executor, c.Config().LoopCap),
		pool: &pool.DefaultPool{
			Cap_:  c.Config().ExecutorCap,
			Name_: c.Config().Name,
			Ctx:   c.Context(),
			Pipe:  pool.GetFifoPipe(c.Config().ExecutorCap),
		},
	}
	res.pool.StartUp()
	return res
}

func (p *ExecutorPoll) Channel(executor Executor) chan Executor {
	return nil
}

func (p *ExecutorPoll) Join(executor Executor) ExecutorGroup {
	p.exJoin <- executor
	return p
}

func (p *ExecutorPoll) Execute() {
	go func() {
		select {
		case ex := <-p.exJoin:
			task := &ExecutorTask{executor: ex}
			p.pool.Run(task)
		}
	}()
}
