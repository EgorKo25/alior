package migrations

import (
	"context"
	"database/sql"

	"github.com/pressly/goose/v3"
)

func init() {
	goose.AddMigrationContext(upCreateTableServices, downCreateTableServices)
}

func upCreateTableServices(ctx context.Context, tx *sql.Tx) error {
	_, err := tx.ExecContext(ctx, `
		CREATE TABLE services (
			id SERIAL PRIMARY KEY,
    		description TEXT,
    		price NUMERIC CHECK (price >= 0)
		);
	`)
	return err
}

func downCreateTableServices(ctx context.Context, tx *sql.Tx) error {
	_, err := tx.ExecContext(ctx, `DROP TABLE services`)
	return err
}
