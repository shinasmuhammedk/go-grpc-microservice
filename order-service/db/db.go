package db

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
)

func ConnectDB() *sql.DB {
	connStr := "user=postgres password=Shinas dbname=grpcmicroorders sslmode=disable"

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

    
	return db
}
