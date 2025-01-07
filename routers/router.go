package routers

import (
	"OnlineMall/api"
	"OnlineMall/middleware"
	"github.com/cloudwego/hertz/pkg/app/server"
)

func RegisterRouters() {
	h := server.Default()
	userGroup := h.Group("/user")
	merchantGroup := h.Group("/merchant")

	userGroup.PUT("/register", api.UserRegister)
	userGroup.POST("/login", api.UserLogin)
	userGroup.POST("/change_username_or_password", middleware.JWTAuthMiddleware(), api.ChangeUserPasswordOrName)

	merchantGroup.PUT("/add_product", middleware.JWTAuthMiddleware(), api.AddProduct)

	h.Spin()
}
