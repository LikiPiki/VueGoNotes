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
  r.Handle("/private", jwtMiddleware.Handler(http.HandlerFunc(Private))).Methods("GET")
  r.Handle("/getNotes/{id}", jwtMiddleware.Handler(http.HandlerFunc(GetNotes))).Methods("GET")
}

func CreateNote(w http.ResponseWriter, r *http.Request) {
  body, err := ioutil.ReadAll(r.Body)

  if err != nil {
    panic(err)
  }

  var note note.Note
  note.Collection = col_notes

  err = json.Unmarshal(body, &note)
  if err != nil {
    panic(err)
  }
  fmt.Println(note)
  ok := note.CreateNote()

  result, err := json.MarshalIndent(map[string]interface{}{
    "ok": ok,
  }, "", "\t")

  if err != nil {
    panic(err)
  }

  w.Write(result)
}

func GetNotes(w http.ResponseWriter, r *http.Request) {
  vars := mux.Vars(r)
  id := vars["id"]
  note := note.Note{Collection: col_notes}
  note.User = bson.ObjectIdHex(id)
  fmt.Println("vars is ", id)
  fl := true
  res, err := note.GetAllById(id)
  if err != nil {
    fl = false
  }
  fmt.Println("result is", res)

  result, err := json.MarshalIndent(map[string]interface{}{
    "ok": fl,
    "notes": res,
  }, "", "\t")

  if err != nil {
    panic(err)
  }

  w.Write(result)
}

func Private(w http.ResponseWriter, r *http.Request) {
  result, err := json.MarshalIndent(map[string]interface{}{
    "success": "this is a private page",
  }, "", "\t")

  if err != nil {
    panic(err)
  }

  w.Write(result)
}

func Login(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)

  if err != nil {
    panic(err)
  }

	var user user.User
	user.Collection = col_users

	err = json.Unmarshal(body, &user)
  if err != nil {
    panic(err)
  }
	fmt.Println(user)

	var fl bool
	if fl, user = user.CheckLogin(); fl {
	  fmt.Println("login success")
	  w.WriteHeader(http.StatusOK)

	  tokenStr, err := createToken(user.Username)
    fmt.Println(tokenStr)

    res, err := json.MarshalIndent(map[string]interface{}{
      "success": true,
      "token": tokenStr,
      "user": user.Username,
      "user_id": user.Id,
    }, "", "\t")

	  if err != nil {
	    panic(err)
    }
	  w.Write(res)
  } else {
    w.WriteHeader(http.StatusBadRequest)
  }

	if err != nil {
		panic(err)
	}

	log.Println(string(body))
}

