package main

import (
	"github.com/auth0/go-jwt-middleware"
	"github.com/dgrijalva/jwt-go"
	"time"
)

type JwtUser struct {
	Username string
	Expire   time.Time
}

func createToken(username string) (string, error) {
	var mySigningKey = []byte("secret")

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"name": username,
		"exp":  time.Now().Add(time.Hour * 24).Unix(),
	})

	tokenStr, err := token.SignedString(mySigningKey)
	return tokenStr, err
}

func checkTokenValid(token string) (bool, error) {
	return false, nil
}

var jwtMiddleware = jwtmiddleware.New(jwtmiddleware.Options{
	ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
		return []byte("secret"), nil
	},
	SigningMethod: jwt.SigningMethodHS256,
})
