package repository

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Repository interface {
	CreateCallback(ctx context.Context, data Callback) error
}

type repo struct {
	db *pgxpool.Pool
}

type Callback struct {
	ID     int       `db:"id"`
	Number string    `db:"number"`
	Date   time.Time `db:"date"`
	Name   string    `db:"name"`
}

func NewRepository(db *pgxpool.Pool) Repository {
	return &repo{db: db}
}

func (r *repo) CreateCallback(ctx context.Context, data Callback) error {
	_, err := r.db.Exec(ctx, `
        INSERT INTO callbacks (number, date, name) 
        VALUES ($1, $2, $3)`,
		data.Number, data.Date, data.Name)
	return err
}
