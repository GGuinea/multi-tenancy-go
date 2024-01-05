package main

import (
	"context"
	"fmt"
	"multitenancy/config"
	"multitenancy/internal"
	"multitenancy/internal/backgroundJobs/workers"
	"multitenancy/internal/drivers/rest"
	dbmigrations "multitenancy/internal/pkg/db-migrations"
	jobprocessor "multitenancy/internal/pkg/jobProcessor"
	"multitenancy/internal/pkg/jobProcessor/migrations"
	"multitenancy/internal/tenants"

	"github.com/gin-gonic/gin"
)

func main() {
	ctx := context.Background()

	cfg := config.NewConfig()

	compositionRoot := internal.NewCompositionRoot(ctx, cfg)
	backgroundJob, err := setupBackgroundJobProcessor(ctx, cfg, compositionRoot)
	fmt.Println(backgroundJob)
	if err != nil {
		panic(err)
	}

	err = dbmigrations.MigrateGlobal(compositionRoot.DbPool)
	if err != nil {
		panic(err)
	}

	router := gin.Default()
	tenantDependencies := tenants.NewTenantDependencies(compositionRoot.DbPool)
	rest.BuildRoutes(router, tenantDependencies)
	router.Run(":8080")
}

func setupBackgroundJobProcessor(ctx context.Context, cfg *config.Config, deps *internal.CompositionRoot) (*jobprocessor.JobProcessorClient, error) {
	err := migrations.PerformStartupRiverMigration(ctx, deps.DbPool)
	if err != nil {
		return nil, err
	}

	backgroundWorkersMgmnt := workers.NewBackgroundJobWorkers(deps.BackgroundWorkers)
	workers.AddDefaultWorker(backgroundWorkersMgmnt)
	workers.AddNewWorker(backgroundWorkersMgmnt, &workers.NewRequestWorker{})

	jobProcessorClient, err := jobprocessor.NewJobProcessorClient(ctx, deps, cfg.BackgroundProcessorConfig, deps.DbPool)
	if err != nil {
		return nil, err
	}

	return jobProcessorClient, nil
}
