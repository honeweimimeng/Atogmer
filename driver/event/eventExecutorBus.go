package event

import (
	"litecluster/driver"
	"litecluster/utils"
)

type ExecutorBus struct {
	eventTrigger   Trigger
	executorAdvice chan []driver.Executor
	eventMap       *utils.SafeMap[int, driver.Executor]
}

func UseExecutorBus(trigger Trigger) *ExecutorBus {
	return &ExecutorBus{
		eventMap:       utils.UseInterSafeMap[int, driver.Executor](),
		eventTrigger:   trigger,
		executorAdvice: make(chan []driver.Executor),
	}
}

func (bus *ExecutorBus) Join(executor driver.Executor) {
	eventExe := executor.(*Executor)
	bus.eventMap.Put(eventExe.event.Id(), executor)
}

func (bus *ExecutorBus) AcceptExecutors() []driver.Executor {
	events := bus.eventTrigger.AcceptEvent()
	res := make([]driver.Executor, len(events))
	for i, e := range events {
		ex := bus.eventMap.Get(e.Id())
		if ex == nil {
			ex = &driver.ExecutorDelegate{DeName: e.Name(), Ctx: nil}
		}
		res[i] = ex
	}
	return res
}
