package service

import (
	"fmt"
	"im-go/helper"
	"im-go/model"
)

type UserService model.User

//插入数据
func (service *UserService) AddUser() {
	//var db = DbEngin
	//tmp := model.User{}
	//
	//
	//fmt.Println(menu)
}

//查询数据
//func (service *UserService) FindUser(input map[string]interface{}) (result model.User) {
func (service *UserService) FindUser() (result model.User) {
	var db = DbEngin
	tmp := model.User{}

	//if _, ok := input["id"]; ok {
	//	db = db.Where("id = ?", input["id"])
	//}
	//if _, ok := input["name"]; ok {
	//	db = db.Where("name = ?", input["name"])
	//}

	if service.Id > 0 {
		db = db.Where("id = ?", service.Id)
	}
	if service.Name != "" {
		db = db.Where("name = ?", service.Name)
	}
	db.First(&tmp)

	if service.Password != "" {
		flag := helper.ValidatePasswd(tmp.Password, service.Password)
		fmt.Println("密码验证：", flag)
		// 验证失败
		if !flag {
			return model.User{}
		}
	}

	return tmp
}
