package main

import (
	"im-go/controller"
	"im-go/service"
	"im-go/setting"
)

func main() {
	//defer 关闭 db
	defer service.CloseDB()
	//init redis
	setting.InitRedis()
	//加载路由 启动http服务
	controller.Init()

}
