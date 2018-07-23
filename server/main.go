package main

import (
	"log"
	"net/http"

	"github.com/likipiki/VueGoNotes/server/db"
	"github.com/likipiki/VueGoNotes/server/routes"

	"github.com/gorilla/mux"
)

type bytes []byte

func main() {
	var err error

	database := db.Connect()
	defer database.Close()

	// routes
	r := mux.NewRouter().StrictSlash(true)
	routes.InitRoutes(r)

	http.Handle("/", r)

	log.Println("Listen and Serve on :3000")
	err = http.ListenAndServe(":3000", nil)
	if err != nil {
		log.Println("Cant run web-server")
	}

}
