package routes

import (
	"github.com/likipiki/VueGoNotes/server/crypt"
	"github.com/likipiki/VueGoNotes/server/db"

	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
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
	password := user.Password
	fmt.Println("pass", password)

	user, err = db.User{}.GetUserByUsername(user.Username)
	if err != nil || (user == db.User{}) {
		returnLoginFailed(w)
		return
	}

	valid := crypt.CheckPassword(user.Password, password)
	fmt.Println("valid is ", valid)

	if !valid {
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
		"status":   "success",
		"username": user.Username,
		"token":    token,
		"id":       user.Id,
	})

}

func Register(w http.ResponseWriter, r *http.Request) {
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
	fmt.Println("Creating", user)

	err = db.User{
		Username: user.Username,
		Password: user.Password,
	}.Create()

	if err != nil {
		returnErrStatus(w, "Cant create user")
		return
	}

	json.NewEncoder(w).Encode(map[string]interface{}{
		"status": "success",
	})

}
