package db

import (
	"log"

	"database/sql"
	_ "github.com/lib/pq"
)

var (
	DB *sql.DB
)

func Connect() (db *sql.DB) {
	var err error

	connStr := "user=sergey dbname=test sslmode=disable"
	DB, err = sql.Open("postgres", connStr)

	if err != nil {
		panic(err)
	}

	log.Println("Connect to postgress success")

	return DB

}
