package middleware

import (
	"OnlineMall/auth"
	"OnlineMall/respond"
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/golang-jwt/jwt/v4"
)

var jwtSecret = auth.JwtSecret

// JWTTokenAuth 中间件
func JWTTokenAuth() app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		// 从请求头中获取 Authorization 字段
		tokenString := c.GetHeader("Authorization")
		//fmt.Println(string(tokenString))//测试用
		if string(tokenString) == "" { //没有token
			c.JSON(consts.StatusUnauthorized, respond.MissingToken)
			c.Abort() // 中断后续流程
			return
		}

		// 解析并验证 Token
		token, err := jwt.Parse(string(tokenString), func(token *jwt.Token) (interface{}, error) {
			// 确保签名方法是我们支持的 HMAC
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, respond.InvalidTokenSingingMethod
			}
			return jwtSecret, nil
		})
		if err != nil || !token.Valid { //token无效
			c.JSON(consts.StatusUnauthorized, respond.InvalidToken)
			c.Abort() // 中断后续流程
			return
		}

		// 将解析出的用户信息存入上下文，供后续使用
		if claims, ok := token.Claims.(jwt.MapClaims); ok {
			c.Set("user_id", claims["user_id"])
			//fmt.Printf("%T", claims["user_id"]) //测试用
			//fmt.Println("Claims:", claims) // 打印出所有解析到的 claims//测试用
		} else {
			c.JSON(consts.StatusUnauthorized, respond.InvalidClaims)
			c.Abort()
		}
	}
}

func JWTTokenAuthForProductHistory() app.HandlerFunc {
	//这个中间件是用于游客和用户都能访问的接口的,如果登录了，就会把用户id放入上下文，用以记录用户的浏览记录
	return func(ctx context.Context, c *app.RequestContext) {
		// 从请求头中获取 Authorization 字段
		tokenString := c.GetHeader("Authorization")
		//fmt.Println(string(tokenString))//测试用
		if string(tokenString) == "" { //没有token
			return //不阻止后续流程，因为这个中间件是用于游客和用户都能访问的接口的
		}

		// 解析并验证 Token
		token, err := jwt.Parse(string(tokenString), func(token *jwt.Token) (interface{}, error) {
			// 确保签名方法是我们支持的 HMAC
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, respond.InvalidTokenSingingMethod
			}
			return jwtSecret, nil
		})
		if err != nil || !token.Valid { //token无效
			c.JSON(consts.StatusUnauthorized, respond.InvalidToken)
			c.Abort() // 中断后续流程
			return
		}

		// 将解析出的用户信息存入上下文，供后续使用
		if claims, ok := token.Claims.(jwt.MapClaims); ok {
			c.Set("user_id", claims["user_id"])
			//fmt.Printf("%T", claims["user_id"]) //测试用
			//fmt.Println("Claims:", claims) // 打印出所有解析到的 claims//测试用
		} else {
			c.JSON(consts.StatusUnauthorized, respond.InvalidClaims)
			c.Abort()
		}
	}
}
