package controller

import (
	"im-go/middleware/validator"
	"net/http"
)

func Register(w http.ResponseWriter, request *http.Request) {
	request.ParseForm()
	postData := validator.User{
		Name:     request.PostForm.Get("name"),
		Password: request.PostForm.Get("password"),
	}
	//前置数据验证
	postData.RegisterValidator(w)

}
