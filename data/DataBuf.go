package data

import "io"

type BufProto interface {
	Writer() io.Writer
	Reader() io.Reader
	Size() int64
}

type Formatter interface {
}
