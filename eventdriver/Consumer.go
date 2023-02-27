package eventdriver

type EventsProto interface {
	Sign(ctx *ProtoCtx)
	Happen(ctx *ProtoCtx)
	CallBack()
	CallBackAsync(func(ctx *ProtoCtx) interface{})
}
