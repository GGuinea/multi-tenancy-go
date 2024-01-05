package tenants

import (
	jobprocessor "multitenancy/internal/pkg/jobProcessor"
	"multitenancy/internal/tenants/app"
	"multitenancy/internal/tenants/app/adapters"

	"github.com/jackc/pgx/v5/pgxpool"
)

type TenantDependencies struct {
	DbPool       *pgxpool.Pool
	JobProcessor *jobprocessor.JobProcessorClient
}

type DependencyTree struct {
	TenantService *app.TenantService
}

func NewTenantDependencies(deps *TenantDependencies) *DependencyTree {
	if deps == nil {
		panic("deps are nil")
	}

	tenantRepository := adapters.NewTenantsPostgresRepository(deps.DbPool)
	tenantService := app.NewTenantService(tenantRepository, deps.JobProcessor)

	return &DependencyTree{
		TenantService: tenantService,
	}
}
