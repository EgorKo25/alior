package database

import (
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jackc/pgx/v5/stdlib"
	"github.com/pressly/goose/v3"
)

func MigrateDatabase(dbConnection *pgxpool.Pool) error {
	return goose.Up(stdlib.OpenDBFromPool(dbConnection), ".")
}
