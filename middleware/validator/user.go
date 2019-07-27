package validator

import (
	"github.com/asaskevich/govalidator"
	"im-go/helper"
	"net/http"
	"regexp"
)

type User struct {
	Name     string `valid:"regname~用户名必须是字母+数字,required~用户名必填,stringlength(3|10)~用户名长度是3到10位"`
	Password string `valid:"regpasswd~密码必须是字母+数字,required~密码必填,stringlength(8|20)~密码长度是8到20位"`
}

func (target User) RegisterValidator(w http.ResponseWriter) {
	//自定义tag验证函数
	govalidator.TagMap["regname"] = govalidator.Validator(func(str string) bool {
		ok, _ := regexp.Match("[A-Za-z].*[\\d]|[\\d].*[A-Za-z]", []byte(str))
		return ok
	})
	govalidator.TagMap["regpasswd"] = govalidator.Validator(func(str string) bool {
		ok, _ := regexp.Match("^[A-Za-z].*[\\d]|[\\d].*[A-Za-z]", []byte(str))
		return ok
	})
	if _, err := govalidator.ValidateStruct(target); err != nil {
		helper.ResponseFail(w, 422, err.Error())
	}
}
