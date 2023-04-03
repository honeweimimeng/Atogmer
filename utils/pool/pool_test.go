package pool

import (
	"container/heap"
	"context"
	"testing"
	"time"
)

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

type DefaultTask struct {
	ctx   context.Context
	count int
}

func (t *DefaultTask) Run() {
	if t.count == 1 {
		time.Sleep(time.Duration(10) * time.Second)
	} else {
		time.Sleep(time.Duration(3) * time.Second)
	}
	println("===>Run：", t.count)
}
func (t *DefaultTask) Ctx() context.Context {
	return t.ctx
}

func (t *DefaultTask) Val() interface{} {
	return t.count
}
func (t *DefaultTask) CompareTo(other interface{}) int {
	i := other.(int) - t.count
	return i
}

func TestPool(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	pipe := &ComparablePipe{}
	pool := &DefaultPool{Cap_: 2, Name_: "d1", Ctx: ctx, Pipe: pipe}
	pool.StartUp()
	tCtx, tCancel := context.WithCancel(context.Background())
	defer tCancel()
	pool.Run(&DefaultTask{ctx: tCtx, count: 1})
	pool.Run(&DefaultTask{ctx: tCtx, count: 2})
	pool.Run(&DefaultTask{ctx: tCtx, count: 3})
	time.Sleep(1 * time.Minute)
}

func TestComparable(t *testing.T) {
	c := GetComparablePipe()
	n := &DefaultTask{ctx: context.Background(), count: 0}
	n1 := &DefaultTask{ctx: context.Background(), count: 5}
	n2 := &DefaultTask{ctx: context.Background(), count: 2}
	n3 := &DefaultTask{ctx: context.Background(), count: 3}
	c.PushTask(n)
	c.PushTask(n1)
	c.PushTask(n2)
	c.PushTask(n3)
	c.PopTask().Run()
	c.PopTask().Run()
	c.PopTask().Run()
	c.PopTask().Run()
}

func TestFifo(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	pool := &DefaultPool{Cap_: 10, Name_: "d1", Ctx: ctx, Pipe: GetFifoPipe(10)}
	pool.StartUp()
	pool.Run(&DefaultTask{ctx: context.Background(), count: 1})
	pool.Run(&DefaultTask{ctx: context.Background(), count: 2})
	pool.Run(&DefaultTask{ctx: context.Background(), count: 3})
	pool.Run(&DefaultTask{ctx: context.Background(), count: 4})
	time.Sleep(1 * time.Minute)
}
