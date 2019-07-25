package controller

import (
	"fmt"
	"im-go/common"
	"im-go/setting"
	"log"
	"net/http"
)

func Init() {
	http.HandleFunc("/helloworld", common.LogPanics(HelloWorld))
	http.HandleFunc("/language", common.LogPanics(Create))
	var appSetting = setting.ServerSetting
	fmt.Println("启动http服务，端口" + appSetting.Port)

	if err := http.ListenAndServe(":"+appSetting.Port, nil); err != nil {
		log.Fatal("ListenAndServe err: ", err)
	}
}
