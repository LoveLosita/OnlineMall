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

func CheckIfCategoryExists(categoryID int) (bool, error) {
	query := "SELECT id FROM categories WHERE id=?"
	rows, err := Db.Query(query, categoryID)
	if err != nil {
		return false, err
	}
	if rows.Next() { //如果有这个分类
		return true, nil
	} else { //如果没有这个分类
		return false, nil
	}
}
