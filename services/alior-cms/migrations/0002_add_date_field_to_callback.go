package migrations

import (
	"context"
	"database/sql"
	"github.com/pressly/goose/v3"
)

func init() {
	goose.AddMigrationContext(upAddDateFieldToCallback, downAddDateFieldToCallback)
}

func upAddDateFieldToCallback(ctx context.Context, tx *sql.Tx) error {
	_, err := tx.ExecContext(ctx, `
		ALTER TABLE callbacks
		ADD COLUMN created_at TIMESTAMP NOT NULL DEFAULT NOW();
	`)
	return err
}

func downAddDateFieldToCallback(ctx context.Context, tx *sql.Tx) error {
	_, err := tx.ExecContext(ctx, `
		ALTER TABLE callbacks
		DROP COLUMN created_at;
	`)
	return err
}
