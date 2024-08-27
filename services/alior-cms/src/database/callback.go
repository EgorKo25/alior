package database

import (
	"context"
	"errors"
	"time"
)

// Limit constant variable represents 1
const Limit = 1

// Offset variable represent global db offset to manage initial, next and previous requests from another services
var Offset = 0

// ICallback declare interaction methods for Callback structure
type ICallback interface {
	CreateCallback(ctx context.Context, data *Callback) error
	GetCallback(ctx context.Context, limit int, offset int) (*Callback, error)
	GetTotalCallbacks(ctx context.Context) (int, error)
	DeleteCallbackByID(ctx context.Context, id int32) error
}

// Callback represents model structure
type Callback struct {
	ID        int32     `db:"id"`
	Name      string    `db:"name"`
	Phone     string    `db:"phone"`
	Type      string    `db:"type"`
	Idea      string    `db:"idea"`
	CreatedAt time.Time `db:"created_at"`
}

// CreateCallback implements ICallback method for callback creation
func (d *Database) CreateCallback(ctx context.Context, data *Callback) error {
	_, err := d.Pool.Exec(ctx, `
        INSERT INTO callbacks (name, phone, type, idea) 
        VALUES ($1, $2, $3, $4)`,
		data.Name, data.Phone, data.Type, data.Idea)
	return err
}

// GetCallback implements ICallback method for getting callback with defined Limit and Offset variables
func (d *Database) GetCallback(ctx context.Context, limit int, offset int) (*Callback, error) {
	rows, err := d.Pool.Query(ctx, `
		SELECT id, name, type, phone, idea, created_at
		FROM callbacks
		ORDER BY created_at ASC
		LIMIT $1 OFFSET $2`,
		limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var callback Callback
	for rows.Next() {
		if err := rows.Scan(&callback.ID, &callback.Name, &callback.Type, &callback.Phone, &callback.Idea, &callback.CreatedAt); err != nil {
			return nil, err
		}
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return &callback, nil
}

// GetTotalCallbacks implements ICallback method for getting total callback rows count in db
func (d *Database) GetTotalCallbacks(ctx context.Context) (int, error) {
	row, err := d.Pool.Query(ctx, `
		SELECT COUNT(*) as total
	FROM callbacks`)
	if err != nil {
		return 0, err
	}
	defer row.Close()

	var total int
	for row.Next() {
		if err := row.Scan(&total); err != nil {
			return 0, err
		}
	}
	return total, nil
}

// DeleteCallbackByID implements ICallback method for deleting callback by ID
func (d *Database) DeleteCallbackByID(ctx context.Context, id int32) error {
	query := `DELETE FROM callbacks WHERE id = $1`
	commit, err := d.Pool.Exec(ctx, query, id)

	if commit.RowsAffected() != 1 {
		return errors.New("object not found")
	}
	return err
}
