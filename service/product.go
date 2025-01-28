package service

import (
	"OnlineMall/auth"
	"OnlineMall/dao"
	"OnlineMall/model"
	"OnlineMall/respond"
	"OnlineMall/utils"
	"fmt"
	"math"
)

func AddProduct(product model.AddProduct, handlerID int) error {
	//1.检查用户权限
	role, err := auth.CheckPermission(handlerID) //检查用户权限
	if err != nil {                              //如果出错
		return err
	}
	if role != "merchant" && role != "admin" { //如果不是商家
		return respond.ErrUnauthorized //返回错误
	}
	//2.检查商品分类是否存在
	result, err := dao.CheckIfCategoryExists(product.CategoryID) //检查商品分类是否存在
	if err != nil {
		return err
	}
	if !result { //如果不存在
		return respond.ErrCategoryNotExists //返回错误
	}
	//3.检查商品信息是否合法
	maxNameLength := 80
	maxDescriptionLength := 10000 // typical max length for text type in MySQL
	maxPrice := 9999999999.99
	maxStock := math.MaxInt32
	if len(product.Name) > int(0.9*float64(maxNameLength)) ||
		len(product.Description) > int(0.9*float64(maxDescriptionLength)) ||
		product.Price > 0.9*maxPrice ||
		product.Stock > int(0.9*float64(maxStock)) {
		return respond.ParamTooLong
	}
	//4.保存商品信息
	return dao.AddProduct(product) //调用dao层函数
}

func ChangeProduct(id int, product model.AddProduct, handlerID int) error {
	//1.检查用户权限
	role, err := auth.CheckPermission(handlerID) //检查用户权限
	if err != nil {                              //如果出错
		return err
	}
	if role != "merchant" && role != "admin" { //如果不是商家及以上
		return respond.ErrUnauthorized //返回错误
	}
	//2.检查商品信息是否合法
	maxNameLength := 80
	maxDescriptionLength := 1000
	maxPrice := 9999999999.99
	maxStock := math.MaxInt32
	if len(product.Name) > int(0.9*float64(maxNameLength)) ||
		len(product.Description) > int(0.9*float64(maxDescriptionLength)) ||
		product.Price > 0.9*maxPrice ||
		product.Stock > int(0.9*float64(maxStock)) {
		return respond.ParamTooLong
	}
	//3.获取商品信息，并且判断是否已经填写对应信息，实现选择性更新
	oldProduct, err := dao.GetProductInfoByID(id, 0) //同时可以检查商品是否存在
	if err != nil {
		return err
	}
	if product.Name == "" { //如果没有填写商品名
		product.Name = oldProduct.Name //则使用原来的商品名
	}
	if product.Description == "" { //如果没有填写商品描述
		product.Description = oldProduct.Description //则使用原来的商品描述
	}
	if product.Price == 0 {
		product.Price = oldProduct.Price
	}
	if product.Stock == 0 {
		product.Stock = oldProduct.Stock
	}
	if product.CategoryID == 0 {
		product.CategoryID = oldProduct.CategoryID
	} else {
		result, err := dao.CheckIfCategoryExists(product.CategoryID) //检查商品分类是否存在
		if err != nil {
			return err
		}
		if !result { //如果不存在
			return respond.ErrCategoryNotExists //返回错误
		}
	}
	if product.ProductImage == "" {
		product.ProductImage = oldProduct.ProductImage
	}
	return dao.UpdateProduct(id, product) //调用dao层函数
}

func ShowProductInManyWays(productID int, keyword string, categoryID int) ([]model.ShowProduct, error) {
	//优先级：商品id>关键字>分类id>全部商品
	if productID != 0 { //如果有商品id
		product, err := dao.GetProductInfoByID(productID, 1)
		if err != nil {
			return nil, err
		}
		return []model.ShowProduct{product}, nil
	}
	if keyword != "" { //如果有关键字
		return dao.GetProductInfoByKeyWord(keyword)
	}
	if categoryID != 0 { //如果有分类id
		//检查分类是否存在
		result, err := dao.CheckIfCategoryExists(categoryID)
		if err != nil {
			return nil, err
		}
		if !result { //如果不存在
			return nil, respond.ErrCategoryNotExists //返回错误
		}
		return dao.ShowACategoryProducts(categoryID)
	}
	return dao.ShowAllProducts() //返回所有商品
}

