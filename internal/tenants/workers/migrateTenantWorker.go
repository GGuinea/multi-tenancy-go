package workers

import (
	"context"
	"fmt"
	dbmigrations "multitenancy/internal/pkg/db-migrations"

	"github.com/riverqueue/river"
)

type MigrateTenantArgs struct {
	TenantName string
}

func (MigrateTenantArgs) Kind() string { return "migrate_tenant" }

type MigrateTenanWorker struct {
	river.WorkerDefaults[MigrateTenantArgs]
}

func (w *MigrateTenanWorker) Work(ctx context.Context, job *river.Job[MigrateTenantArgs]) error {
	fmt.Println("MigrateTenanWorker: start", job.Args.TenantName)
	err := dbmigrations.CreateSchemaForTenant(job.Args.TenantName)
	if err != nil {
		return err
	}

	err = dbmigrations.MigrateTenant(job.Args.TenantName)
	if err != nil {
		return err
	}

	fmt.Println("MigrateTenanWorker: end")
	return nil
}
