package driver

import "litecluster/utils"

type EventExecutorBus struct {
	eventTrigger   EventTrigger
	executorAdvice chan []Executor
	eventMap       *utils.SafeMap[AcceptableEvent, Executor]
}

func UseExecutorBus(trigger EventTrigger) *EventExecutorBus {
	return &EventExecutorBus{
		eventMap:       utils.UseInterSafeMap[AcceptableEvent, Executor](),
		eventTrigger:   trigger,
		executorAdvice: make(chan []Executor),
	}
}

func (bus *EventExecutorBus) Join(executor Executor) {
	eventExe := executor.(*EventExecutor)
	bus.eventMap.Put(eventExe.event, executor)
}

func (bus *EventExecutorBus) AcceptExecutors() []Executor {
	events := bus.eventTrigger.AcceptEvent()
	res := make([]Executor, len(events))
	for i, e := range events {
		ex := bus.eventMap.Get(e)
		if ex == nil {
			ex = &ExecutorDelegate{DeName: e.Name(), Ctx: nil}
		}
		res[i] = ex
	}
	return res
}
