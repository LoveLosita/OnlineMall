package utils

import "OnlineMall/model"

func MapToSlice(inputMap map[model.ShowProduct]int) []model.MapToSlice {
	var resultList []model.MapToSlice
	for key, value := range inputMap {
		resultList = append(resultList, model.MapToSlice{Key: key, Value: value})
	}
	return resultList
}
