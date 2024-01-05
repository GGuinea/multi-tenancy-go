package app

import (
	"context"
	"multitenancy/internal/tenants/app/ports"
	"multitenancy/internal/tenants/domain"
	"multitenancy/internal/tenants/workers"

	"github.com/google/uuid"
)

type TenantService struct {
	tenantsRepository ports.TenantsRepository
	jobProcessor      ports.JobProcessor
}

func NewTenantService(tenantsRepository ports.TenantsRepository, jobProcessor ports.JobProcessor) *TenantService {
	return &TenantService{tenantsRepository: tenantsRepository, jobProcessor: jobProcessor}
}

func (t *TenantService) ListTenants(ctx context.Context) ([]domain.Tenant, error) {
	tenants, err := t.tenantsRepository.List(ctx)
	if err != nil {
		return nil, err
	}

	return tenants, nil
}

func (t *TenantService) AddTenant(ctx context.Context, tenant domain.TenantRequestDto) (string, error) {
	uuid := uuid.New().String()

	err := t.tenantsRepository.Save(ctx, domain.Tenant{
		ID:     uuid,
		Name:   tenant.Name,
		Active: tenant.Active,
	})

	if err != nil {
		return "", err
	}

	// this code should be executed asynchronously
	// err = dbmigrations.CreateSchemaForTenant(tenant.Name)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	//
	// err = dbmigrations.MigrateTenant(tenant.Name)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	t.jobProcessor.ScheduleNewJob(ctx, workers.MigrateTenantArgs{
		TenantName: tenant.Name,
	})
	// this code should be executed asynchronously

	return uuid, nil
}

func (t *TenantService) ReadById(ctx context.Context, id string) (domain.Tenant, error) {
	tenant, err := t.tenantsRepository.ReadById(ctx, id)
	if err != nil {
		return domain.Tenant{}, err
	}

	return tenant, nil
}
