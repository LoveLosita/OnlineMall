package api

import (
	"OnlineMall/model"
	"OnlineMall/respond"
	"OnlineMall/service"
	"context"
	"errors"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"strconv"
)

func RateAndReviewProduct(ctx context.Context, c *app.RequestContext) {
	handlerID := int(c.GetFloat64("user_id")) //从上下文中获取用户的id
	// 从请求中获取评论信息
	var review model.AddReview
	err := c.BindJSON(&review)
	if err != nil {
		c.JSON(consts.StatusBadRequest, respond.WrongParamType)
		return
	}
	// 评论商品
	review.UserID = handlerID
	err = service.RatingAndReviewProduct(review)
	if err != nil {
		switch {
		case errors.Is(err, respond.ErrProductNotExists), errors.Is(err, respond.ErrUserDidntBuyThisProduct),
			errors.Is(err, respond.ErrUserHasAlreadyReviewed), errors.Is(err, respond.ErrRatingOutOfRange),
			errors.Is(err, respond.ErrCommentTooLong), errors.Is(err, respond.MissingParam):
			c.JSON(consts.StatusUnauthorized, err)
			return
		case errors.Is(err, respond.ErrOrderNotExists):
			c.JSON(consts.StatusInternalServerError, err)
		default:
			c.JSON(consts.StatusInternalServerError, respond.InternalError(err))
			return
		}
	}
	c.JSON(consts.StatusOK, respond.Ok)
}

func ReplyToReview(ctx context.Context, c *app.RequestContext) {
	handlerID := int(c.GetFloat64("user_id")) //从上下文中获取用户的id
	// 从请求中获取评论信息
	var review model.ReplyToReview
	err := c.BindJSON(&review)
	if err != nil {
		c.JSON(consts.StatusBadRequest, respond.WrongParamType)
		return
	}
	// 回复评论
	err = service.ReplyToReview(handlerID, review)
	if err != nil {
		switch {
		case errors.Is(err, respond.MissingParam), errors.Is(err, respond.ErrCommentTooLong),
			errors.Is(err, respond.ErrParentNotExists):
			c.JSON(consts.StatusBadRequest, err)
			return
		default:
			c.JSON(consts.StatusInternalServerError, respond.InternalError(err))
			return
		}
	}
	c.JSON(consts.StatusOK, respond.Ok)
}

func ShowAProductReviews(ctx context.Context, c *app.RequestContext) {
	//1.从请求中获取商品id
	productID := c.Query("product_id")
	intProductID, err := strconv.ParseInt(productID, 10, 64)
	if err != nil {
		c.JSON(consts.StatusBadRequest, respond.WrongParamType)
		return
	}
	//2.获取商品的评论
	reviews, err := service.BuildReviewTree(int(intProductID))
	if err != nil {
		switch {
		case errors.Is(err, respond.EmptyProductReviews), errors.Is(err, respond.ErrProductNotExists):
			c.JSON(consts.StatusNotFound, err)
			return
		default:
			c.JSON(consts.StatusInternalServerError, respond.InternalError(err))
			return
		}
	}
	c.JSON(consts.StatusOK, respond.Respond(respond.Ok, reviews))
}

func SearchForAProductReview(ctx context.Context, c *app.RequestContext) {
	//1.从请求中获取商品id
	productID := c.Query("product_id")
	intProductID, err := strconv.ParseInt(productID, 10, 64)
	if err != nil {
		c.JSON(consts.StatusBadRequest, respond.WrongParamType)
		return
	}
	//2.从请求中获取关键词
	keyword := c.Query("keyword")
	if keyword == "" {
		c.JSON(consts.StatusBadRequest, respond.MissingParam)
		return
	}
	//3.从上下文中获取用户id
	handlerID := int(c.GetFloat64("user_id"))
	//4.获取商品的评论
	review, err := service.BuildReviewTree2(int(intProductID), keyword, handlerID)
	if err != nil {
		switch {
		case errors.Is(err, respond.ErrUnauthorized), errors.Is(err, respond.CantFindReview):
			c.JSON(consts.StatusNotFound, err)
			return
		default:
			c.JSON(consts.StatusInternalServerError, respond.InternalError(err))
			return
		}
	}
	c.JSON(consts.StatusOK, respond.Respond(respond.Ok, review))
}

func DeleteReview(ctx context.Context, c *app.RequestContext) {
	//1.从请求中获取评论id
	reviewID := c.Query("id")
	intReviewID, err := strconv.ParseInt(reviewID, 10, 64)
	if err != nil {
		c.JSON(consts.StatusBadRequest, respond.WrongParamType)
		return
	}
	//2.从上下文中获取用户id
	handlerID := int(c.GetFloat64("user_id"))
	//3.删除评论
	err = service.DeleteReview(handlerID, int(intReviewID))
	if err != nil {
		switch {
		case errors.Is(err, respond.ErrUnauthorized), errors.Is(err, respond.ErrReviewNotExists):
			c.JSON(consts.StatusUnauthorized, err)
			return
		default:
			c.JSON(consts.StatusInternalServerError, respond.InternalError(err))
			return
		}
	}
	c.JSON(consts.StatusOK, respond.Ok)
}
