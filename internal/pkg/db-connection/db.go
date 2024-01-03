package dbconnection

import (
	"multitenancy/config"
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Public func(tenant string) (*pgx.Conn, error)

func GetDbPool(ctx context.Context, config *config.DbConfig) (*pgxpool.Pool, error) {
	return pgxpool.New(ctx, buildConnectionString(config, "public"))
}

func buildConnectionString(dbConfig *config.DbConfig, schema string) string {
	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable&search_path=%s",
		dbConfig.User, dbConfig.Pass, dbConfig.Host, dbConfig.Port, dbConfig.DbName, schema)
}
