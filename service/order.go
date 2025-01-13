package service

import (
	"OnlineMall/dao"
	"OnlineMall/model"
	"OnlineMall/respond"
)

func PlaceAnOrder(order model.PlaceOrder) (model.ReturnOrder, error) {
	//1.首先检验数量是否合法
	for i := 0; i < len(order.Items); i++ {
		if order.Items[i].Quantity < 1 || order.Items[i].Quantity > 999 {
			return model.ReturnOrder{}, respond.ErrQuantityTooLarge
		}
	}
	//2.计算商品单价*数量
	var product model.ShowProduct
	var err error
	for i := 0; i < len(order.Items); i++ {
		product, err = dao.GetProductInfoByID(order.Items[i].ProductID, 0)
		if err != nil {
			return model.ReturnOrder{}, err
		}
		order.Items[i].Price = product.Price * float64(order.Items[i].Quantity)
	}
	//3.计算总价
	order.TotalPrice = 0
	for i := 0; i < len(order.Items); i++ {
		order.TotalPrice += order.Items[i].Price
	}
	//4.把订单信息存入数据库
	order.Status = "completed" //订单状态设置为已完成
	returnOrder, err := dao.PlaceAnOrder(order)
	if err != nil {
		return model.ReturnOrder{}, err
	}
	return returnOrder, nil
}
