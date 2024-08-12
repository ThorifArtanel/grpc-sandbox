package postgres

import (
	"database/sql"

	_ "github.com/lib/pq"
)

// Not used, here as a placeholder example
func Conn() (*sql.DB, error) {
	db, err := sql.Open("postgres", "")
	if err != nil {
		return nil, err
	}
	return db, nil
}
