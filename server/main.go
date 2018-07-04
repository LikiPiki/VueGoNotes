package main

import (
  "Notes/server/config"

	"net/http"
	"log"
	"github.com/gorilla/mux"
)

var conf config.Config

func init() {
	// load config from file
	err := config.ReadConfig()

	if err != nil {
		panic(err)
	}
}

func main() {

	var err error

	r := mux.NewRouter()

	initRoutes(r)
	http.Handle("/", r)

	println("server listening on port", config.Cnf.Port)
	err = http.ListenAndServe(config.Cnf.Port, nil)

	if err != nil {
		log.Println(err)
	}
}
