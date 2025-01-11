package api

import (
	"OnlineMall/model"
	"OnlineMall/respond"
	"OnlineMall/service"
	"context"
	"errors"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

func AddProductToCart(ctx context.Context, c *app.RequestContext) {
	//1.获取请求体
	var requestBody model.AddToCart
	handlerID := int(c.GetFloat64("user_id"))
	if err := c.Bind(&requestBody); err != nil {
		c.JSON(consts.StatusBadRequest, respond.WrongParamType)
		return
	}
	//2.添加商品到购物车
	err := service.AddProductToCart(handlerID, requestBody.ProductID, requestBody.Quantity)
	if err != nil {
		switch {
		case errors.Is(err, respond.ErrProductAlreadyInCart), errors.Is(err, respond.ErrProductNotExists):
			c.JSON(consts.StatusBadRequest, err)
			return
		default:
			c.JSON(consts.StatusInternalServerError, respond.InternalError(err))
			return
		}
	}
	c.JSON(consts.StatusOK, respond.Ok)
}
