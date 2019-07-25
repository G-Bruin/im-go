package controller

import (
	"im-go/helper"
	"im-go/service"
	"net/http"
)

var serviceLanguage service.LanguageService

func Create(w http.ResponseWriter, request *http.Request) {
	//获取post解析参数
	//request.ParseForm()
	//name := request.PostForm.Get("name")
	//code := request.PostForm.Get("code")

	//获取get请求数据
	vars := request.URL.Query()
	name := vars.Get("name")
	code := vars.Get("code")

	user, err := serviceLanguage.Create(name, code)
	if err != nil {
		helper.ResponseFail(w, 400, "失败了")
	} else {
		helper.ResponseOk(w, 200, user, "")
	}
}
