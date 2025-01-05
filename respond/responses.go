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

var (
	Ok = Response{ //正常
		Status: "10000",
		Info:   "success",
	}

	WrongID = Response{ //用户ID错误
		Status: "40001",
		Info:   "wrong userid",
	}

	WrongPwd = Response{ //密码错误
		Status: "40002",
		Info:   "wrong password",
	}

	InvalidID = Response{ //用户ID无效
		Status: "40003",
		Info:   "the userid already exists",
	}

	MissingParam = Response{ //缺少参数
		Status: "40004",
		Info:   "missing param",
	}

	WrongParamType = Response{ //参数错误
		Status: "40005",
		Info:   "wrong param type",
	}
)