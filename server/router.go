package main

import (
	"Notes/server/db"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
  "io/ioutil"
  "log"
)

func initRoutes(r *mux.Router) {
  r.HandleFunc("/login", Login).Methods("POST")
}

func Login(w http.ResponseWriter, r *http.Request) {
  body, err := ioutil.ReadAll(r.Body)

  if err != nil {
    returnErrStatus(w,"Cant read body")
    return
  }

  var user db.User
  err = json.Unmarshal(body, &user)
  if err != nil {
    returnErrStatus(w,"Cant json body")
    return
  }

  user, err = db.User{}.GetUserByUsername(user.Username)
  if err != nil || (user == db.User{}) {
    log.Println("login err")
    returnLoginFailed(w)
    return
  }
  var token string
  token, err = createToken(user.Username)
  if err != nil {
    log.Println("create token err")
    returnLoginFailed(w)
    return
  }

  json.NewEncoder(w).Encode(map[string]interface{}{
    "success": "success",
    "username": user.Username,
    "token": token,
    "id": user.Id,
  })

}

/*
 *  Utils
 */

func returnErrStatus(w http.ResponseWriter, errorString string) {
  err := json.NewEncoder(w).Encode(map[string]interface{}{
    "status": "error",
    "error": errorString,
  })
  if err != nil {
    log.Println("Cant encode error response", err)
  }
}

func returnLoginFailed(w http.ResponseWriter) {
  err := json.NewEncoder(w).Encode(map[string]interface{}{
    "login": "failed",
  })
  if err != nil {
    log.Println("Cant encode error response", err)
  }
}
