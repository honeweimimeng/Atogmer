package pool

import (
	"context"
	"testing"
	"time"
)

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
	pool := &DefaultPool{cap: 2, name: "d1", ctx: ctx, pipe: pipe}
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
