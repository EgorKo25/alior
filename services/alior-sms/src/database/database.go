package database

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"

	_ "alior-sms/migrations"
	"alior-sms/src/types"
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

func (db *DB) InsertService(ctx context.Context, service *types.Service) (int32, error) {
	query := `INSERT INTO $1
    (description, price) VALUES ($2, $3, $4) RETURNING id`
	tablePath := `services.public.services` // TODO: Брать путь из переменных окружения
	return service.ID, db.pool.QueryRow(ctx, query, tablePath, service.Name, service.Description, service.Price).Scan(&service.ID)
}

func (d *DB) GetServiceByID(ctx context.Context, id int32) (*types.Service, error) {
	tablePath := `services.public.services`
	query := `SELECT name, description, price FROM $1 WHERE id = $2`
	service := &types.Service{ID: id}

	return service, d.pool.QueryRow(ctx, query, tablePath, id).Scan(&service.Name, &service.Description, &service.Price)
}
