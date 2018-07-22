package routes

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// default open routes
func InitRoutes(r *mux.Router) {
	r.HandleFunc("/login", Login).Methods("POST")

	JWTRoutes(r)
}

// JWT private routes
func JWTRoutes(r *mux.Router) {
	// notes handlers
	r.Handle("/getNotes/{id}", jwtMiddleware.Handler(http.HandlerFunc(GetAllNotes))).Methods("GET")
	r.Handle("/createNote/{id}", jwtMiddleware.Handler(http.HandlerFunc(CreateNote))).Methods("POST")
}

/*
 *  Utils
 */
func returnErrStatus(w http.ResponseWriter, errorString string) {
	err := json.NewEncoder(w).Encode(map[string]interface{}{
		"status": "error",
		"error":  errorString,
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
