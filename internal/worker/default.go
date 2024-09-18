package worker

import (
	"context"
	"log/slog"
)

var defaultPool *Pool

func SetupDefaultPool(lg *slog.Logger, workerCount int) {
	defaultPool = NewPool(lg, workerCount)
	defaultPool.Start(context.Background())
}

func StopDefaultPool() {
	if defaultPool != nil {
		defaultPool.Stop()
	}
}

func AddTask(task Task) {
	defaultPool.AddTask(task)
}
