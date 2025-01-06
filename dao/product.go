package dao

import "OnlineMall/model"

func AddProduct(product model.Product) error {
	query := "INSERT INTO products(name,description,price,stock,category_id,product_image) VALUES(?,?,?,?,?,?)"
	_, err := Db.Exec(query, product.Name, product.Description, product.Price, product.Stock, product.CategoryID, product.ProductImage)
	if err != nil {
		return err
	}
	return nil
}
