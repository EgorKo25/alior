package repository

import (
	"callback_service/src/database"
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
)

type IRepository interface {
	CreateCallback(ctx context.Context, data Callback) error
}

type CallbackRepository struct {
	db *pgxpool.Pool
}

type Callback struct {
	ID    int    `db:"id"`
	Name  string `db:"name"`
	Phone string `db:"phone"`
	Type  string `db:"type"`
	Idea  string `db:"idea"`
}

func NewRepository(db *database.Database) IRepository {
	return &CallbackRepository{db: db.Pool}
}

func (r *CallbackRepository) CreateCallback(ctx context.Context, data Callback) error {
	_, err := r.db.Exec(ctx, `
        INSERT INTO callbacks (name, phone, type, idea) 
        VALUES ($1, $2, $3, $4)`,
		data.Name, data.Phone, data.Type, data.Idea)
	return err
}
