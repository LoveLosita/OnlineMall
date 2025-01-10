package api

import (
	"OnlineMall/respond"
	"OnlineMall/service"
	"context"
	"errors"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

func RefreshTokenHandler(ctx context.Context, c *app.RequestContext) {
	var requestBody struct {
		RefreshToken string `json:"refresh_token"`
	}
	if err := c.Bind(&requestBody); err != nil {
		c.JSON(consts.StatusBadRequest, respond.WrongParamType)
		return
	}
	tokens, err := service.RefreshTokenHandler(requestBody.RefreshToken)
	if err != nil {
		switch {
		case errors.Is(err, respond.InvalidRefreshToken), errors.Is(err, respond.InvalidClaims),
			errors.Is(err, respond.InvalidTokenSingingMethod): //如果是无效刷新令牌或者无效claims或者无效签名方法
			c.JSON(consts.StatusBadRequest, err)
			return
		default:
			c.JSON(consts.StatusInternalServerError, respond.InternalError(err))
		}
	}
	c.JSON(consts.StatusOK, respond.Respond(respond.Ok, tokens))
}
