package driver

type LooperGroup interface {
	Executor
	Registry(channel Channel)
	Next() LooperGroup
}
