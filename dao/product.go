package dao

import (
	"OnlineMall/model"
	"OnlineMall/respond"
)

func AddProduct(product model.AddProduct) error {
	query := "INSERT INTO products(name,description,price,stock,category_id,product_image) VALUES(?,?,?,?,?,?)"
	_, err := Db.Exec(query, product.Name, product.Description, product.Price, product.Stock, product.CategoryID, product.ProductImage)
	if err != nil {
		return err
	}
	return nil
}

func UpdateProduct(id int, product model.AddProduct) error {
	query := "UPDATE products SET name=?,description=?,price=?,stock=?,category_id=?,product_image=? WHERE id=?"
	_, err := Db.Exec(query, product.Name, product.Description, product.Price, product.Stock, product.CategoryID, product.ProductImage, id)
	if err != nil {
		return err
	}
	return nil
}

func GetProductInfoByID(productID int, getMethod int) (model.ShowProduct, error) { //计入热度计算
	query := "SELECT * FROM products WHERE id=?"
	rows, err := Db.Query(query, productID)
	if err != nil {
		return model.ShowProduct{}, err
	}
	var product model.ShowProduct
	for rows.Next() { //如果有这个商品
		err = rows.Scan(&product.ID, &product.Name, &product.Description, &product.Price, &product.Stock, &product.CategoryID,
			&product.Popularity, &product.AveRating, &product.ProductImage, &product.CreatedAt, &product.UpdatedAt)
		if err != nil {
			return model.ShowProduct{}, err
		}
		//如果getMethod==1，说明是通过点击商品详情页进入的，需要增加热度；其他都是后端计算调用，不需要增加热度
		if getMethod == 1 {
			query = "UPDATE products SET popularity=popularity+1 WHERE id=?"
			_, err = Db.Exec(query, productID)
			if err != nil {
				return model.ShowProduct{}, err
			}
		}
		return product, nil
	}
	return model.ShowProduct{}, respond.ErrProductNotExists
}

func GetProductInfoByKeyWord(keyword string) ([]model.ShowProduct, error) {
	query := "SELECT * FROM products WHERE name LIKE ?"
	rows, err := Db.Query(query, "%"+keyword+"%")
	if err != nil {
		return nil, err
	}
	var products []model.ShowProduct
	for rows.Next() {
		var product model.ShowProduct
		err = rows.Scan(&product.ID, &product.Name, &product.Description, &product.Price, &product.Stock, &product.CategoryID,
			&product.Popularity, &product.AveRating, &product.ProductImage, &product.CreatedAt, &product.UpdatedAt)
		if err != nil {
			return nil, err
		}
		products = append(products, product)
	}
	if len(products) == 0 { //如果没有找到商品
		return nil, respond.CantFindProduct
	}
	return products, nil
}

func ShowAllProducts() ([]model.ShowProduct, error) {
	query := "SELECT * FROM products"
	rows, err := Db.Query(query)
	if err != nil {
		return nil, err
	}
	var products []model.ShowProduct
	for rows.Next() {
		var product model.ShowProduct
		err = rows.Scan(&product.ID, &product.Name, &product.Description, &product.Price, &product.Stock, &product.CategoryID,
			&product.ProductImage, &product.Popularity, &product.CreatedAt, &product.UpdatedAt)
		if err != nil {
			return nil, err
		}
		products = append(products, product)
	}
	if len(products) == 0 { //如果没有商品
		return nil, respond.EmptyProductList
	}
	return products, nil
}

func ShowACategoryProducts(categoryID int) ([]model.ShowProduct, error) {
	query := "SELECT * FROM products WHERE category_id=?"
	rows, err := Db.Query(query, categoryID)
	if err != nil {
		return nil, err
	}
	var products []model.ShowProduct
	for rows.Next() {
		var product model.ShowProduct
		err = rows.Scan(&product.ID, &product.Name, &product.Description, &product.Price, &product.Stock, &product.CategoryID,
			&product.Popularity, &product.AveRating, &product.ProductImage, &product.CreatedAt, &product.UpdatedAt)
		if err != nil {
			return nil, err
		}
		products = append(products, product)
	}
	if len(products) == 0 { //如果没有商品
		return nil, respond.EmptyProductList
	}
	return products, nil
}

func DeleteProduct(id int) error {
	query := "DELETE FROM products WHERE id=?"
	_, err := Db.Exec(query, id)
	if err != nil {
		return err
	}
	return nil
}

func UpdateRating(productID int, aveRating float64) error {
	query := "UPDATE products SET ave_rating=? WHERE id=?"
	_, err := Db.Exec(query, aveRating, productID)
	if err != nil {
		return err
	}
	return nil
}
