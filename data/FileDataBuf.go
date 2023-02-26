package data

import (
	"io"
	"os"
)

type FileDataBuf struct {
	Path string
	file *os.File
}

func (n *FileDataBuf) Remove() {
	if n.FileExits() {
		err := os.Remove(n.Path)
		if err != nil {
			panic(err)
		}
	}
}

func (n *FileDataBuf) Writer() io.Writer {
	file := n.FileExitAndCreate()
	return file
}

func (n *FileDataBuf) Reader() io.Reader {
	file, err := os.OpenFile(n.Path, os.O_RDONLY, 0666)
	if err != nil {
		return nil
	}
	return file
}

func (n *FileDataBuf) Size() int64 {
	file := n.FileExitAndCreate()
	fi, _ := file.Stat()
	return fi.Size()
}

func (n *FileDataBuf) FileExitAndCreate() *os.File {
	if n.file != nil {
		return n.file
	}
	exits := n.FileExits()
	if exits {
		fi, err := os.OpenFile(n.Path, os.O_RDONLY|os.O_CREATE|os.O_APPEND, 0666)
		if err != nil {
			panic(err)
		}
		n.file = fi
		return fi
	}
	if !exits {
		fi, err := os.Create(n.Path)
		if err != nil {
			panic(err)
		}
		n.file = fi
		return fi
	}
	return nil
}

func (n *FileDataBuf) FileExits() bool {
	_, err := os.Stat(n.Path)
	return !os.IsNotExist(err)
}
