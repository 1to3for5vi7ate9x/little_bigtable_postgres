package bttest

import (
	"context"

	"github.com/jmoiron/sqlx"
)

func CreateTables(ctx context.Context, db *sqlx.DB) error {
	// row_key is BYTEA (not TEXT): Bigtable row keys are arbitrary bytes and can
	// contain 0x00, which Postgres rejects in a text/utf8 column (SQLSTATE 22021).
	// bytea also compares byte-wise, matching Bigtable's lexicographic key ordering
	// used by the Ascend* range scans.
	query := `CREATE TABLE IF NOT EXISTS rows_t (
		parent TEXT NOT NULL,
		table_id TEXT NOT NULL,
		row_key BYTEA NOT NULL,
		families BYTEA NOT NULL,
		PRIMARY KEY (parent, table_id, row_key)
		)`
	_, err := db.ExecContext(ctx, query)
	if err != nil {
		return err
	}

	query = `CREATE TABLE IF NOT EXISTS tables_t (
		parent TEXT NOT NULL,
		table_id TEXT NOT NULL,
		metadata BYTEA NOT NULL,
		PRIMARY KEY  (parent, table_id)
		)`
	_, err = db.ExecContext(ctx, query)
	if err != nil {
		return err
	}
	return nil
}
