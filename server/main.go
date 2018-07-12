package main

import (
  "Notes/server/config"

	"net/http"
	"log"

	"github.com/gorilla/mux"
  "github.com/globalsign/mgo"
  "fmt"
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
	r := mux.NewRouter().StrictSlash(true)
	//api := r.PathPrefix("/api/").Subrouter()
	//initRoutes(api)
  initRoutes(r)

  // default handler for dev only
  //http.Handle("/", http.FileServer(http.Dir("../dist")))
  //r.PathPrefix("/").Handler(http.FileServer(http.Dir("../dist")))
  http.Handle("/", r)

  r.Walk(func(route *mux.Route, router *mux.Router, ancestors []*mux.Route) error {
    t, err := route.GetPathTemplate()
    if err != nil {
      return err
    }
    fmt.Println(t)
    return nil
  })

  println("server listening on port", config.Cnf.Port)
	err = http.ListenAndServe(config.Cnf.Port, nil)

	if err != nil {
		log.Println(err)
	}
}
