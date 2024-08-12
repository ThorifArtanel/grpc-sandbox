package ddb

import (
	"database/sql"

	_ "github.com/marcboeker/go-duckdb"
)

func Conn() (*sql.DB, error) {
	db, err := sql.Open("duckdb", "internal/app/ddb/default.duckdb")
	if err != nil {
		return nil, err
	}
	return db, nil
}
