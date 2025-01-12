package dao

import (
	"OnlineMall/model"
	"fmt"
	"time"
)

func PlaceAnOrder(order model.PlaceOrder) (model.ReturnOrder, error) {
	//1.先把订单信息存入orders表
	query := "INSERT INTO orders(user_id,total_price,status,address) VALUES(?,?,?,?)"
	_, err := Db.Exec(query, order.UserID, order.TotalPrice, order.Status, order.Address)
	if err != nil {
		fmt.Println(2)
		return model.ReturnOrder{}, err
	}
	//2.获取刚刚插入的订单的id
	query = "SELECT id FROM orders WHERE user_id=? AND total_price=? AND status=? AND address=?"
	rows, err := Db.Query(query, order.UserID, order.TotalPrice, order.Status, order.Address)
	if err != nil {
		return model.ReturnOrder{}, err
	}
	var orderID int
	for rows.Next() { //只有一条记录
		err = rows.Scan(&orderID)
		if err != nil {
			return model.ReturnOrder{}, err
		}
	}
	//3.把订单中的商品信息存入order_products表
	for _, item := range order.Items {
		query = "INSERT INTO order_items(order_id,product_id,quantity,price) VALUES(?,?,?,?)"
		_, err = Db.Exec(query, orderID, item.ProductID, item.Quantity, item.Price)
		if err != nil {
			return model.ReturnOrder{}, err
		}
	}
	//4.获取数据库中的订单时间
	query = "SELECT order_date FROM orders WHERE id=?"
	rows, err = Db.Query(query, orderID)
	if err != nil {
		return model.ReturnOrder{}, err
	}
	var orderDate time.Time
	for rows.Next() { //只有一条记录
		err = rows.Scan(&orderDate)
		if err != nil {
			return model.ReturnOrder{}, err
		}
	}
	//5.返回订单信息
	return model.ReturnOrder{
		ID:        orderID,
		Status:    order.Status,
		OrderDate: orderDate.Format("2006-01-02 15:04:05"),
	}, nil
}
