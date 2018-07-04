package main

import (
	"net/http"
	"encoding/json"
  "io/ioutil"
  "log"
  "fmt"
  "time"

	"github.com/gorilla/mux"
  "github.com/dgrijalva/jwt-go"
  "github.com/auth0/go-jwt-middleware"
)

var myHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
  println("worked")
})

var jwtMiddleware = jwtmiddleware.New(jwtmiddleware.Options{
  ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
    return []byte("secret"), nil
  },
  SigningMethod: jwt.SigningMethodHS256,
})

func initRoutes(r *mux.Router) {
  r.HandleFunc("/login", Login).Methods("POST")
	r.HandleFunc("/", Test).Methods("GET")

	// private handle
  r.Handle("/private", jwtMiddleware.Handler(http.HandlerFunc(Private))).Methods("GET")
}

func Test(w http.ResponseWriter, r *http.Request) {

	result, err := json.MarshalIndent(map[string]interface{}{
		"success": true,
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
  println("logining")
	body, err := ioutil.ReadAll(r.Body)

  if err != nil {
    panic(err)
  }

	var user User
	err = json.Unmarshal(body, &user)
	fmt.Println(user)

	if (user.CheckLogin()) {
	  w.WriteHeader(http.StatusOK)

    var mySigningKey = []byte("secret")
    // create token
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
      "name": "admin",
      "exp": time.Now().Add(time.Hour * 24).Unix(),
    })
    tokenStr, err := token.SignedString(mySigningKey)
    if err != nil {
      panic(err)
    }

    res, err := json.MarshalIndent(map[string]interface{}{
      "success": true,
      "token": tokenStr,
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
