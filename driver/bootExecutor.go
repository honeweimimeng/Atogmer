package driver

import (
	"context"
	"github.com/sirupsen/logrus"
	"litecluster/utils/pool"
)

type BootExecutor struct {
	ctx  ExecutorContext
	pool pool.Pool
}

func UseBootExecutor(c ExecutorContext) *BootExecutor {
	res := &BootExecutor{
		ctx: c,
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

func (e *BootExecutor) Execute() {
	for _, ex := range e.Context().Group().AcceptExecutors() {
		task := &BootExecutorTask{executor: ex}
		e.pool.Run(task)
	}
}

func (e *BootExecutor) Context() ExecutorContext {
	return e.ctx
}

func (e *BootExecutor) Name() string {
	return ""
}

type BootExecutorTask struct {
	executor Executor
}

func (e *BootExecutorTask) Run() {
	e.Logger().Println("event executor task start Name:", e.executor.Name())
	e.executor.Execute()
	e.executor.Context().Channel(e.executor) <- e.executor
	e.Logger().Println("event executor task finish Name:", e.executor.Name())
}

func (e *BootExecutorTask) Ctx() context.Context {
	return e.executor.Context().Context()
}

func (e *BootExecutorTask) Interrupt() context.CancelFunc {
	e.Logger().Println("event executor task Interrupt Name:", e.executor.Name())
	return e.executor.Context().Interrupt()
}

func (e *BootExecutorTask) Logger() logrus.StdLogger {
	return e.executor.Context().Config().Logger
}
