package dao

import (
	"OnlineMall/model"
	"OnlineMall/respond"
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

func GetUserHashedPassword(username string) (string, error) {
	var password string
	query := "SELECT password FROM users WHERE username=?"
	rows, err := Db.Query(query, username)
	if err != nil {
		return "", err
	}
	if rows.Next() { //如果有这个用户
		err = rows.Scan(&password)
		if err != nil {
			return "", err
		}
		return password, nil
	}
	return "", respond.WrongName //找不到用户
}

func GetUserID(username string) (int, error) {
	var id int
	query := "SELECT id FROM users WHERE username=?"
	rows, err := Db.Query(query, username)
	if err != nil {
		return 0, err
	}
	if rows.Next() { //如果有这个用户
		err = rows.Scan(&id) //将用户id赋值给id
		if err != nil {
			return 0, err
		}
		return id, nil
	}
	return 0, respond.WrongName //找不到用户
}
