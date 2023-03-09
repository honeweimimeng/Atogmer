package data

import "io"

type IOProto interface {
	Writer0() io.Writer
	Reader0() io.Reader
	Size() int64
}
