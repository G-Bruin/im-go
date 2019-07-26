package service

import (
	"fmt"
	"im-go/model"
)

type UserService model.User

func (service *UserService) AddUser() {

	menu := map[string]interface{}{
		"name": service.Name,
		"id":   service.ID,
	}
	fmt.Println(menu)
}
