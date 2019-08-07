package controller

import (
	"fmt"
	"im-go/helper"
	"im-go/service"
	"im-go/validator"
	"net/http"
)

var userControl service.UserService

func Register(w http.ResponseWriter, request *http.Request) {
	request.ParseForm()
	postData := validator.User{
		Name:     request.PostForm.Get("name"),
		Password: request.PostForm.Get("password"),
	}
	fmt.Println(postData)
	//前置数据验证
	postData.RegisterValidator(w)
	//定义map
	//test := make(map[string]interface{})
	//test["name"] = postData.Name
	userControl.Name = postData.Name
	result := userControl.FindUser()
	if result.Id > 0 {
		helper.ResponseFail(w, 403, "用户已经存在")
	}
	userControl.Password = postData.Password

	helper.ResponseOk(w, 200, result, "获取数据成功")

}
