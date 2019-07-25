package main

import (
	"im-go/controller"
	"im-go/service"
)

func main() {
	//defer 关闭 db
	defer service.CloseDB()
	//加载路由 启动http服务
	controller.Init()
}
