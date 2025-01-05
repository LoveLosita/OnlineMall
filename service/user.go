package service

import (
	"OnlineMall/dao"
	"OnlineMall/model"
	"OnlineMall/respond"
)

func UserRegister(user model.User) error {
	if user.Username == "" || user.Email == "" || user.Password == "" || user.FullName == "" ||
		user.PhoneNumber == "" || user.Nickname == "" || user.QQ == "" || user.Avatar == "" ||
		user.Gender == "" || user.Bio == "" { //检查是否有空字段
		return respond.MissingParam
	}
	result, err := dao.IfUsernameExists(user.Username) //调用dao层的方法
	if err != nil {
		return err
	}
	if result {
		return respond.InvalidID
	}
	err = dao.UserRegister(user) //调用dao层的方法
	if err != nil {
		return err
	}
	return nil
}
