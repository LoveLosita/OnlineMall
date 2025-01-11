package dao

import "OnlineMall/model"

func AddToCart(userID, productID, quantity int) error { //添加商品到购物车
	query := "INSERT INTO carts(user_id,product_id,quantity) VALUES(?,?,?)"
	_, err := Db.Exec(query, userID, productID, quantity)
	if err != nil {
		return err
	}
	return nil
}

func IfProductExistsInYourCart(productID int, userID int) (bool, error) { //检查商品是否已经在购物车中
	query := "SELECT id FROM carts WHERE product_id=? AND user_id=?"
	rows, err := Db.Query(query, productID, userID)
	if err != nil {
		return false, err
	}
	if rows.Next() { //如果有这个商品
		return true, nil
	} else { //如果没有这个商品
		return false, nil
	}
}

func GetUserCart(userID int) ([]model.ProductInCart, error) { //获取用户购物车全部商品
	query := "SELECT * FROM carts WHERE user_id=?"
	rows, err := Db.Query(query, userID)
	if err != nil {
		return nil, err
	}
	var products []model.ProductInCart
	for rows.Next() {
		var product model.ProductInCart
		err = rows.Scan(&product.ID, &product.UserID, &product.ProductID, &product.Quantity)
		if err != nil {
			return nil, err
		}
		products = append(products, product)
	}
	return products, nil
}

func GetSingleProductInCart(userID, productID int) (model.ProductInCart, error) { //获取购物车中单个商品
	query := "SELECT * FROM carts WHERE user_id=? AND product_id=?"
	rows, err := Db.Query(query, userID, productID)
	if err != nil {
		return model.ProductInCart{}, err
	}
	var product model.ProductInCart
	for rows.Next() {
		err = rows.Scan(&product.ID, &product.UserID, &product.ProductID, &product.Quantity, &product.CreatedAt, &product.UpdatedAt)
		if err != nil {
			return model.ProductInCart{}, err
		}
		return product, nil
	}
	return model.ProductInCart{}, nil
}

func UpdateProductQuantityInCart(userID, productID, quantity int) error { //更新购物车中商品数量
	query := "UPDATE carts SET quantity=? WHERE user_id=? AND product_id=?"
	_, err := Db.Exec(query, quantity, userID, productID)
	if err != nil {
		return err
	}
	return nil
}
