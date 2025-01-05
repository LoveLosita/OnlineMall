package dao

import (
	"OnlineMall/model"
	"log"
)

func UserRegister(user model.User) error {
	query := `
		INSERT INTO users (
			username, email, password, full_name, phone_number, nickname, qq, avatar, gender, bio) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
	`
	_, err := Db.Exec(query, user.Username, user.Email, user.Password, user.FullName, user.PhoneNumber, user.Nickname,
		user.QQ, user.Avatar, user.Gender, user.Bio)
	if err != nil {
		log.Printf("Error inserting user: %v", err)
		return err
	}
	return nil
}

func IfUsernameExists(name string) (bool, error) {
	query := "SELECT id FROM users WHERE username=?"
	rows, err := Db.Query(query, name)
	if err != nil {
		return true, err
	}
	if rows.Next() {
		return true, nil
	} else {
		return false, nil
	}
}
