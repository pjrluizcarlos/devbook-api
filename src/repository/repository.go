package repository

import (
	"database/sql"
	"devbook-api/src/database"
	"log"
)

func GetConnection() *sql.DB {
	db, error := database.Connect(); if error != nil {
		log.Fatal("Error connecting to database", error)
	}

	return db
}
