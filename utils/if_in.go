package utils

import "OnlineMall/model"

func InIntSlice(slice []int, target int) bool {
	for _, v := range slice {
		if v == target {
			return true
		}
	}
	return false
}

func InMapSlice(slice []model.MapToSlice, target model.ShowProduct) bool {
	for _, value := range slice {
		if value.Key == target {
			return true
		}
	}
	return false
}
