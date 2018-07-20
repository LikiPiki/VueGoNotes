package main

import (
	"Notes/server/db"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func initRoutes(r *mux.Router) {
	r.HandleFunc("/test", Test).Methods("GET")
}

func Test(w http.ResponseWriter, r *http.Request) {
	users, err := db.GetAll()

	if err != nil {
		fmt.Println("Can not get users")
	}

	json.NewEncoder(w).Encode(map[string]interface{}{
		"users":   users,
		"success": true,
	})
}
