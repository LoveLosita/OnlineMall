package service

import (
	"OnlineMall/dao"
	"OnlineMall/model"
	"OnlineMall/respond"
	"OnlineMall/utils"
)

func UserRegister(user model.User) error {
	//检查是否有空字段
	if user.Username == "" || user.Email == "" || user.Password == "" || user.FullName == "" ||
		user.PhoneNumber == "" || user.Nickname == "" || user.QQ == "" || user.Avatar == "" ||
		user.Gender == "" || user.Bio == "" {
		return respond.MissingParam
	}
	// 检查字段长度是否超过90%
	if len(user.Username) > 45 || len(user.Email) > 90 || len(user.Password) > 229 ||
		len(user.FullName) > 90 || len(user.PhoneNumber) > 18 || len(user.Nickname) > 45 ||
		len(user.QQ) > 18 || len(user.Avatar) > 229 {
		return respond.ParamTooLong
	}
	//检查性别是否合法
	if user.Gender != "male" && user.Gender != "female" && user.Gender != "other" {
		return respond.WrongGender
	}
	result, err := dao.IfUsernameExists(user.Username) //调用dao层的方法
	if err != nil {
		return err
	}
	if result {
		return respond.InvalidName
	}
	hashedPwd, err := utils.HashPassword(user.Password) //调用utils层的方法
	if err != nil {
		return err
	}
	user.Password = hashedPwd    //将user的密码字段改为加密后的密码
	err = dao.UserRegister(user) //调用dao层的方法
	if err != nil {
		return err
	}
	return nil
}

func UserLogin(user model.LoginUser) (bool, string, error) {
	hashedPwd, err := dao.GetUserHashedPassword(user.Username) //调用dao层的方法
	if err != nil {
		return false, "", err
	}
	result, err := utils.CompareHashPwdAndPwd(hashedPwd, user.Password) //比较密码是否匹配
	if err != nil {                                                     //其他错误
		return false, "", err
	} else if !result { //密码不匹配
		return false, "", respond.WrongPwd
	}
	id, err := dao.GetUserID(user.Username) //获取用户id
	if err != nil {
		return false, "", err
	}
	jwtkey, err := utils.GenerateJWT(id) //生成jwt key
	if err != nil {                      //其他错误
		return false, "", err
	}
	return true, jwtkey, nil
}

func ChangeUserPwdOrName(handlerID int, user model.ChangePasswordAndUsernameUser) error {
	if user.OldPassword == "" { //检查是否有空字段
		return respond.MissingParam //返回参数过少的错误
	}
	//先通过传入的handlerID获取用户名
	name, err := dao.GetUserName(handlerID)
	if err != nil {
		return err
	}
	//首先检验旧密码是否正确
	hashedPwd, err := dao.GetUserHashedPassword(name) //调用dao层的方法
	if err != nil {
		return err
	}
	result, err := utils.CompareHashPwdAndPwd(hashedPwd, user.OldPassword) //比较密码是否匹配
	if err != nil {                                                        //其他错误
		return err
	} else if !result { //密码不匹配
		return respond.WrongPwd
	}
	//如果有的信息为空，说明不修改，需要调取原来的信息来填上
	if user.NewPassword == "" {
		user.NewPassword = hashedPwd
	} else { //非空，给密码加密
		hashedPassword, err := utils.HashPassword(user.NewPassword)
		if err != nil {
			return err
		}
		user.NewPassword = hashedPassword
	}
	if user.NewUsername == "" {
		oldUser, err := dao.GetUserInfoByID(handlerID)
		if err != nil {
			return err
		}
		user.NewUsername = oldUser.Username
	}
	//接下来才开始填入信息
	err = dao.ChangeUserPasswordOrName(handlerID, user.NewPassword, user.NewUsername)
	if err != nil {
		return err
	}
	return nil
}
