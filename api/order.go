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

func PlaceOrder(ctx context.Context, c *app.RequestContext) {
	// 1.从请求中获取订单信息和用户信息
	handlerID := int(c.GetFloat64("user_id")) //从上下文中获取用户的id
	var order model.PlaceOrder
	err := c.BindJSON(&order)
	if err != nil {
		c.JSON(consts.StatusBadRequest, respond.WrongParamType)
		return
	}
	// 2.保存订单信息
	order.UserID = handlerID
	returnOrder, err := service.PlaceAnOrder(order)
	if err != nil {
		switch {
		case errors.Is(err, respond.ErrProductNotExists), errors.Is(err, respond.ErrQuantityTooLarge):
			c.JSON(consts.StatusBadRequest, err)
			return
		default:
			c.JSON(consts.StatusInternalServerError, respond.InternalError(err))
			return
		}
	}
	// 3.返回结果
	c.JSON(consts.StatusOK, respond.Respond(respond.Ok, returnOrder))
}
