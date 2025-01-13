package respond

type Response struct { //响应结构体
	Status string `json:"status"`
	Info   string `json:"info"`
}

type FinalResponse struct { //最终响应结构体
	Status string      `json:"status"`
	Info   string      `json:"info"`
	Data   interface{} `json:"data"`
}

// 实现error接口
func (r Response) Error() string { // 实现 error 接口
	return r.Info
}

func Respond(response Response, data interface{}) FinalResponse { //传入一个响应结构体和数据，返回一个最终响应结构体
	var finalResponse FinalResponse
	finalResponse.Status = response.Status
	finalResponse.Info = response.Info
	finalResponse.Data = data
	return finalResponse
}

func InternalError(err error) Response { //服务器错误
	return Response{
		Status: "500",
		Info:   err.Error(),
	}
}

var ( //请求相关的响应
	Ok = Response{ //正常
		Status: "10000",
		Info:   "success",
	}

	WrongName = Response{ //用户名错误
		Status: "40001",
		Info:   "wrong username",
	}

	WrongPwd = Response{ //密码错误
		Status: "40002",
		Info:   "wrong password",
	}

	InvalidName = Response{ //用户名无效
		Status: "40003",
		Info:   "the username already exists",
	}

	MissingParam = Response{ //缺少参数
		Status: "40004",
		Info:   "missing param",
	}

	WrongParamType = Response{ //参数错误
		Status: "40005",
		Info:   "wrong param type",
	}

	ParamTooLong = Response{ //参数过长
		Status: "40006",
		Info:   "param too long",
	}

	WrongUsernameOrPwd = Response{ //用户名或密码错误
		Status: "40007",
		Info:   "wrong username or password",
	}

	WrongGender = Response{ //性别错误
		Status: "40008",
		Info:   "wrong gender",
	}

	MissingToken = Response{ //缺少token
		Status: "40009",
		Info:   "missing token",
	}

	InvalidTokenSingingMethod = Response{ //jwt token签名方法无效
		Status: "40010",
		Info:   "invalid signing method",
	}

	InvalidToken = Response{ //无效token
		Status: "40011",
		Info:   "invalid token",
	}

	InvalidClaims = Response{ //无效声明
		Status: "40012",
		Info:   "invalid claims",
	}

	WrongUserID = Response{ //用户ID错误
		Status: "40013",
		Info:   "wrong userid",
	}

	ErrUnauthorized = Response{ //未授权，没有权限
		Status: "40014",
		Info:   "unauthorized",
	}

	ErrCategoryNotExists = Response{ //分类不存在
		Status: "40015",
		Info:   "category not exists",
	}

	ErrCategoryNameExists = Response{ //分类名已存在
		Status: "40016",
		Info:   "category name exists",
	}

	ErrProductNotExists = Response{ //商品不存在
		Status: "40017",
		Info:   "product not exists",
	}

	CantFindProduct = Response{ //找不到商品
		Status: "40018",
		Info:   "can't find product",
	}

	EmptyProductList = Response{ //商品列表为空
		Status: "40019",
		Info:   "product list is empty",
	}
	InvalidRefreshToken = Response{ //刷新令牌无效
		Status: "40020",
		Info:   "invalid refresh token",
	}
	ErrProductAlreadyInCart = Response{ //商品已在购物车中
		Status: "40021",
		Info:   "product already in cart",
	}
	ErrQuantityTooLarge = Response{ //数量太大
		Status: "40022",
		Info:   "quantity too large",
	}
	ErrUserDidntBuyThisProduct = Response{ //用户没有购买这个商品
		Status: "40024",
		Info:   "user didn't buy this product",
	}
	ErrUserHasAlreadyReviewed = Response{ //用户已经评论过了
		Status: "40025",
		Info:   "user has already reviewed",
	}
	ErrRatingOutOfRange = Response{ //评分超出范围
		Status: "40026",
		Info:   "rating out of range",
	}
	ErrCommentTooLong = Response{ //评论太长
		Status: "40027",
		Info:   "comment too long",
	}
)

var ( //服务器错误
	ErrOrderNotExists = Response{ //订单不存在
		Status: "50001",
		Info:   "order not exists",
	}
)
