package controller

import (
	"im-go/helper"
	"im-go/service"
	"net/http"
)

var helloService service.EmailService

func HelloWorld(w http.ResponseWriter, request *http.Request) {
	map2 := make(map[string]string, 100)
	map2["key"] = "dddd"
	helper.ResponseOk(w, 200, map2, "asdasda")
	helloService.AddFriend()

	//fmt.Println(DbEngin)
	//data := returnData{Title: "sdada"}
	//fmt.Println(data)
	//common.RenderTemplate(w, "hello", &data)
}
