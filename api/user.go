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
		case errors.Is(err, respond.InvalidName), errors.Is(err, respond.MissingParam), errors.Is(err, respond.ParamTooLong): //如果是无效ID或者缺少参数的错误
			c.JSON(consts.StatusBadRequest, err)
			return
		default:
			c.JSON(consts.StatusInternalServerError, respond.InternalError(err))
			return
		}
	}
	c.JSON(consts.StatusOK, respond.Ok)
}

func UserLogin(ctx context.Context, c *app.RequestContext) {
	var user model.LoginUser
	err := c.BindJSON(&user)
	if err != nil {
		c.JSON(consts.StatusBadRequest, respond.WrongParamType)
		return
	}
	result, jwtKey, err := service.UserLogin(user)
	if err != nil {
		switch {
		case errors.Is(err, respond.WrongName), errors.Is(err, respond.WrongPwd), errors.Is(err, respond.WrongGender): //如果是无效ID或者密码错误或者性别错误
			c.JSON(consts.StatusBadRequest, respond.WrongUsernameOrPwd)
			return
		default:
			c.JSON(consts.StatusInternalServerError, respond.InternalError(err))
			return
		}
	}
	var jwt model.JWTKey
	jwt.Key = jwtKey
	if result { //登录成功
		c.JSON(consts.StatusOK, respond.Respond(respond.Ok, jwt))
	} else { //密码错误
		c.JSON(consts.StatusBadRequest, respond.WrongPwd)
	}
}
