package main

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (us User) CheckLogin() (bool) {
  if (us.Username == "admin" && us.Password == "admin") {
    return true
  }
  return false
}
