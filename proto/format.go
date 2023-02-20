package proto

type Format interface {
	Wrapper(wrapper AbstractWrapper)
	Decoder() *BuffDecoder
	Encoder() *BuffEncoder
}

type DefaultFormat struct {
}

func (f *DefaultFormat) Wrapper(wrapper AbstractWrapper) {

}

func (f *DefaultFormat) Decoder() *BuffDecoder {

}

func (f *DefaultFormat) Encoder() *BuffEncoder {

}