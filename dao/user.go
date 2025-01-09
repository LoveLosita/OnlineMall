package dao

import (
	"OnlineMall/model"
	"OnlineMall/respond"
	"log"
)

func UserRegister(user model.User) error { //注册用户
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

func IfUsernameExists(name string) (bool, error) { //检查用户名是否存在
	query := "SELECT id FROM users WHERE username=?"
	rows, err := Db.Query(query, name)
	if err != nil {
		return true, err
	}
	if rows.Next() { //如果有这个用户
		return true, nil
	} else {
		return false, nil
	}
}

func GetUserHashedPassword(username string) (string, error) { //获取用户密码
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

func GetUserID(username string) (int, error) { //获取用户id
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

func GetUserName(id int) (string, error) {
	var userName string
	query := "SELECT username FROM users WHERE id=?"
	rows, err := Db.Query(query, id)
	if err != nil {
		return "", err
	}
	if rows.Next() {
		err = rows.Scan(&userName)
		if err != nil {
			return "", err
		}
		return userName, nil
	}
	return "", respond.WrongUserID
}

func GetUserInfoByID(id int) (model.User, error) { //通过id获取用户信息
	var user model.User
	query := "SELECT id, username, email, full_name, phone_number, nickname, qq, avatar,gender,bio,role FROM users WHERE id=?"
	rows, err := Db.Query(query, id)
	if err != nil {
		return user, err
	}
	if rows.Next() { //如果有这个用户
		err = rows.Scan(&user.ID, &user.Username, &user.Email, &user.FullName, &user.PhoneNumber, &user.Nickname, &user.QQ,
			&user.Avatar, &user.Gender, &user.Bio, &user.Role)
		if err != nil {
			return user, err
		}
		return user, nil
	}
	return user, respond.WrongUserID //找不到用户
}

func ChangeUserPasswordOrName(id int, password, name string) error { //修改用户密码或用户名
	query := "UPDATE users SET password=?, username=? WHERE id=?"
	_, err := Db.Exec(query, password, name, id)
	if err != nil {
		return err
	}
	return nil
}

func UpdateUserInfo(id int, user model.ChangeInfoUser) error { //更新用户信息
	query := "UPDATE users SET email=?, full_name=?, phone_number=?, nickname=?, qq=?,avatar=?,gender=?,bio=? WHERE id=?"
	_, err := Db.Exec(query, user.Email, user.FullName, user.PhoneNumber, user.Nickname, user.QQ, user.Avatar, user.Gender, user.Bio, id)
	if err != nil {
		return err
	}
	return nil
}
