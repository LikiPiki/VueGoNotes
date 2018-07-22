package routes

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"projects/Notes/server/db"
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

	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		returnErrStatus(w, "Cant read body")
		return
	}

	var note db.Note
	fmt.Println("note is ", note)
	err = json.Unmarshal(body, &note)
	if err != nil {
		returnErrStatus(w, "Cant json body")
		return
	}
	fmt.Println("note is ", note)

	json.NewEncoder(w).Encode(map[string]interface{}{
		"status": "success",
	})
}
