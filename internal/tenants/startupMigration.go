package tenants

import "context"

func VerifyTenantsMigrations(ctx context.Context, deps *DependencyTree) error {
	service := deps.TenantService
	allTenants, err := service.ListTenants(ctx)
	if err != nil {
		return err
	}

	for _, tenant := range allTenants {
		err = service.MigrateTenant(ctx, tenant.Name)
		if err != nil {
			return err
		}
	}

	return nil
}
