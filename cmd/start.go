package cmd

import (
	"OnlineMall/dao"
	"OnlineMall/routers"
	"fmt"
)

func Start() {
	err := dao.ConnectDB()
	if err != nil {
		fmt.Println(err)
	}
	routers.RegisterRouters() //注册路由并启动服务
}
