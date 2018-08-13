package routes

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/likipiki/VueGoNotes/server/db"
)

func GetAllTags(w http.ResponseWriter, r *http.Request) {
	userID := r.Header.Get("user_id")

	tags, err := db.Tag{}.GetAll(userID)
	if err != nil {
		ErrorJSON(w, err.Error())
		return
	}

	json.NewEncoder(w).Encode(map[string]interface{}{
		"success": "success",
		"tags":    tags,
	})
}

func DeleteTag(w http.ResponseWriter, r *http.Request) {
	userID := r.Header.Get("user_id")
	var tag db.Tag
	err := tag.Remove(userID)
	if err != nil {
		ErrorJSON(w, err.Error())
		return
	}

	json.NewEncoder(w).Encode(map[string]interface{}{
		"status": "success",
	})
}

func CreateTag(w http.ResponseWriter, r *http.Request) {
	userID := r.Header.Get("user_id")

	var tag db.Tag
	err := json.NewDecoder(r.Body).Decode(&tag)
	if err != nil {
		ErrorJSON(w, err.Error())
		return
	}
	var parsed uint64
	parsed, err = strconv.ParseUint(userID, 10, 32)

	if err != nil {
		ErrorJSON(w, err.Error())
		return
	}

	tag.UserID = uint(parsed)

	if err != nil {
		ErrorJSON(w, err.Error())
		return
	}
	err = tag.Create()

	if err != nil {
		ErrorJSON(w, err.Error())
	}

	json.NewEncoder(w).Encode(map[string]interface{}{
		"status": "success",
	})
}
