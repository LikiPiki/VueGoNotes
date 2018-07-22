package db

import (
	"database/sql"
	"log"
	"projects/Notes/server/crypt"

	_ "github.com/lib/pq"
)

var (
	DB *sql.DB
)

// Connect to postgres database
func Connect() (db *sql.DB) {
	var err error

	connStr := "user=admin password='admin' dbname=notes host=localhost sslmode=disable"
	DB, err = sql.Open("postgres", connStr)

	if err != nil {
		panic(err)
	}

	log.Println("Connect to postgress success")

	chechFirstConnect()

	return DB

}

func chechFirstConnect() {
	var id int
	err := DB.QueryRow(
		"SELECT id FROM users WHERE username = $1",
		"admin",
	).Scan(&id)

	if id == 0 {
		pass, _ := crypt.CryptPassword("admin")
		_, err = DB.Query(
			"INSERT INTO users (username, password, is_admin) VALUES ($1, $2, $3)",
			"admin", pass, true,
		)
		if err != nil {
			panic(err)
		}
	}

}
