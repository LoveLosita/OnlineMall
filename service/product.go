package service

import (
	"OnlineMall/auth"
	"OnlineMall/dao"
	"OnlineMall/model"
	"OnlineMall/respond"
	"math"
)

func AddProduct(product model.Product, handlerID int) error {
	role, err := auth.CheckPermission(handlerID) //检查用户权限
	if err != nil {                              //如果出错
		return err
	}
	if role != "merchant" && role != "admin" { //如果不是商家
		return respond.ErrUnauthorized //返回错误
	}
	result, err := dao.CheckIfCategoryExists(product.CategoryID) //检查商品分类是否存在
	if err != nil {
		return err
	}
	if !result { //如果不存在
		return respond.ErrCategoryNotExists //返回错误
	}
	maxNameLength := 80
	maxDescriptionLength := 10000 // typical max length for text type in MySQL
	maxPrice := 9999999999.99
	maxStock := math.MaxInt32
	//检查商品信息是否合法
	if len(product.Name) > int(0.9*float64(maxNameLength)) ||
		len(product.Description) > int(0.9*float64(maxDescriptionLength)) ||
		product.Price > 0.9*maxPrice ||
		product.Stock > int(0.9*float64(maxStock)) {
		return respond.ParamTooLong
	}
	return dao.AddProduct(product) //调用dao层函数
}
