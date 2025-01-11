package api

import (
	"OnlineMall/model"
	"OnlineMall/respond"
	"OnlineMall/service"
	"context"
	"errors"
	"fmt"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"strconv"
)

func AddProduct(ctx context.Context, c *app.RequestContext) {
	handlerID := int(c.GetFloat64("user_id")) //从上下文中获取用户的id
	// 1.从请求中获取商品信息
	product := model.AddProduct{}
	err := c.BindJSON(&product)
	if err != nil {
		c.JSON(consts.StatusBadRequest, respond.WrongParamType)
		return
	}
	// 2.保存商品信息
	err = service.AddProduct(product, handlerID)
	if err != nil {
		switch {
		case errors.Is(err, respond.ErrUnauthorized), errors.Is(err, respond.ErrCategoryNotExists), errors.Is(err, respond.ParamTooLong):
			c.JSON(consts.StatusUnauthorized, err)
			return
		default:
			c.JSON(consts.StatusInternalServerError, respond.InternalError(err))
			return
		}
	}
	c.JSON(consts.StatusOK, respond.Ok)
}

func ChangeProduct(ctx context.Context, c *app.RequestContext) {
	handlerID := int(c.GetFloat64("user_id")) //从上下文中获取用户的id
	// 1.从请求中获取商品信息
	product := model.AddProduct{}
	id := c.Query("id")
	intID, err := strconv.ParseInt(id, 10, 0)
	if err != nil {
		c.JSON(consts.StatusBadRequest, respond.WrongParamType)
		return
	}
	err = c.BindJSON(&product)
	if err != nil {
		c.JSON(consts.StatusBadRequest, respond.WrongParamType)
		return
	}
	// 2.保存商品信息
	err = service.ChangeProduct(int(intID), product, handlerID)
	if err != nil {
		switch {
		case errors.Is(err, respond.ErrUnauthorized), errors.Is(err, respond.ErrCategoryNotExists),
			errors.Is(err, respond.ErrProductNotExists), errors.Is(err, respond.ParamTooLong):
			c.JSON(consts.StatusUnauthorized, err)
			return
		default:
			c.JSON(consts.StatusInternalServerError, respond.InternalError(err))
			return
		}
	}
	c.JSON(consts.StatusOK, respond.Ok)
}

func GetProductsInManyWays(ctx context.Context, c *app.RequestContext) {
	// 1.从请求中获取商品信息
	productID := c.Query("product_id")
	keyWord := c.Query("keyword")
	categoryID := c.Query("category_id")
	//如果没有传入参数，则默认为0
	if productID == "" {
		productID = "0"
	}
	if categoryID == "" {
		categoryID = "0"
	}
	intProductID, err := strconv.ParseInt(productID, 10, 0)
	if err != nil {
		fmt.Println(1)
		c.JSON(consts.StatusBadRequest, respond.WrongParamType)
		return
	}
	intCategoryID, err := strconv.ParseInt(categoryID, 10, 0)
	if err != nil {
		fmt.Println(2)
		c.JSON(consts.StatusBadRequest, respond.WrongParamType)
		return
	}
	//2.调用service层函数
	products, err := service.ShowProductInManyWays(int(intProductID), keyWord, int(intCategoryID))
	if err != nil {
		switch {
		case errors.Is(err, respond.CantFindProduct), errors.Is(err, respond.EmptyProductList), errors.Is(err, respond.ErrProductNotExists):
			c.JSON(consts.StatusNotFound, err)
			return
		default:
			c.JSON(consts.StatusInternalServerError, respond.InternalError(err))
			return
		}
	}
	c.JSON(consts.StatusOK, respond.Respond(respond.Ok, products))
}

func DeleteProduct(ctx context.Context, c *app.RequestContext) {
	handlerID := int(c.GetFloat64("user_id")) //从上下文中获取用户的id
	// 1.从请求中获取商品id
	id := c.Query("id")
	intID, err := strconv.ParseInt(id, 10, 0)
	if err != nil {
		c.JSON(consts.StatusBadRequest, respond.WrongParamType)
		return
	}
	// 2.删除商品
	err = service.DeleteProduct(int(intID), handlerID)
	if err != nil {
		switch {
		case errors.Is(err, respond.ErrUnauthorized), errors.Is(err, respond.ErrProductNotExists):
			c.JSON(consts.StatusUnauthorized, err)
			return
		default:
			c.JSON(consts.StatusInternalServerError, respond.InternalError(err))
			return
		}
	}
	c.JSON(consts.StatusOK, respond.Ok)
}
