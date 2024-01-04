package tenants

import (
	"multitenancy/internal/tenants/app"
	"multitenancy/internal/tenants/app/adapters"

	"github.com/jackc/pgx/v5/pgxpool"
)

type TenantDependencies struct {
	DbPool *pgxpool.Pool
}

type DependencyTree struct {
	TenantService *app.TenantService
}

func NewTenantDependencies(dbPool *pgxpool.Pool) *DependencyTree {
	if dbPool == nil {
		panic("dbPool is nil")
	}

	tenantRepository := adapters.NewTenantsPostgresRepository(dbPool)
	tenantService := app.NewTenantService(tenantRepository)

	return &DependencyTree{
		TenantService: tenantService,
	}
}
