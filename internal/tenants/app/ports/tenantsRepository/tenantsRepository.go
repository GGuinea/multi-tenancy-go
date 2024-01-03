package tenantsrepository

import (
	"context"
	"multitenancy/internal/tenants/domain"
)

type TenantsRepository interface {
	List(context.Context) ([]domain.Tenant, error)
}
