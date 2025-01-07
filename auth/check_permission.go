package auth

import "OnlineMall/dao"

func CheckPermission(handlerID int) (string, error) {
	user, err := dao.GetUserInfoByID(handlerID) //通过handlerID获取用户信息
	if err != nil {
		return "", err
	}
	return user.Role, nil
}
