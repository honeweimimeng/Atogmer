package driver

import (
	"context"
	"github.com/sirupsen/logrus"
)

type Executor interface {
	Name() string
	Execute()
	Context() ExecutorContext
}

type ExecutorGroup interface {
	Join(executor Executor)
	AcceptExecutors() []Executor
}

type ExecutorContext interface {
	Config() *ExecutorConfig
	Boot() Executor
	Context() context.Context
	Group() ExecutorGroup
	Channel(executor Executor) chan Executor
	Interrupt() context.CancelFunc
}

type ExecutorCtxProcesses interface {
	Process(ctx ExecutorContext)
}

type ExecutorDelegate struct {
	exe    Executor
	DeName string
	Ctx    ExecutorContext
}

func (de *ExecutorDelegate) Name() string {
	if de.exe == nil {
		return de.DeName
	}
	return de.exe.Name()
}

func (de *ExecutorDelegate) Execute() {
	de.exe.Execute()
}

func (de *ExecutorDelegate) Context() ExecutorContext {
	if de.exe == nil {
		return de.Ctx
	}
	return de.exe.Context()
}

type ExecutorConfig struct {
	ExecutorCap uint32
	Name        string
	Logger      logrus.StdLogger
}

func DefaultConfig() *ExecutorConfig {
	return &ExecutorConfig{
		ExecutorCap: 1000,
		Name:        "exConfig0",
		Logger:      logrus.New(),
	}
}
