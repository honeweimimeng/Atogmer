package event

import "atogmer/driver"

type DefaultLoopStrategy struct {
	result []driver.Executor
}

func (l *DefaultLoopStrategy) Strategy(executors []driver.Executor) []driver.Executor {
	res := make([]driver.Executor, 0)
	mergeHandle := make([]driver.Executor, 0)
	for _, item := range executors {
		handle, ok := item.(Handler)
		if ok {
			if len(handle.Events()) > 1 {
				res = append(res, item)
				continue
			}
			mergeHandle = append(mergeHandle, item)
		}
	}
	l.result = res
	return l.mergeHandle(mergeHandle)
}

func (l *DefaultLoopStrategy) mergeHandle(merge []driver.Executor) []driver.Executor {
	return l.result
}

func (l *DefaultLoopStrategy) Provide(group driver.ExecutorGroup, exes []driver.Executor) {
	for _, item := range exes {
		group.Join(item)
	}
}
