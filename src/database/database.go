package database

import (
	"devbook-api/src/config"
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func Connect() (*sql.DB, error) {
	db, error := sql.Open("mysql", config.ConnectionString); if error != nil {
		return nil, error
	}

	error = db.Ping(); if error != nil {
		db.Close()
		return nil, error
	}

	return db, nil
}