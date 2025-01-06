package service

import (
	"OnlineMall/dao"
	"OnlineMall/model"
)

func AddProduct(product model.Product) error {
	return dao.AddProduct(product)
}
