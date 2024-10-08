package database

import (
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jackc/pgx/v5/stdlib"
	"github.com/pressly/goose/v3"
)

func MigrateDatabase(dbConnection *pgxpool.Pool, migrationDir string) error {
	DB := stdlib.OpenDBFromPool(dbConnection)
	defer DB.Close()

	return goose.Up(DB, migrationDir)
}
