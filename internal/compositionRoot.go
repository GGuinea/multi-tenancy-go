package internal

import (
	"async_worker/config"
	dbconnection "async_worker/internal/pkg/db-connection"
	"context"

	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/riverqueue/river"
)

type CompositionRoot struct {
	BackgroundWorkers *river.Workers
	DbPool            *pgxpool.Pool
}

func NewCompositionRoot(ctx context.Context, config *config.Config) *CompositionRoot {
	dbPool, err := dbconnection.GetDbPool(ctx, &config.Db)
	if err != nil {
		panic(err)
	}

	backgroundWorkers, err := initBackgroundJobWorkers()
	if err != nil {
		panic(err)
	}

	return &CompositionRoot{DbPool: dbPool, BackgroundWorkers: backgroundWorkers}
}

func initBackgroundJobWorkers() (*river.Workers, error) {
	workers := river.NewWorkers()
	if workers == nil {
		return nil, fmt.Errorf("failed to create workers")
	}

	return workers, nil
}
