package service

import (
	"OnlineMall/dao"
	"OnlineMall/respond"
)

func AddProductToCart(userID, productID, quantity int) error {
	//1.检查商品是否存在
	_, err := dao.GetProductInfoByID(productID)
	if err != nil {
		return err
	}
	//2.检查商品是否已经在购物车中
	exists, err := dao.IfProductExistsInYourCart(productID, userID)
	if err != nil {
		return err
	}
	//3.如果存在，获取商品数量
	if exists {
		product, err := dao.GetSingleProductInCart(userID, productID)
		if err != nil {
			return err
		}
		if product.Quantity == quantity { //如果在购物车中并且数量相同，认定为重复添加，阻止添加
			return respond.ErrProductAlreadyInCart
		} //如果在购物车中但数量不同，则进行下一步
	}
	//4.如果在购物车中但数量不同，更新购物车中商品数量为新数量
	err = dao.UpdateProductQuantityInCart(userID, productID, quantity)
	if err != nil {
		return err
	}
	//5.如果不在购物车中，添加到购物车
	err = dao.AddToCart(userID, productID, quantity)
	if err != nil {
		return err
	}
	return nil
}
