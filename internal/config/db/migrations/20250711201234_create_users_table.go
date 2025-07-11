package migrations

import (
	"context"
	"database/sql"

	"github.com/pressly/goose/v3"
)

func init() {
	goose.AddMigrationContext(upCreateUsersTable, downCreateUsersTable)
}

func upCreateUsersTable(ctx context.Context, tx *sql.Tx) error {
	_, err := tx.Exec(`
		CREATE TABLE users(
		id BIGSERIAL PRIMARY KEY,
		first_name VARCHAR(30) NOT NULL,
		last_name VARCHAR(30) NOT NULL,
		user_name VARCHAR(50) NOT NULL UNIQUE,
		email VARCHAR(150) NOT NULL UNIQUE,
		password CHAR(60) NOT NULL
	);`)

	if err != nil {
		return err
	}
	return nil
}

func downCreateUsersTable(ctx context.Context, tx *sql.Tx) error {
	_, err := tx.Exec("DROP TABLE users;")

	if err != nil {
		return err
	}
	return nil
}
