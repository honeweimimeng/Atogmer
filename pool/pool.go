package pool

import (
	"context"
	"sync/atomic"
)

type Pool interface {
	StartUp()
	Name() string
	Run(task Task)
	WorkerCount() uint32
	Cap() uint32
	ExHandler(ctx context.Context)
}

type DefaultPool struct {
	name          string
	workCount     uint32
	cap           uint32
	ctx           context.Context
	acceptChannel chan Task
	pipe          TaskPipe
}

func (p *DefaultPool) StartUp() {
	p.acceptChannel = make(chan Task, p.cap)
	go func() {
		for {
			select {
			case t := <-p.acceptChannel:
				p.pipe.PushTask(t)
			case <-p.ctx.Done():
				println("done msg:", p.ctx.Err().Error())
				return
			default:
				atomic.AddUint32(&p.workCount, 1)
				if atomic.LoadUint32(&p.workCount) <= atomic.LoadUint32(&p.cap) {
					w := &Worker{Next: func() Task {
						return p.pipe.PopTask()
					}}
					w.doWork(p.pipe.PopTask())
				}
			}
		}
	}()
}

func (p *DefaultPool) Name() string {
	return p.name
}

func (p *DefaultPool) Run(task Task) {
	p.acceptChannel <- task
}

func (p *DefaultPool) WorkerCount() uint32 {
	return p.workCount
}

func (p *DefaultPool) Cap() uint32 {
	return p.cap
}

func (p *DefaultPool) ExHandler(ctx context.Context) {

}
