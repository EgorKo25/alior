package migrations

import (
	"context"
	"database/sql"
	"github.com/pressly/goose/v3"
)

func init() {
	goose.AddMigrationContext(upCreateCallbackTable, downCreateCallbackTable)
}

func upCreateCallbackTable(ctx context.Context, tx *sql.Tx) error {
	_, err := tx.ExecContext(ctx, `
		CREATE TABLE IF NOT EXISTS callbacks (
			id SERIAL PRIMARY KEY,
    		name VARCHAR(32) NOT NULL,
    		number VARCHAR(12) NOT NULL,
    		date TIMESTAMP NOT NULL
		);
	`)
	return err
}

func downCreateCallbackTable(ctx context.Context, tx *sql.Tx) error {
	_, err := tx.ExecContext(ctx, `DROP TABLE IF EXISTS callbacks`)
	return err
}
