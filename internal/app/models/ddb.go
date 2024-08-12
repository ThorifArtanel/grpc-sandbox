package models

import (
	"database/sql"

	"github.com/google/uuid"
)

type DDB struct {
	DB *sql.DB
}

func (m *DDB) ReGenDB() error {
	_, err := m.DB.Exec("DROP TABLE IF EXISTS users")
	if err != nil {
		return err
	}

	_, err = m.DB.Exec("CREATE TABLE users(id VARCHAR, firstname VARCHAR, lastname VARCHAR, PRIMARY KEY (id))")
	if err != nil {
		return err
	}

	_, err = m.DB.Exec("INSERT INTO users(id, firstname, lastname) VALUES ($1, $2, $3)", uuid.New().String(), "Asep", "Saepuloh")
	if err != nil {
		return err
	}

	return nil
}
