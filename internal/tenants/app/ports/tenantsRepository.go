package ports

import (
	"context"
	"multitenancy/internal/tenants/domain"
)

type TenantsRepository interface {
	List(context.Context) ([]domain.Tenant, error)
	Save(context.Context, domain.Tenant) error
	ReadById(context.Context, string) (domain.Tenant, error)
}
