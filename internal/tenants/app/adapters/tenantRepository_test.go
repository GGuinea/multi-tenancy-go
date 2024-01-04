package adapters

import (
	"context"
	"multitenancy/config"
	dbconnection "multitenancy/internal/pkg/db-connection"
	"multitenancy/internal/tenants/domain"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func TestShouldSaveTenant(t *testing.T) {
	// Given
	ctx := context.Background()
	config := config.NewConfig()
	dbPool, err := dbconnection.GetDbPool(ctx, &config.Db)
	if err != nil {
		t.Fatal(err)
	}
	tenantsRepository := NewTenantsPostgresRepository(dbPool)
	tenant := domain.Tenant{
		Name:   "test-tenant",
		Active: true,
	}
	// When
	err = tenantsRepository.Save(ctx, tenant)

	// Then
	if err != nil {
		t.Fatal(err)
	}
}

func TestShouldReadTenantByUUID(t *testing.T) {
	ctx := context.Background()
	config := config.NewConfig()
	dbPool, err := dbconnection.GetDbPool(ctx, &config.Db)
	if err != nil {
		t.Fatal(err)
	}

	tenantsRepository := NewTenantsPostgresRepository(dbPool)
	tenant := domain.Tenant{
		ID:     uuid.New().String(),
		Name:   "test-tenant",
		Active: true,
	}

	err = tenantsRepository.Save(ctx, tenant)

	if err != nil {
		t.Fatal(err)
	}

	tenantRead, err := tenantsRepository.ReadById(ctx, tenant.ID)

	if err != nil {
		t.Fatal(err)
	}

	require.Equal(t, tenant, tenantRead)
}
