package routes

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"projects/Notes/server/db"
)

func Login(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		returnErrStatus(w, "Cant read body")
		return
	}

	var user db.User
	err = json.Unmarshal(body, &user)
	if err != nil {
		returnErrStatus(w, "Cant json body")
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
		"success":  "success",
		"username": user.Username,
		"token":    token,
		"id":       user.Id,
	})

}
