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

func GetProductInfo(productID int) (model.ShowProduct, error) {
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
