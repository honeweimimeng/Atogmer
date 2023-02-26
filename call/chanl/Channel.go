package chanl

import (
	"bufio"
	"io"
	"litecluster/data"
	"log"
)

type IOChannel interface {
	Write(b []byte) int
	ReadBytes(size int) []byte
	ReadString() string
	Init()
}

type BufChannel struct {
	Adapter     data.BufProto
	defaultSize int
}

func (c *BufChannel) Init() {
	defSie := 4096
	if c.defaultSize == 0 || c.defaultSize < defSie {
		c.defaultSize = defSie
	}
}

func (c *BufChannel) Write(b []byte) int {
	w := bufio.NewWriterSize(c.Adapter.Writer(), len(b))
	size := 0
	for i := len(b); i > 0; i = i - c.defaultSize {
		si, ei := c.prefix(int64(i), int64(c.defaultSize), b)
		s, err := w.Write(b[si:ei])
		size = size + s
		if err == io.EOF {
			return 0
		}
		err = w.Flush()
		if err != nil {
			log.Println("flush error :", err)
		}
	}
	return size
}

func (c *BufChannel) ReadBytes(size int) []byte {
	onOfRead := int64(c.defaultSize)
	var bytes []byte
	if size == -1 {
		bytes = make([]byte, c.Adapter.Size())
	} else {
		bytes = make([]byte, size)
	}
	r := bufio.NewReader(c.Adapter.Reader())
	var lens int
	for i := c.Adapter.Size(); i > 0; i = i - int64(lens) {
		si, ei := c.prefix(i, onOfRead, bytes)
		_, lens = c.DoReadBytes(r, bytes[si:ei])
		if lens == -1 {
			break
		}
	}
	return bytes
}

func (c *BufChannel) prefix(i, onOfRead int64, bytes []byte) (int64, int64) {
	sIdx := int64(len(bytes)) - i
	eIdx := sIdx + onOfRead
	if eIdx > int64(len(bytes)) {
		eIdx = int64(len(bytes))
	}
	return sIdx, eIdx
}

func (c *BufChannel) DoReadBytes(r *bufio.Reader, bytes []byte) ([]byte, int) {
	size, err := r.Read(bytes)
	if err == io.EOF {
		return make([]byte, 0), -1
	}
	return bytes, size
}

func (c *BufChannel) ReadString() string {
	res := c.ReadBytes(-1)
	return string(res)
}
