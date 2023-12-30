package internal

import (
	"async_worker/config"
	"context"

	"fmt"

	"github.com/riverqueue/river"
)

type CompositionRoot struct {
	BackgroundWorkers *river.Workers
}

func NewCompositionRoot(ctx context.Context, config *config.Config) *CompositionRoot {
	backgroundWorkers, err := initBackgroundJobWorkers()
	if err != nil {
		panic(err)
	}

	return &CompositionRoot{BackgroundWorkers: backgroundWorkers}
}

func initBackgroundJobWorkers() (*river.Workers, error) {
	workers := river.NewWorkers()
	if workers == nil {
		return nil, fmt.Errorf("failed to create workers")
	}

	return workers, nil
}
