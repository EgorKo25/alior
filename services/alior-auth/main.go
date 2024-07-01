package main

import (
	"context"
	"log"

	"alior-auth/src/database"

	_ "alior-auth/migrations"

	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {
	ctx := context.Background()
	pool, err := pgxpool.New(ctx, "postgres://auth:aliorAuth@authdb:5432/auth?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	err = database.MigrateDatabase(pool)
	if err != nil {
		log.Fatal(err)
	}
}
