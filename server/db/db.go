package db

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

var (
	DB *sql.DB
)

// Connect to postgres database
func Connect() (db *sql.DB) {
	var err error

	connStr := "password='postgres' dbname=notes sslmode=disable"
	DB, err = sql.Open("postgres", connStr)

	if err != nil {
		panic(err)
	}

	log.Println("Connect to postgress success")

	err = chechFirstConnect()

	if err != nil {
		panic(err)
	}

	return DB

}

func chechFirstConnect() error {
	var id int
	err := DB.QueryRow(
		"SELECT id FROM users WHERE username = $1",
		"admin",
	).Scan(&id)

	if id == 0 {

		newUser := User{
			Username: "admin",
			Password: "admin",
			IsAdmin:  true,
		}

		err = newUser.Create()
		if err != nil {
			return err
		}

	}

	return nil
}
