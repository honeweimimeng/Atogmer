package chanl

type SelectableChannel interface {
	IOChannel
	Bind0()
}

type DefaultSelectChannel struct {
}
