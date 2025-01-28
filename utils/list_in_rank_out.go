package utils

import "OnlineMall/model"

func ListInRankOut(inputList []model.MapToSlice) []model.MapToSlice { //输入键和次数，输出键和排名
	var result []model.MapToSlice
	var count int
	for i := 0; i < len(inputList); i++ {
		count = 1
		for j := 0; j < len(inputList); j++ {
			if inputList[j].Value > inputList[i].Value && i != j {
				count++
			}
		}
		result = append(result, model.MapToSlice{Key: inputList[i].Key, Value: count})
	}
	return result
}

func ProductInRankOut(products []model.ShowProduct) []model.MapToSlice {
	var result []model.MapToSlice
	var count int
	for i := 0; i < len(products); i++ {
		count = 1
		for j := 0; j < len(products); j++ {
			if products[j].Popularity > products[i].Popularity && i != j {
				count++
			}
		}
		result = append(result, model.MapToSlice{Key: products[i], Value: count})
	}
	return result
}