func DeleteProduct(id int, handlerID int) error {
	//1.检查用户权限
	role, err := auth.CheckPermission(handlerID) //检查用户权限
	if err != nil {                              //如果出错
		return err
	}
	if role != "merchant" && role != "admin" { //如果不是商家及以上
		return respond.ErrUnauthorized //返回错误
	}
	//2.检查商品是否存在
	_, err = dao.GetProductInfoByID(id, 0) //检查商品是否存在
	if err != nil {
		return err
	}
	//3.删除商品
	return dao.DeleteProduct(id) //调用dao层函数
}

func AddUserProductHistory(userID, productID int) error {
	return dao.AddUserProductHistory(userID, productID) //调用dao层函数
}

func ShowUserProductHistory(userID int) ([]model.ShowProduct, error) {
	//1.检查用户是否存在
	_, err := dao.GetUserInfoByID(userID) //检查用户是否存在
	if err != nil {
		return nil, err
	}
	//2.获取用户浏览记录
	return dao.GetUserProductHistory(userID) //调用dao层函数
}

func SortProduct(getMethod int, userID int) ([]model.ShowProduct, error) {
	//先获取全部商品
	products, err := dao.ShowAllProducts()
	if err != nil {
		return nil, err
	}
	if getMethod == 1 { //1.游客访问，按照热度降序排序
		for i := 0; i < len(products); i++ {
			for j := 0; j < len(products)-1; j++ {
				if products[j].Popularity < products[j+1].Popularity {
					products[j], products[j+1] = products[j+1], products[j]
				}
			}
		}
		return products, nil
	} else if getMethod == 2 { //2.用户访问，按照常看程度和内容热度计算后，按照分数降序排序；剩下的没看过的商品再按照热度排序
		//2.1.获取用户历史记录，接下来以历史记录为基础，和所有商品的热度进行分数计算
		history, err := dao.GetUserProductHistory(userID)
		if err != nil {
			return nil, err
		}
		//2.2.调用函数计数历史记录
		resultMap := CountUserHistory(history)
		//2.3.转换成列表
		resultList := utils.MapToSlice(resultMap)
		//2.4.历史访问次数->排名
		historyRankList := utils.ListInRankOut(resultList)
		//2.5.产品热度->排名
		popRankList := utils.ProductInRankOut(products)
		//2.6.计算平均排名
		var scoreList []model.FloatMapToSlice
		var score float64
		for i := 0; i < len(historyRankList); i++ {
			for j := 0; j < len(popRankList); j++ {
				if popRankList[j].Key == historyRankList[i].Key {
					score = float64(popRankList[j].Value)*0.5 + float64(historyRankList[i].Value)*0.5
					scoreList = append(scoreList, model.FloatMapToSlice{Key: popRankList[j].Key, Value: score})
					fmt.Println(popRankList[j].Key, score)
				}
			}
		}
		//2.7.根据分数从低到高排序（分数越低的越好）
		for i := 0; i < len(scoreList); i++ {
			for j := 0; j < len(scoreList)-1; j++ {
				if scoreList[j].Value > scoreList[j+1].Value {
					scoreList[j], scoreList[j+1] = scoreList[j+1], scoreList[j]
				}
			}
		}
		//2.8.将最终顺序加入产品切片
		var finalProducts []model.ShowProduct
		for i := 0; i < len(scoreList); i++ {
			finalProducts = append(finalProducts, scoreList[i].Key)
		}
		//2.9.再把所有商品根据热度排序
		for i := 0; i < len(products); i++ {
			for j := 0; j < len(products)-1; j++ {
				if products[j].Popularity < products[j+1].Popularity {
					products[j], products[j+1] = products[j+1], products[j]
				}
			}
		}
		//2.10.再把不在上面的商品加入产品切片
		for i := 0; i < len(products); i++ {
			if !utils.InMapSlice(historyRankList, products[i]) {
				finalProducts = append(finalProducts, products[i])
			}
		}
		//2.11.返回最终结果
		return finalProducts, nil
	} else {
		return nil, fmt.Errorf("wrong get method(sv->product.go->sortProduct)")
	}
}

func CountUserHistory(history []model.ShowProduct) map[model.ShowProduct]int { //计数
	resultMap := make(map[model.ShowProduct]int)
	for i := 0; i < len(history); i++ {
		_, exists := resultMap[history[i]]
		if !exists {
			resultMap[history[i]] = 1
		} else {
			resultMap[history[i]]++
		}
	}
	return resultMap
}
