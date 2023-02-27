package eventdriver

type EventHandle interface {
	Handle(ctx *ProtoCtx) interface{}
}

type EventsProto interface {
	EventHandle
	Sign(ctx *ProtoCtx)
	Process(ctx *ProtoCtx)
	CallBack()
	CallBackAsync(func(ctx *ProtoCtx) interface{})
}
