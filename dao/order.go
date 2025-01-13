package dao

import (
	"OnlineMall/model"
	"OnlineMall/respond"
	"database/sql"
	"time"
)

func PlaceAnOrder(order model.PlaceOrder) (model.ReturnOrder, error) {
	//1.先把订单信息存入orders表
	query := "INSERT INTO orders(user_id,total_price,status,address) VALUES(?,?,?,?)"
	_, err := Db.Exec(query, order.UserID, order.TotalPrice, order.Status, order.Address)
	if err != nil {
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

func IfUserBoughtThisProduct(userID int, productID int) (bool, error) { //检查用户是否有这个订单
	query := "SELECT order_id FROM order_items WHERE product_id=?"
	rows, err := Db.Query(query, productID)
	if err != nil {
		return false, err
	}
	var orderID int
	var rows2 *sql.Rows
	var getUserID int
	for rows.Next() { //如果找到了对应的商品，开始遍历
		err = rows.Scan(&orderID) //获取订单id
		if err != nil {
			return false, err
		}
		query = "SELECT user_id FROM orders WHERE id=?" //在orders表单中获取对应订单
		rows2, err = Db.Query(query, orderID)
		if err != nil {
			return false, err
		}
		if rows2.Next() { //如果有对应订单（此处不采用遍历，因为结果只可能唯一）,一般都是有的，如果没有就是内部错误，写入订单时出问题了
			err = rows2.Scan(&getUserID) //获取下单人的userid
			if err != nil {
				return false, err
			}
			//接下来开始和本用户的id比较
			if getUserID == userID { //相等，即本用户就是下单人
				return true, nil
			} else {
				continue
			}
		} else {
			return false, respond.ErrOrderNotExists //属于内部错误，订单表单和订单商品表单不匹配
		}
	}
	//到这里，说明没找到商品或者用户
	return false, nil
}
