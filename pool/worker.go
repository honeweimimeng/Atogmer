package pool

type Worker struct {
	pool Pool
	Next func() Task
}

func (w *Worker) doWork(task Task) {
	if task == nil {
		return
	}
	go func() {
		for {
			if task != nil {
				task.Run()
			}
			func() {
				defer func() {
					if r := recover(); r != nil {
						w.pool.ExHandler(task.Ctx())
						return
					}
				}()
			}()
			task = w.Next()
		}
	}()
}
