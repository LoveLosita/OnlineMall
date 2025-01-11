package service

import (
	"OnlineMall/auth"
	"OnlineMall/dao"
	"OnlineMall/model"
	"OnlineMall/respond"
	"math"
)

func AddProduct(product model.AddProduct, handlerID int) error {
	//1.检查用户权限
	role, err := auth.CheckPermission(handlerID) //检查用户权限
	if err != nil {                              //如果出错
		return err
	}
	if role != "merchant" && role != "admin" { //如果不是商家
		return respond.ErrUnauthorized //返回错误
	}
	//2.检查商品分类是否存在
	result, err := dao.CheckIfCategoryExists(product.CategoryID) //检查商品分类是否存在
	if err != nil {
		return err
	}
	if !result { //如果不存在
		return respond.ErrCategoryNotExists //返回错误
	}
	//3.检查商品信息是否合法
	maxNameLength := 80
	maxDescriptionLength := 10000 // typical max length for text type in MySQL
	maxPrice := 9999999999.99
	maxStock := math.MaxInt32
	if len(product.Name) > int(0.9*float64(maxNameLength)) ||
		len(product.Description) > int(0.9*float64(maxDescriptionLength)) ||
		product.Price > 0.9*maxPrice ||
		product.Stock > int(0.9*float64(maxStock)) {
		return respond.ParamTooLong
	}
	//4.保存商品信息
	return dao.AddProduct(product) //调用dao层函数
}

func ChangeProduct(id int, product model.AddProduct, handlerID int) error {
	//1.检查用户权限
	role, err := auth.CheckPermission(handlerID) //检查用户权限
	if err != nil {                              //如果出错
		return err
	}
	if role != "merchant" && role != "admin" { //如果不是商家及以上
		return respond.ErrUnauthorized //返回错误
	}
	//2.检查商品信息是否合法
	maxNameLength := 80
	maxDescriptionLength := 1000
	maxPrice := 9999999999.99
	maxStock := math.MaxInt32
	if len(product.Name) > int(0.9*float64(maxNameLength)) ||
		len(product.Description) > int(0.9*float64(maxDescriptionLength)) ||
		product.Price > 0.9*maxPrice ||
		product.Stock > int(0.9*float64(maxStock)) {
		return respond.ParamTooLong
	}
	//3.获取商品信息，并且判断是否已经填写对应信息，实现选择性更新
	oldProduct, err := dao.GetProductInfoByID(id) //同时可以检查商品是否存在
	if err != nil {
		return err
	}
	if product.Name == "" { //如果没有填写商品名
		product.Name = oldProduct.Name //则使用原来的商品名
	}
	if product.Description == "" { //如果没有填写商品描述
		product.Description = oldProduct.Description //则使用原来的商品描述
	}
	if product.Price == 0 {
		product.Price = oldProduct.Price
	}
	if product.Stock == 0 {
		product.Stock = oldProduct.Stock
	}
	if product.CategoryID == 0 {
		product.CategoryID = oldProduct.CategoryID
	} else {
		result, err := dao.CheckIfCategoryExists(product.CategoryID) //检查商品分类是否存在
		if err != nil {
			return err
		}
		if !result { //如果不存在
			return respond.ErrCategoryNotExists //返回错误
		}
	}
	if product.ProductImage == "" {
		product.ProductImage = oldProduct.ProductImage
	}
	return dao.UpdateProduct(id, product) //调用dao层函数
}

func ShowProductInManyWays(productID int, keyword string, categoryID int) ([]model.ShowProduct, error) {
	//优先级：商品id>关键字>分类id>全部商品
	if productID != 0 { //如果有商品id
		product, err := dao.GetProductInfoByID(productID)
		if err != nil {
			return nil, err
		}
		return []model.ShowProduct{product}, nil
	}
	if keyword != "" { //如果有关键字
		return dao.GetProductInfoByKeyWord(keyword)
	}
	if categoryID != 0 { //如果有分类id
		return dao.ShowACategoryProducts(categoryID)
	}
	return dao.ShowAllProducts() //返回所有商品
}

func DeleteProduct(id int, handlerID int) error {
	//1.检查用户权限
	role, err := auth.CheckPermission(handlerID) //检查用户权限
	if err != nil {                              //如果出错
		return err
	}
	if role != "merchant" && role != "admin" { //如果不是商家及以上
		return respond.ErrUnauthorized //返回错误
	}
	//2.检查商品是否存在
	_, err = dao.GetProductInfoByID(id) //检查商品是否存在
	if err != nil {
		return err
	}
	//3.删除商品
	return dao.DeleteProduct(id) //调用dao层函数
}
