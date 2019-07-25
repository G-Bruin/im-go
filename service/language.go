package service

import (
	"fmt"
	"web/model"
)

type LanguageService struct {
}

func (service *LanguageService) Create(name string, code string) (language model.Language, err error) {
	fmt.Println(2222)
	tmp := model.Language{
		Name: name,
		Code: code,
	}
	fmt.Println(tmp)
	DbEngin.Create(&tmp)
	//res := DbEngin.Where("name = ? AND code = ?", name, code).Find(&tmp)
	fmt.Println(DbEngin)
	return tmp, nil
}
