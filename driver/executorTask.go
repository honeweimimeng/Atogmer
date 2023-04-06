package driver

import (
	"context"
	"github.com/sirupsen/logrus"
)

type ExecutorTask struct {
	executor Executor
}

func (e *ExecutorTask) Run() {
	e.Logger().Println("event executor task start, Name:", e.executor.Name())
	e.executor.Execute()
	e.Logger().Println("event executor task finish, Name:", e.executor.Name())
}

func (e *ExecutorTask) Ctx() context.Context {
	return e.executor.Context().Context()
}

func (e *ExecutorTask) Interrupt() context.CancelFunc {
	e.Logger().Println("event executor task Interrupt Name:", e.executor.Name())
	return e.executor.Context().Interrupt()
}

func (e *ExecutorTask) Logger() logrus.StdLogger {
	return e.executor.Context().Config().Logger
}
