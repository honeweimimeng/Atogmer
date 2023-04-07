package driver

import (
	"github.com/honeweimimeng/atogmer/utils/pool"
	"sync"
)

type ExecutorPoll struct {
	exJoin chan Executor
	ctx    ExecutorContext
	pool   pool.Pool
	wait   sync.WaitGroup
}

func NewExecutorPoll(c ExecutorContext) *ExecutorPoll {
	res := &ExecutorPoll{
		ctx:    c,
		exJoin: make(chan Executor, c.Config().ExecutorCap),
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

func (p *ExecutorPoll) WaitFinish() bool {
	p.wait.Wait()
	return true
}

func (p *ExecutorPoll) Channel(executor Executor) chan Executor {
	return nil
}

func (p *ExecutorPoll) Join(executor Executor) ExecutorGroup {
	p.wait.Add(1)
	p.exJoin <- executor
	return p
}

func (p *ExecutorPoll) Execute() {
	go func() {
		for {
			select {
			case ex := <-p.exJoin:
				p.wait.Done()
				task := &ExecutorTask{executor: ex}
				p.pool.Run(task)
			}
		}
	}()
}
