package tenantspostgresrepository

import (
	"context"
	"multitenancy/internal/tenants/domain"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

type TenantsPostgresRepository struct {
	DbPool *pgxpool.Pool
}

type TenantDto struct {
	Id        string    `db:"id"`
	Name      string    `db:"name"`
	Active    bool      `db:"active"`
	CreatedAt time.Time `db:"created_at"`
}

func NewTenantsPostgresRepository(dbPool *pgxpool.Pool) *TenantsPostgresRepository {
	return &TenantsPostgresRepository{DbPool: dbPool}
}

func (tpr *TenantsPostgresRepository) List(ctx context.Context) ([]domain.Tenant, error) {
	tx, err := tpr.DbPool.Begin(ctx)
	defer tx.Rollback(ctx)

	if err != nil {
		return nil, err
	}

	rows, err := tx.Query(ctx, "SELECT id, name, active FROM tenant")

	if err != nil {
		return nil, err
	}

	tenants := make([]domain.Tenant, 0)

	for rows.Next() {
		var tenant TenantDto
		err := rows.Scan(&tenant.Id, &tenant.Name, &tenant.Active)

		if err != nil {
			return nil, err
		}

		tenants = append(tenants, domain.Tenant{
			ID:     tenant.Id,
			Name:   tenant.Name,
			Active: tenant.Active,
		})
	}

	err = tx.Commit(ctx)

	if err != nil {
		return tenants, err
	}

	return tenants, nil
}
