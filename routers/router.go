package routers

import (
	"OnlineMall/api"
	"github.com/cloudwego/hertz/pkg/app/server"
)

func RegisterRouters() {
	h := server.Default()
	userGroup := h.Group("/user")

	userGroup.PUT("/register", api.UserRegister)

	h.Spin()
}
