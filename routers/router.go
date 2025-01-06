package routers

import (
	"OnlineMall/api"
	"github.com/cloudwego/hertz/pkg/app/server"
)

func RegisterRouters() {
	h := server.Default()
	userGroup := h.Group("/user")
	//merchantGroup := h.Group("/merchant")

	userGroup.PUT("/register", api.UserRegister)
	userGroup.POST("/login", api.UserLogin)

	//merchantGroup.PUT("/add_product", api.MerchantAddProduct)

	h.Spin()
}
