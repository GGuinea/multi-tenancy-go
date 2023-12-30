package main

import (
	"async_worker/config"
	"async_worker/internal"
	"async_worker/internal/backgroundJobs/workers"
	"async_worker/internal/drivers/rest"
	dbconnection "async_worker/internal/pkg/db-connection"
	jobprocessor "async_worker/internal/pkg/jobProcessor"
	"async_worker/internal/pkg/jobProcessor/migrations"
	someresource "async_worker/internal/someResource"
	"context"

	"github.com/gin-gonic/gin"
)

func main() {
	ctx := context.Background()

	cfg := config.NewConfig()
	compositionRoot := internal.NewCompositionRoot(ctx, cfg)
	backgroundJob, err := setupBackgroundJobProcessor(ctx, cfg, compositionRoot)
	if err != nil {
		panic(err)
	}

	defer backgroundJob.DbPool.Close()

	router := gin.Default()
	someResourceDependencies := someresource.NewSomeResourceDependencies(backgroundJob)
	rest.BuildRoutes(router, someResourceDependencies)
	router.Run(":8080")
}

func setupBackgroundJobProcessor(ctx context.Context, cfg *config.Config, deps *internal.CompositionRoot) (*jobprocessor.JobProcessorClient, error) {
	dbPool, err := dbconnection.GetDbPool(ctx, &cfg.Db)

	if err != nil {
		return nil, err
	}

	err = migrations.PerformStartupRiverMigration(ctx, dbPool)
	if err != nil {
		return nil, err
	}

	backgroundWorkersMgmnt := workers.NewBackgroundJobWorkers(deps.BackgroundWorkers)
	workers.AddDefaultWorker(backgroundWorkersMgmnt)
	workers.AddNewWorker(backgroundWorkersMgmnt, &workers.NewRequestWorker{})

	jobProcessorClient, err := jobprocessor.NewJobProcessorClient(ctx, deps, cfg.BackgroundProcessorConfig, dbPool)
	if err != nil {
		return nil, err
	}

	return jobProcessorClient, nil
}
