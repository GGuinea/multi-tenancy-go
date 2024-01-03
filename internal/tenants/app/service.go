package app

import (
	"context"
	tenantsrepository "multitenancy/internal/tenants/app/ports/tenantsRepository"
	"multitenancy/internal/tenants/domain"
)

type TenantService struct {
	tenantsRepository tenantsrepository.TenantsRepository
}

func NewTenantService(tenantsRepository tenantsrepository.TenantsRepository) *TenantService {
	return &TenantService{tenantsRepository: tenantsRepository}
}

func (t *TenantService) ListTenants(ctx context.Context) ([]domain.Tenant, error) {
	tenants, err := t.tenantsRepository.List(ctx)
	if err != nil {
		return nil, err
	}

	return tenants, nil
}
