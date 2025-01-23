package service

import (
	"OnlineMall/auth"
	"OnlineMall/dao"
	"OnlineMall/model"
	"OnlineMall/respond"
	"OnlineMall/utils"
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
	err = dao.AddReviewOrReply(review)
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

func ReplyToReview(handlerID int, reply model.ReplyToReview) error {
	//1.检查参数合法性
	if reply.ReplyToID == 0 || reply.Reply == "" {
		return respond.MissingParam
	}
	if len(reply.Reply) > 1000 {
		return respond.ErrCommentTooLong
	}
	//2.检查parent_id是否存在
	result, err := dao.IfReviewExists(reply.ReplyToID)
	if err != nil {
		return err
	}
	if !result { //如果不存在
		return respond.ErrParentNotExists
	}
	//3.将回复结构体信息转写进review结构体
	var review model.AddReview
	review.UserID = handlerID
	review.ProductID = -1 //因为是回复，所以不需要product_id
	review.ParentID = &reply.ReplyToID
	review.Rating = -1 //因为是回复，所以不需要rating
	review.Comment = reply.Reply
	//4.调用dao函数完成回复
	err = dao.AddReviewOrReply(review)
	if err != nil {
		return err
	}
	return nil
}

func BuildReviewTree(productID int) ([]model.ShowReview, error) { //构建评论树
	//1.首先检查是否有这个商品
	_, err := dao.GetProductInfoByID(productID, 0)
	if err != nil {
		return nil, err
	}
	//2.获取所有评论
	reviews, err := dao.GetAProductReviews(productID) //获取所有评论
	if len(reviews) == 0 {                            //如果评论为空
		return nil, respond.EmptyProductReviews
	}
	if err != nil {
		return nil, err
	}
	//3.构建评论树
	for i := len(reviews) - 1; i >= 0; i-- { //从后往前遍历
		if reviews[i].ParentID != nil { //如果有父评论
			for j := 0; j < i; j++ { //寻找父评论
				if *reviews[i].ParentID == reviews[j].ID { //如果找到了父评论
					reviews[j].Replies = append(reviews[j].Replies, reviews[i]) //将子评论添加到父评论的Replies中
					break
				}
			}
		}
	}
	var resultList []model.ShowReview //定义一个评论列表
	for _, review := range reviews {  //遍历评论
		if review.ParentID == nil { //如果没有父评论
			resultList = append(resultList, review) //将评论添加到评论列表中
		}
	}
	return resultList, nil
}

func BuildReviewTree2(productID int, keyword string, handlerID int) ([]model.ShowReview, error) { //构建评论树
	//1.首先检查权限
	result, err := auth.CheckPermission(handlerID)
	if err != nil {
		return nil, err
	}
	if result != "merchant" && result != "admin" { //如果不是商家或管理员
		return nil, respond.ErrUnauthorized
	}
	//2.检查是否有这个商品
	_, err = dao.GetProductInfoByID(productID, 0)
	if err != nil {
		return nil, err
	}
	//3.获取所有评论
	reviews, err := dao.SearchForProductReviews(productID, keyword) //搜索评论
	if len(reviews) == 0 {                                          //如果评论为空
		return nil, respond.CantFindReview
	}
	if err != nil {
		return nil, err
	}
	//4.构建评论树
	var appendSlice []int
	for i := len(reviews) - 1; i >= 0; i-- { //从后往前遍历
		if reviews[i].ParentID != nil { //如果有父评论
			for j := 0; j < i; j++ { //寻找父评论
				if *reviews[i].ParentID == reviews[j].ID { //如果找到了父评论
					reviews[j].Replies = append(reviews[j].Replies, reviews[i]) //将子评论添加到父评论的Replies中
					appendSlice = append(appendSlice, reviews[i].ID)            //将子评论的ID添加到appendSlice中
					break
				}
			}
		}
	}
	var resultList []model.ShowReview //定义一个评论列表
	for _, review := range reviews {  //遍历评论
		if review.ParentID == nil || !utils.InIntSlice(appendSlice, review.ID) { //如果没有父评论或者ID不在appendSlice中
			resultList = append(resultList, review) //将评论添加到评论列表中
		}
	}
	return resultList, nil
}

func DeleteReview(handlerID int, reviewID int) error {
	//1.检查权限
	role, err := auth.CheckPermission(handlerID)
	if err != nil {
		return err
	}
	if role != "admin" { //如果不是管理员
		return respond.ErrUnauthorized
	}
	//2.检查评论是否存在
	result, err := dao.IfReviewExists(reviewID)
	if err != nil {
		return err
	}
	if !result { //如果不存在
		return respond.ErrReviewNotExists
	}
	//获取商品id
	productID, err := dao.GetProductIDByReviewID(reviewID)
	if err != nil {
		return err
	}
	//3.删除评论
	err = dao.DeleteReview(reviewID)
	if err != nil {
		return err
	}
	//4.删除评论后更新商品的评分
	err = UpdateAverageRating(productID)
	if err != nil {
		return err
	}
	return nil
}
