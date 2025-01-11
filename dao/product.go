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

func GetProductInfoByID(productID int) (model.ShowProduct, error) {
	query := "SELECT * FROM products WHERE id=?"
	rows, err := Db.Query(query, productID)
	if err != nil {
		return model.ShowProduct{}, err
	}
	var product model.ShowProduct
	for rows.Next() { //如果有这个商品
		err = rows.Scan(&product.ID, &product.Name, &product.Description, &product.Price, &product.Stock, &product.CategoryID,
			&product.ProductImage, &product.Popularity, &product.CreatedAt, &product.UpdatedAt)
		if err != nil {
			return model.ShowProduct{}, err
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
			&product.ProductImage, &product.Popularity, &product.CreatedAt, &product.UpdatedAt)
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

func DeleteProduct(id int) error {
	query := "DELETE FROM products WHERE id=?"
	_, err := Db.Exec(query, id)
	if err != nil {
		return err
	}
	return nil
}
