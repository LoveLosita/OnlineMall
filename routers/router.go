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
	adminGroup := h.Group("/admin")

	//分组依据为使用对应功能需要的最低权限
	h.GET("show_product", api.GetProductsInManyWays)

	userGroup.PUT("/register", api.UserRegister)
	userGroup.POST("/login", api.UserLogin)
	userGroup.POST("/change_username_or_password", middleware.JWTAuthMiddleware(), api.ChangeUserPasswordOrName)
	userGroup.GET("/refresh_token", middleware.JWTAuthMiddleware(), api.RefreshTokenHandler)
	userGroup.POST("/add_to_cart", middleware.JWTAuthMiddleware(), api.AddProductToCart)
	userGroup.POST("/place_order", middleware.JWTAuthMiddleware(), api.PlaceOrder)
	userGroup.POST("/rate_and_review", middleware.JWTAuthMiddleware(), api.RateAndReviewProduct)

	merchantGroup.PUT("/add_product", middleware.JWTAuthMiddleware(), api.AddProduct)
	merchantGroup.PUT("/add_category", middleware.JWTAuthMiddleware(), api.AddCategory)
	merchantGroup.POST("/change_product", middleware.JWTAuthMiddleware(), api.ChangeProduct)
	merchantGroup.DELETE("/delete_product", middleware.JWTAuthMiddleware(), api.DeleteProduct)

	adminGroup.POST("/change_user_info", middleware.JWTAuthMiddleware(), api.ChangeUserInfo)
	adminGroup.DELETE("/delete_user", middleware.JWTAuthMiddleware(), api.DeleteUser)

	h.Spin()
}
