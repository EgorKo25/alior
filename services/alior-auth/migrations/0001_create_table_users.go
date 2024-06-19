package migrations

import (
	"context"
	"database/sql"

	"github.com/pressly/goose/v3"
)

func init() {
	goose.AddMigrationContext(upCreateTableUsers, downCreateTableUsers)
}

func upCreateTableUsers(ctx context.Context, tx *sql.Tx) error {
	_, err := tx.ExecContext(ctx, `
		CREATE TABLE users (
			id SERIAL PRIMARY KEY,
    		full_name VARCHAR(255),
    		email VARCHAR(255),
    		phone_number VARCHAR(12),
    		password VARCHAR(255)
		);
	`)
	return err
}

func downCreateTableUsers(ctx context.Context, tx *sql.Tx) error {
	// This code is executed when the migration is rolled back.
	_, err := tx.ExecContext(ctx, `DROP TABLE users`)
	return err
}
