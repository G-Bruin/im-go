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

	//client := setting.RdEngin
	//var err error
	//err = client.Set("433242342tt", "value888888", 0).Err()
	//if err != nil {
	//	panic(err)
	//}

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

func LanguageFind(w http.ResponseWriter, request *http.Request) {

	vars := request.URL.Query()
	id := vars.Get("id")
	test := make(map[string]interface{})
	test["id"] = id
	user, err := serviceLanguage.Find(test)
	if err != nil {
		helper.ResponseFail(w, 400, "失败了")
	} else {
		helper.ResponseOk(w, 200, user, "")
	}
}
