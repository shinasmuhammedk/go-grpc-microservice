package db

import (
	"database/sql"
	"log"
)

func ConnectDB() *sql.DB {
	connStr := "user=postgres password=Shinas dbname=grpcUsers sslmode=disable"

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	return db
}
