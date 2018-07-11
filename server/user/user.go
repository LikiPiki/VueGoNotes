package user

import (
  "github.com/asaskevich/govalidator"
  "github.com/globalsign/mgo/bson"
  "github.com/globalsign/mgo"
  "fmt"
  "log"
)

type User struct {
  Collection *mgo.Collection `json:"-" bson:"-"`

  Id bson.ObjectId `json:"_id" bson:"_id,omitempty"`
	Username string `json:"username" valid:"required,alphanum,runelength(3|30)"`
	Password string `json:"password" valid:"required,alphanum,runelength(3|30)"`
}

func (us User) ValidateUser() (bool, error) {
  fl, err := govalidator.ValidateStruct(&us)
  return fl, err
}

func (us User) CheckLogin() (bool, User) {

  valid, err := us.ValidateUser()
  fmt.Println("valid")

  if !valid {
    return false, us
  }

  var res User
  err = us.Collection.Find(us).One(&res)
  if err != nil {
    log.Println("error find user", err)
    return false, res
  } else {
    return true, res
  }

  return false, res
}
