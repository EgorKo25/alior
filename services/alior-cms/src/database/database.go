package database

import (
	_ "callback_service/migrations"
	"callback_service/src/config"
	"context"
	"database/sql"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jackc/pgx/v5/stdlib"
	"github.com/pressly/goose/v3"
)

type IMigrator interface {
	Up(db *sql.DB, dir string) error
}

type Migrator struct{}

func (m *Migrator) Up(db *sql.DB, dir string) error {
	return goose.Up(db, dir)
}

type Database struct {
	Pool *pgxpool.Pool
}

func New(ctx context.Context, cfg *config.Config) (*Database, error) {
	pool, err := pgxpool.New(ctx, cfg.Database.Url)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	err = runMigrations(pool, &Migrator{})
	if err != nil {
		return nil, fmt.Errorf("migrations failed: %w", err)
	}

	return &Database{Pool: pool}, nil
}

func runMigrations(pool *pgxpool.Pool, migrator IMigrator) error {
	db := stdlib.OpenDBFromPool(pool)
	return migrator.Up(db, "migrations")
}

func (d *Database) Close() {
	d.Pool.Close()
}
