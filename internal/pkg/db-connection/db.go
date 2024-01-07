package dbconnection

import (
	"context"
	"fmt"
	"multitenancy/config"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

func GetDbPool(ctx context.Context, config *config.DbConfig) (*pgxpool.Pool, error) {
	return pgxpool.New(ctx, buildConnectionString(config, "public"))
}

func GetDbConnectionForTenant(ctx context.Context, config *config.DbConfig, tenant string) (*pgx.Conn, error) {
	return pgx.Connect(ctx, buildConnectionString(config, tenant))
}

func buildConnectionString(dbConfig *config.DbConfig, schema string) string {
	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable&search_path=%s",
		dbConfig.User, dbConfig.Pass, dbConfig.Host, dbConfig.Port, dbConfig.DbName, schema)
}
