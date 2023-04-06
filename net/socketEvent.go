package net

import (
	"litecluster/driver/event"
)

type READ struct {
	id int
}

func (r *READ) Name() string {
	return "READ"
}
func (r *READ) Id() int {
	return r.id
}

type WRITE struct {
	id int
}

func (w *WRITE) Name() string {
	return "WRITE"
}
func (w *WRITE) Id() int {
	return w.id
}

type CLOSE struct {
	id int
}

func (c *CLOSE) Name() string {
	return "CLOSE"
}
func (c *CLOSE) Id() int {
	return c.id
}

func Events(id int) []event.Proto {
	return []event.Proto{
		&READ{id: id},
		&WRITE{id: id},
		&CLOSE{id: id},
	}
}
