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
		case errors.Is(err, respond.InvalidName), errors.Is(err, respond.MissingParam),
			errors.Is(err, respond.ParamTooLong), errors.Is(err, respond.WrongGender): //如果是无效ID或者缺少参数的错误
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
	result, tokens, err := service.UserLogin(user)
	if err != nil {
		switch {
		case errors.Is(err, respond.WrongName), errors.Is(err, respond.WrongPwd),
			errors.Is(err, respond.WrongGender): //如果是无效ID或者密码错误或者性别错误
			c.JSON(consts.StatusBadRequest, respond.WrongUsernameOrPwd)
			return
		default:
			c.JSON(consts.StatusInternalServerError, respond.InternalError(err))
			return
		}
	}
	if result { //登录成功
		c.JSON(consts.StatusOK, respond.Respond(respond.Ok, tokens))
	} else { //密码错误
		c.JSON(consts.StatusBadRequest, respond.WrongPwd)
	}
}

func ChangeUserPasswordOrName(ctx context.Context, c *app.RequestContext) {
	handlerID := c.GetFloat64("user_id")
	var user model.ChangePasswordAndUsernameUser
	err := c.BindJSON(&user)
	if err != nil {
		c.JSON(consts.StatusBadRequest, respond.WrongParamType)
		return
	}
	err = service.ChangeUserPwdOrName(int(handlerID), user)
	if err != nil {
		switch {
		case errors.Is(err, respond.WrongPwd), errors.Is(err, respond.WrongName), errors.Is(err, respond.MissingParam): //如果是密码错误或者用户名错误或者缺少参数的错误
			c.JSON(consts.StatusUnauthorized, err)
			return
		default:
			c.JSON(consts.StatusInternalServerError, respond.InternalError(err))
			return
		}
	}
	c.JSON(consts.StatusOK, respond.Ok)
}

func ChangeUserInfo(ctx context.Context, c *app.RequestContext) {
	//1.从请求中获取handlerID和targetID
	handlerID := c.GetFloat64("user_id")
	targetID := c.Query("id")
	intID, err := strconv.ParseInt(targetID, 10, 64)
	var user model.ChangeInfoUser
	err = c.BindJSON(&user)
	if err != nil {
		c.JSON(consts.StatusBadRequest, respond.WrongParamType)
		return
	}
	//2.调用service层函数
	err = service.ChangeUserInfo(int(handlerID), int(intID), user)
	if err != nil {
		switch {
		case errors.Is(err, respond.WrongGender), errors.Is(err, respond.WrongName), errors.Is(err, respond.MissingParam),
			errors.Is(err, respond.ParamTooLong), errors.Is(err, respond.ErrUnauthorized): //如果是性别错误或者用户名错误或者缺少参数或者参数过长的错误
			c.JSON(consts.StatusBadRequest, err)
			return
		default:
			c.JSON(consts.StatusInternalServerError, respond.InternalError(err))
			return
		}
	}
	//3.返回成功
	c.JSON(consts.StatusOK, respond.Ok)
}

func DeleteUser(ctx context.Context, c *app.RequestContext) {
	handlerID := c.GetFloat64("user_id")
	targetID := c.Query("id")
	intID, err := strconv.ParseInt(targetID, 10, 64)
	if err != nil {
		c.JSON(consts.StatusBadRequest, respond.WrongParamType)
		return
	}
	err = service.DeleteUser(int(handlerID), int(intID))
	if err != nil {
		switch {
		case errors.Is(err, respond.WrongUserID), errors.Is(err, respond.ErrUnauthorized): //如果是用户id错误或者权限错误
			c.JSON(consts.StatusBadRequest, err)
			return
		default:
			c.JSON(consts.StatusInternalServerError, respond.InternalError(err))
			return
		}
	}
	c.JSON(consts.StatusOK, respond.Ok)
}
