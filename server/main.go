package main

import (
	db "Notes/server/db"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type bytes []byte

func main() {
	var err error

	database := db.Connect()
	defer database.Close()

	// routes
	r := mux.NewRouter().StrictSlash(true)
	initRoutes(r)

	http.Handle("/", r)

	fmt.Println("Listen and Serve on :3000")
	err = http.ListenAndServe(":3000", nil)
	if err != nil {
		log.Println("Cant run web-server")
	}

}
