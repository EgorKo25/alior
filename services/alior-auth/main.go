package main

import (
	"context"
	"log"

	"alior-digital/src/database"

	_ "alior-digital/migrations"

	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {
	ctx := context.Background()
	pool, err := pgxpool.New(ctx, "postgres://auth:aliorAuth@localhost:5432/auth?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	err = database.MigrateDatabase(pool)
	if err != nil {
		log.Fatal(err)
	}
}
