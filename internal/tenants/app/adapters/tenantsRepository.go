package adapters

import (
	"context"
	"multitenancy/internal/tenants/domain"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

type TenantsPostgresRepository struct {
	DbPool *pgxpool.Pool
}

func NewTenantsPostgresRepository(dbPool *pgxpool.Pool) *TenantsPostgresRepository {
	return &TenantsPostgresRepository{DbPool: dbPool}
}

func (tpr *TenantsPostgresRepository) List(ctx context.Context) ([]domain.Tenant, error) {
	tx, err := tpr.DbPool.Begin(ctx)

	if err != nil {
		return nil, err
	}
	defer tx.Rollback(ctx)

	rows, err := tx.Query(ctx, "SELECT id, name, active FROM tenant")

	if err != nil {
		return nil, err
	}

	tenants := make([]domain.Tenant, 0)

	for rows.Next() {
		var tenant domain.Tenant
		err := rows.Scan(&tenant.ID, &tenant.Name, &tenant.Active)

		if err != nil {
			return nil, err
		}

		tenants = append(tenants, tenant)
	}

	err = tx.Commit(ctx)

	if err != nil {
		return tenants, err
	}

	return tenants, nil
}

func (tpr *TenantsPostgresRepository) Save(ctx context.Context, tenant domain.Tenant) error {
	tx, err := tpr.DbPool.Begin(ctx)

	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	_, err = tx.Exec(ctx, "INSERT INTO tenant VALUES ($1, $2, $3, $4)", tenant.ID, tenant.Name, time.Now(), tenant.Active)

	if err != nil {
		return err
	}

	err = tx.Commit(ctx)
	if err != nil {
		return err
	}
	return nil
}

func (tpr *TenantsPostgresRepository) ReadById(ctx context.Context, uuid string) (domain.Tenant, error) {
	tx, err := tpr.DbPool.Begin(ctx)
	if err != nil {
		return domain.Tenant{}, err
	}

	defer tx.Rollback(ctx)

	var tenant domain.Tenant
	err = tx.QueryRow(ctx, "SELECT id, name, active FROM tenant WHERE id = $1", uuid).Scan(&tenant.ID, &tenant.Name, &tenant.Active)

	if err != nil {
		return domain.Tenant{}, err
	}

	return tenant, nil
}
