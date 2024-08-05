package database

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"

	_ "alior-sms/migrations"
)

type DB struct {
	pool *pgxpool.Pool
}

func NewDB(ctx context.Context, connString string, migrationDir string) (*DB, error) {
	pool, err := pgxpool.New(ctx, connString)
	if err != nil {
		return nil, err
	}

	if err := MigrateDatabase(pool, migrationDir); err != nil {
		pool.Close()
		return nil, err
	}

	return &DB{pool: pool}, nil
}
