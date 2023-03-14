package driver

type Channel interface {
	Looper() EventLooper
	Register(looper EventLooper, future ChannelFuture)
	Write()
	Read()
	Chain() ChannelChain
	Flush()
}
