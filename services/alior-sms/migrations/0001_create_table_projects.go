package migrations

import (
	"context"
	"database/sql"

	"github.com/pressly/goose/v3"
)

func init() {
	goose.AddMigrationContext(upCreateTableProjects, downCreateTableProjects)
}

func upCreateTableProjects(ctx context.Context, tx *sql.Tx) error {
	_, err := tx.ExecContext(ctx, `
		CREATE TABLE Projects (
			id SERIAL PRIMARY KEY,
			URL TEXT,
    		description TEXT,
		);
	`)

	return err
}

func downCreateTableProjects(ctx context.Context, tx *sql.Tx) error {
	_, err := tx.ExecContext(ctx, `DROP TABLE Projects`)
	return err
}
