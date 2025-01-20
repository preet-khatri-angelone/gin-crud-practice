package utils

import (
	"CRUD-GIN/db"
	"CRUD-GIN/model"
	"fmt"
)

func CreateUser(user *model.User) bool {
	hashedPassword := Encrypt(user.Password)
	if hashedPassword == "" {
		fmt.Println("error in password encryption")
		return false
	}
	user.Password = hashedPassword

 	_, err := db.DB.Exec(`INSERT INTO users(username, password) VALUES ($1, $2)`, user.Username, user.Password)
	if err != nil {
		fmt.Println("error in inserting user in DB", err)
		return false
	}
	return true
}

func FetchUser(user *model.User) (*model.User, bool)  {
	fetchedUser := &model.User{}
	row := db.DB.QueryRow(`SELECT * FROM users WHERE username = $1`, user.Username)

	if err := row.Scan(&fetchedUser.ID, &fetchedUser.Username, &fetchedUser.Password); err != nil {
		fmt.Println("error in fetching the user", err)
		return nil, false
	}

	if ok := Validate(user.Password, fetchedUser.Password); !ok {
		fmt.Println("error in validating the password")
		return nil, false
	}

	return fetchedUser, true
}
