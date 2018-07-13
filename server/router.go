package main

import (
	"net/http"
	"encoding/json"
  "io/ioutil"
  "log"
  "fmt"

  "github.com/gorilla/mux"
  "Notes/server/user"
  "Notes/server/note"
  "github.com/globalsign/mgo/bson"
)

func initRoutes(r *mux.Router) {
  r.HandleFunc("/login", Login).Methods("POST")

	// private handle
  r.Handle("/createNote", jwtMiddleware.Handler(http.HandlerFunc(CreateNote))).Methods("POST")
  r.Handle("/getNotes/{id}", jwtMiddleware.Handler(http.HandlerFunc(GetNotes))).Methods("GET")
}

func CreateNote(w http.ResponseWriter, r *http.Request) {
  body, err := ioutil.ReadAll(r.Body)

  if err != nil {
    panic(err)
  }

  note := note.Note{
    Collection: col_notes,
  }

  err = json.Unmarshal(body, &note)
  if err != nil {
    panic(err)
  }

  fmt.Println(note)

  ok, err := note.CreateNote()

  if err != nil {
    panic(err)
  }

  json.NewEncoder(w).Encode(map[string]interface{}{
    "success": ok,
  })

}

func GetNotes(w http.ResponseWriter, r *http.Request) {
  vars := mux.Vars(r)
  id := vars["id"]

  note := note.Note{
    Collection: col_notes,
  }

  note.User = bson.ObjectIdHex(id)
  res, err := note.GetAllById()
  if err != nil {
    panic(err)
  }

  fmt.Println("result is", res)

  json.NewEncoder(w).Encode(res)
}

func Login(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)

  if err != nil {
    panic(err)
  }

	user := user.User {
	  Collection: col_users,
  }

	err = json.Unmarshal(body, &user)
  if err != nil {
    panic(err)
  }
	fmt.Println(user)

	var fl bool
	if fl, user = user.CheckLogin(); fl {
	  w.WriteHeader(http.StatusOK)

	  tokenStr, err := createToken(user.Username)
    fmt.Println(tokenStr)

	  if err != nil {
	    panic(err)
    }

    json.NewEncoder(w).Encode(map[string]interface{}{
      "status": true,
      "user": user.Username,
      "token": tokenStr,
      "_id": user.Id,
    })
  } else {
    json.NewEncoder(w).Encode(map[string]interface{}{
      "status": false,
      "error": "Error in login",
    })
  }

	if err != nil {
		panic(err)
	}

	log.Println(string(body))
}

