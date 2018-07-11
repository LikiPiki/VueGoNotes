package main

import (
  "Notes/server/config"

	"net/http"
	"log"

	"github.com/gorilla/mux"
  "github.com/globalsign/mgo"
)

var conf config.Config

var (
  col_users *mgo.Collection
  col_notes *mgo.Collection
)

func init() {
	// load config from file
	err := config.ReadConfig()

	if err != nil {
		panic(err)
	}

}

func main() {

  // init mongodb
  session, err := mgo.Dial("127.0.0.1:27017")
  if err != nil {
    panic(err)
  }
  defer session.Close()
  session.SetSafe(&mgo.Safe{})

  col_users = session.DB("notes").C("users")
  col_notes = session.DB("notes").C("notes")

  // run web server
	r := mux.NewRouter()

	initRoutes(r)

	// default handler for dev only

  //http.Handle("/", http.FileServer(http.Dir("../dist")))
  http.Handle("/", r)

  println("server listening on port", config.Cnf.Port)
	err = http.ListenAndServe(config.Cnf.Port, nil)

	if err != nil {
		log.Println(err)
	}
}
