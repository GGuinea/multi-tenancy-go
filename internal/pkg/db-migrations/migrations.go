package dbmigrations

import (
	"embed"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jackc/pgx/v5/stdlib"
	"github.com/pressly/goose/v3"
)

//go:embed migrations/global/*.sql
var embedMigrations embed.FS

func Migrate(pool *pgxpool.Pool) error {
	goose.SetDialect("postgres")
	goose.SetBaseFS(embedMigrations)

	db := stdlib.OpenDBFromPool(pool)

	if err := goose.Up(db, "migrations/global"); err != nil {
		return err
	}

	if err := db.Close(); err != nil {
		return err
	}

	return nil
}
