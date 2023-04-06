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

type ExecutorContext interface {
	Config() *ExecutorConfig
	Group() ExecutorGroup
	SetGroup(group ExecutorGroup)
	Context() context.Context
	Interrupt() context.CancelFunc
	Process() ExecutorProcess
	GroupRule() GroupRule
}

type ExecutorGroup interface {
	Join(executor Executor) ExecutorGroup
	Execute()
	Channel(executor Executor) chan Executor
}

type GroupRule interface {
	Strategy(executors []Executor) []Executor
	Provide(group ExecutorGroup, exes []Executor)
}

type ExecutorProcess interface {
	Context(ctx ExecutorContext)
	Process(executor Executor)
	Group(group ExecutorGroup)
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
