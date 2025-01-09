package service

import (
	"OnlineMall/auth"
	"OnlineMall/dao"
	"OnlineMall/respond"
)

func AddCategory(handlerID int, name string, description string) error {
	maxNameLength := 80
	maxDescriptionLength := 1000
	//检查分类信息是否合法
	if len(name) > int(0.9*float64(maxNameLength)) || len(description) > int(0.9*float64(maxDescriptionLength)) {
		return respond.ParamTooLong
	}
	role, err := auth.CheckPermission(handlerID) //检查用户权限
	if err != nil {                              //如果出错
		return err
	}
	if role != "merchant" && role != "admin" { //如果不是商家
		return respond.ErrUnauthorized //返回错误
	}
	//检查分类名是否已存在
	result, err := dao.CheckIfCategoryNameExists(name)
	if err != nil {
		return err
	}
	if result { //如果存在
		return respond.ErrCategoryNameExists //返回错误
	}
	return dao.AddCategory(name, description) //调用dao层函数
}
