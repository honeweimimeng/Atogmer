package driver

type AcceptableEvent interface {
	Name() string
	Id() int
	Msg() interface{}
}

type EventExecutor struct {
	event AcceptableEvent
	child Executor
	ctx   *EventContext
}

func (e *EventExecutor) Execute() {
	for {
		select {
		case ex := <-e.ctx.Channel(e):
			e.Context().Group().Join(ex)
		default:
			chi := e.child
			if chi != nil {
				chi.Execute()
			}
		}
	}
}

func (e *EventExecutor) Context() ExecutorContext {
	return e.ctx
}

func (e *EventExecutor) Name() string {
	return ""
}
