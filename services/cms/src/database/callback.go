package database

import (
	"context"
)

type ICallback interface {
	CreateCallback(ctx context.Context, data Callback) error
}

type Callback struct {
	ID    int    `db:"id"`
	Name  string `db:"name"`
	Phone string `db:"phone"`
	Type  string `db:"type"`
	Idea  string `db:"idea"`
}

func (d *Database) CreateCallback(ctx context.Context, data Callback) error {
	_, err := d.Pool.Exec(ctx, `
        INSERT INTO callbacks (name, phone, type, idea) 
        VALUES ($1, $2, $3, $4)`,
		data.Name, data.Phone, data.Type, data.Idea)
	return err
}
