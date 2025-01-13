package service

import (
	"OnlineMall/dao"
	"OnlineMall/model"
	"OnlineMall/respond"
)

func RatingAndReviewProduct(review model.AddReview) error {
	//1.首先检查是否有这个商品
	_, err := dao.GetProductInfoByID(review.ProductID, 0)
	if err != nil {
		return err
	}
	//2.检查该用户是否购买过这个商品
	result, err := dao.IfUserBoughtThisProduct(review.UserID, review.ProductID)
	if err != nil {
		return err
	}
	if !result { //如果没有购买过
		return respond.ErrUserDidntBuyThisProduct
	}
	//3.检查是否已经评论过
	result, err = dao.IfUserHasReviewedThisProduct(review.UserID, review.ProductID)
	if err != nil {
		return err
	}
	if result { //如果已经评论过
		return respond.ErrUserHasAlreadyReviewed
	}
	//4.检查参数是否合法
	if review.Rating == -1 && review.Comment == "" { //如果评分和评论都为空
		return respond.MissingParam
	}
	if review.Rating < 1 || review.Rating > 5 { //如果没有被标记为-1，那么评分必须在1-5之间
		return respond.ErrRatingOutOfRange
	}
	if len(review.Comment) > 1000 { //评论字数限制1000
		return respond.ErrCommentTooLong
	}
	//5.看是否要自动评论
	if review.Comment == "" { //如果评论为空
		if review.Rating == 5 {
			review.Comment = "该用户觉得该商品很好"
		} else if review.Rating >= 3 {
			review.Comment = "该用户觉得该商品还行"
		} else if review.Rating >= 2 {
			review.Comment = "该用户觉得该商品一般"
		} else if review.Rating >= 1 {
			review.Comment = "该用户觉得该商品较差"
		}
	}
	//6.评论
	err = dao.RateAndReviewProduct(review)
	if err != nil {
		return err
	}
	//7.更新商品的评分
	err = UpdateAverageRating(review.ProductID)
	if err != nil {
		return err
	}
	return nil
}

func UpdateAverageRating(productID int) error {
	//1.获取所有评分
	ratings, err := dao.GetProductRatings(productID)
	if err != nil {
		return err
	}
	//2.计算平均分
	var sum int
	for _, rating := range ratings {
		sum += rating
	}
	var aveRating float64
	aveRating = float64(sum) / float64(len(ratings))
	//3.更新商品的平均分
	err = dao.UpdateRating(productID, aveRating)
	if err != nil {
		return err
	}
	return nil
}
