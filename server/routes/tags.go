package routes

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/likipiki/VueGoNotes/server/db"
)

func GetAllTags(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID := vars["user_id"]

	tags, err := db.Tag{}.GetAll(userID)
	if err != nil {
		log.Println("err is ", err)
		returnLoginFailed(w)
		return
	}

	json.NewEncoder(w).Encode(map[string]interface{}{
		"success": "success",
		"notes":   tags,
	})
}

func CreateTag(w http.ResponseWriter, r *http.Request) {

	var tag db.Tag
	err := json.NewDecoder(r.Body).Decode(&tag)

	if err != nil {
		returnErrStatus(w, "Cant json body")
		return
	}
	fmt.Println("tag is ", tag)

	err = tag.Create()
	log.Println("result is, ", err)

	if err != nil {
		returnErrStatus(w, "Cant create note")
	}

	json.NewEncoder(w).Encode(map[string]interface{}{
		"status": "success",
	})
}
