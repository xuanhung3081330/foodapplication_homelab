package config

import (
	"database/sql"
	_ "github.com/lib/pq"
)

var DB *sql.DB

func ConnectDB() error {
	connStr := ""
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return err
	}

	if err = db.Ping(); err != nil {
		return err
	}

	DB = db
	return nil
}
