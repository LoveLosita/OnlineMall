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

	h.GET("/search_products", api.SearchForProducts)
	h.GET("/show_all_products", api.ShowAllProducts)
	h.GET("/show_category_products", api.ShowACategoryProducts)
	h.GET("/view_product", middleware.JWTTokenAuthForProductHistory(), api.ShowSingleProduct)
	h.GET("/show_product_reviews", api.ShowAProductReviews)

	userGroup.PUT("/register", api.UserRegister)
	userGroup.POST("/login", api.UserLogin)
	userGroup.POST("/change_username_or_password", middleware.JWTTokenAuth(), api.ChangeUserPasswordOrName)
	userGroup.GET("/refresh_token", middleware.JWTTokenAuth(), api.RefreshTokenHandler)
	userGroup.POST("/add_to_cart", middleware.JWTTokenAuth(), api.AddProductToCart)
	userGroup.POST("/place_order", middleware.JWTTokenAuth(), api.PlaceOrder)
	userGroup.POST("/rate_and_review", middleware.JWTTokenAuth(), api.RateAndReviewProduct)
	userGroup.POST("/reply_to_review", middleware.JWTTokenAuth(), api.ReplyToReview)
	userGroup.GET("/search_product_in_cart", middleware.JWTTokenAuth(), api.SearchForProductsInCart)
	userGroup.GET("/show_products_in_cart", middleware.JWTTokenAuth(), api.GetUserCart)
	userGroup.GET("/show_history", middleware.JWTTokenAuth(), api.ShowUserViewProductHistory)

	merchantGroup.PUT("/add_product", middleware.JWTTokenAuth(), api.AddProduct)
	merchantGroup.PUT("/add_category", middleware.JWTTokenAuth(), api.AddCategory)
	merchantGroup.POST("/change_product", middleware.JWTTokenAuth(), api.ChangeProduct)
	merchantGroup.DELETE("/delete_product", middleware.JWTTokenAuth(), api.DeleteProduct)
	merchantGroup.GET("/search_for_comments", middleware.JWTTokenAuth(), api.SearchForAProductReview)

	adminGroup.POST("/change_user_info", middleware.JWTTokenAuth(), api.ChangeUserInfo)
	adminGroup.DELETE("/delete_user", middleware.JWTTokenAuth(), api.DeleteUser)
	adminGroup.DELETE("/delete_comment", middleware.JWTTokenAuth(), api.DeleteReview)

	h.Spin()
}
