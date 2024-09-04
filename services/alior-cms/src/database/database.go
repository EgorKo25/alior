package database

import (
	_ "callback_service/migrations" // blank import to run migrations
	"callback_service/src/config"
	"context"
	"database/sql"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jackc/pgx/v5/stdlib"
	"github.com/pressly/goose/v3"
)

// IMigrator declare method to run migrations
type IMigrator interface {
	Up(db *sql.DB, dir string) error
}

// Migrator is an empty structure to implement IMigrator method
type Migrator struct{}

// Up implements IMigrator method
func (m *Migrator) Up(db *sql.DB, dir string) error {
	return goose.Up(db, dir)
}

// Database structure to store db pool
type Database struct {
	Pool *pgxpool.Pool
}

// New is a function to create new db pool and run migrations
func New(ctx context.Context, cfg *config.Config) (*Database, error) {
	pool, err := pgxpool.New(ctx, cfg.Database.URL)
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

// Close is a Database structure method to close connections pool
func (d *Database) Close() {
	d.Pool.Close()
}
