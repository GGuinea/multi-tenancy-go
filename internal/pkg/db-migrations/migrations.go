package dbmigrations

import (
	"context"
	"embed"
	"multitenancy/config"
	dbconnection "multitenancy/internal/pkg/db-connection"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jackc/pgx/v5/stdlib"
	"github.com/pressly/goose/v3"
)

//go:embed migrations/global/*.sql
var globalMigrations embed.FS

//go:embed migrations/tenant/*.sql
var tenantMigrations embed.FS

func MigrateGlobal(pool *pgxpool.Pool) error {
	goose.SetDialect("postgres")
	goose.SetBaseFS(globalMigrations)

	db := stdlib.OpenDBFromPool(pool)

	defer db.Close()

	if err := goose.Up(db, "migrations/global"); err != nil {
		return err
	}

	return nil
}

func MigrateTenant(tenantName string) error {
	ctx := context.Background()
	goose.SetDialect("postgres")
	goose.SetBaseFS(tenantMigrations)

	dbconn, err := dbconnection.GetDbConnectionForTenant(ctx, &config.NewConfig().Db, tenantName)

	if err != nil {
		return err
	}

	db := stdlib.OpenDB(*dbconn.Config())
	defer db.Close()

	if err := goose.Up(db, "migrations/tenant"); err != nil {
		return err
	}

	dbconn.Exec(ctx, "set search_path to public")
	return nil
}

func CreateSchemaForTenant(tenantName string) error {
	ctx := context.Background()
	dbconn, err := dbconnection.GetDbConnectionForTenant(ctx, &config.NewConfig().Db, tenantName)

	if err != nil {
		return err
	}

	_, err = dbconn.Exec(ctx, "CREATE SCHEMA IF NOT EXISTS "+tenantName)

	if err != nil {
		return err
	}

	dbconn.Exec(ctx, "set search_path to public")
	return nil
}
