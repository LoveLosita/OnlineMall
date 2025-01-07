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

func AddProduct(ctx context.Context, c *app.RequestContext) {
	handlerID := int(c.GetFloat64("user_id")) //从上下文中获取用户的id
	// 从请求中获取商品信息
	product := model.Product{}
	err := c.BindJSON(&product)
	if err != nil {
		c.JSON(consts.StatusBadRequest, respond.WrongParamType)
		return
	}
	// 保存商品信息
	err = service.AddProduct(product, handlerID)
	if err != nil {
		switch {
		case errors.Is(err, respond.ErrUnauthorized), errors.Is(err, respond.ErrCategoryNotExists):
			c.JSON(consts.StatusUnauthorized, err)
			return
		default:
			c.JSON(consts.StatusInternalServerError, respond.InternalError(err))
			return
		}
	}
	c.JSON(consts.StatusOK, respond.Ok)
}
