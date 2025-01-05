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

func UserRegister(ctx context.Context, c *app.RequestContext) {
	var user model.User
	err := c.BindJSON(&user)
	if err != nil {
		c.JSON(consts.StatusBadRequest, respond.WrongParamType)
		return
	}
	err = service.UserRegister(user)
	if err != nil {
		switch {
		case errors.Is(err, respond.InvalidID), errors.Is(err, respond.MissingParam): //如果是无效ID或者缺少参数的错误
			c.JSON(consts.StatusBadRequest, respond.WrongParamType)
			return
		default:
			c.JSON(consts.StatusInternalServerError, respond.InternalError(err))
			return
		}
	}
	c.JSON(consts.StatusOK, respond.Ok)
}
