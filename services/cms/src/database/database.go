package database

import (
	"callback_service/src/config"
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jackc/pgx/v5/stdlib"
	"github.com/pressly/goose/v3"

	_ "callback_service/migrations"
)

type Database struct {
	Pool *pgxpool.Pool
}

func New(ctx context.Context, cfg *config.Config) (*Database, error) {
	pool, err := pgxpool.New(ctx, cfg.Database.Url)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	err = runMigrations(pool)
	if err != nil {
		return nil, fmt.Errorf("migrations failed: %w", err)
	}

	return &Database{Pool: pool}, nil
}

func runMigrations(pool *pgxpool.Pool) error {
	db := stdlib.OpenDBFromPool(pool)
	return goose.Up(db, "migrations")
}

func (db *Database) Close() {
	db.Pool.Close()
}
