package workers

import (
	"fmt"

	"github.com/riverqueue/river"
)

type backgroundJobWorkers struct {
	Workers *river.Workers
}

func NewBackgroundJobWorkers(workers *river.Workers) *backgroundJobWorkers {
	return &backgroundJobWorkers{
		Workers: workers,
	}
}

func AddNewWorker[T river.JobArgs](currentWorkers *backgroundJobWorkers, newWorker river.Worker[T]) error {
	err := river.AddWorkerSafely(currentWorkers.Workers, newWorker)
	if err != nil {
		return fmt.Errorf("failed to add new worker: %w", err)
	}
	return nil
}
