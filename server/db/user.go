package db

import "fmt"

type User struct {
	Id uint `sql:"id,pk" json:"id"`

	Password string `sql:"password" json:"password"`
	Username string `sql:"username,unique" json:"username"`
	IsAdmin  bool   `sql:"is_admin" json:"is_admin"`
}
type Users []User

func (u User) GetAll() (Users, error) {

	rows, err := DB.Query(
	  "SELECT id, username, password, is_admin FROM users",
	)

	if err != nil {
		return nil, err
	}

	var users Users
	var user User

	for rows.Next() {
		err = rows.Scan(
			&user.Id,
			&user.Username,
			&user.Password,
			&user.IsAdmin,
		)

		if err != nil {
			return nil, err
		}

		users = append(users, user)

	}

	defer rows.Close()

	return users, nil
}

func (user User) GetUserByUsername(username string) (User, error) {
  fmt.Println("finding", username)
  err := DB.QueryRow(
    "SELECT id, username, password, is_admin FROM users WHERE username = $1",
    username,
  ).Scan(
    &user.Id,
    &user.Username,
    &user.Password,
    &user.IsAdmin,
  )

  if err != nil {
    return User{}, err
  }

  fmt.Println(user)

  return user, nil

}
