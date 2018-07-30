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
	r.HandleFunc("/register", Register).Methods("POST")

	JWTRoutes(r)
}

// JWT private routes
func JWTRoutes(r *mux.Router) {
	r.Handle("/tags", jwtMiddleware.Handler(http.HandlerFunc(GetAllTags))).Methods("GET")
	r.Handle("/tags", jwtMiddleware.Handler(http.HandlerFunc(CreateTag))).Methods("POST")
	// notes handlers
	r.Handle("/notes/{id}/{note_id}", jwtMiddleware.Handler(http.HandlerFunc(GetOneNote))).Methods("GET")
	r.Handle("/notes/{id}", jwtMiddleware.Handler(http.HandlerFunc(GetAllNotes))).Methods("GET")
	r.Handle("/notes/{id}", jwtMiddleware.Handler(http.HandlerFunc(RemoveNote))).Methods("DELETE")
	r.Handle("/notes", jwtMiddleware.Handler(http.HandlerFunc(CreateNote))).Methods("POST")
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
