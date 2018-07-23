package routes

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/likipiki/VueGoNotes/server/db"
)

func GetAllNotes(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	notes, err := db.Note{}.GetAll(id)
	if err != nil {
		log.Println("err is ", err)
		returnLoginFailed(w)
		return
	}

	json.NewEncoder(w).Encode(map[string]interface{}{
		"success": "success",
		"notes":   notes,
	})
}

func CreateNote(w http.ResponseWriter, r *http.Request) {

	var note db.Note
	err := json.NewDecoder(r.Body).Decode(&note)

	if err != nil {
		returnErrStatus(w, "Cant json body")
		return
	}
	fmt.Println("note is ", note)

	var result bool
	result, err = db.Note{}.Create(note)
	log.Println("result is, ", err)
	if err != nil {
		returnErrStatus(w, "Cant create note")
	}

	json.NewEncoder(w).Encode(map[string]interface{}{
		"status":  "success",
		"created": result,
	})
}

func RemoveNote(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	err := db.Note{}.RemoveNote(id)
	if err != nil {
		returnErrStatus(w, "Cant delete by id ")
	}

	json.NewEncoder(w).Encode(map[string]interface{}{
		"status": "success",
	})

}
