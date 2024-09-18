package worker

import (
	"context"
	"log/slog"
)

type Pool struct {
	lg          *slog.Logger
	tasks       chan Task
	workerCount int
	cancelFunc  context.CancelFunc
}

func NewPool(lg *slog.Logger, workerCount int) *Pool {
	return &Pool{
		lg:          lg,
		tasks:       make(chan Task, workerCount * 5),
		workerCount: workerCount,
	}
}

func (p *Pool) Start(ctx context.Context) {
	ctx, cancel := context.WithCancel(ctx)
	p.cancelFunc = cancel

	for i := 0; i < p.workerCount; i++ {
		go p.worker(ctx)
	}
}

func (p *Pool) Stop() {
	if p.cancelFunc != nil {
		p.cancelFunc()
	}
}

func (p *Pool) AddTask(task Task) {
	p.tasks <- task
}

func (p *Pool) worker(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			return
		case task := <-p.tasks:
			err := task.Execute()
			if err != nil {
				p.lg.Error(err.Error())
			}
		}
	}
}
