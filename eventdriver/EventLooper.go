package eventdriver

import "litecluster/call/chanl"

type EventLooper struct {
	selectChannel chanl.SelectableChannel
}

func (loop *EventLooper) Select(channel chanl.SelectableChannel) {
	loop.selectChannel = channel
	loop.bind()
}

func (loop *EventLooper) bind() {
	loop.selectChannel.Bind0()
}
