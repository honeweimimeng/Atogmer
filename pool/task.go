package pool

import (
	"container/heap"
	"context"
)

type Task interface {
	Run()
	Ctx() context.Context
}

type TaskPipe interface {
	PushTask(task Task)
	PopTask() Task
}

type Comparable interface {
	Val() interface{}
	CompareTo(other interface{}) int
}

func GetComparablePipe() *ComparablePipe {
	arr := make([]Comparable, 0)
	r := &ComparablePipe{heap: arr}
	return r
}

type ComparablePipe struct {
	heap []Comparable
}

func (t *ComparablePipe) PushTask(task Task) {
	t.Push(task)
}

func (t *ComparablePipe) PopTask() Task {
	ro := t.Pop()
	if ro == nil {
		return nil
	}
	return ro.(Task)
}

func (t *ComparablePipe) Push(task any) {
	c, ok := task.(Comparable)
	if !ok {
		panic("task must type of Comparable")
	}
	t.heap = append(t.heap, c)
	heap.Init(t)
}

func (t *ComparablePipe) Pop() any {
	var v Comparable = nil
	var r Task = nil
	if t.Len() > 0 {
		t.heap, v = t.heap[:t.Len()-1], t.heap[t.Len()-1]
	}
	r, ok := v.(Task)
	if !ok && v != nil {
		panic("node must type of Task")
	}
	return r
}

func (t *ComparablePipe) Less(i, j int) bool {
	return t.heap[i].CompareTo(t.heap[j].Val()) < 0
}

func (t *ComparablePipe) Swap(i, j int) {
	t.heap[i], t.heap[j] = t.heap[j], t.heap[i]
}

func (t *ComparablePipe) Len() int {
	return len(t.heap)
}
