package service

import (
	"im-go/model"
)

type LanguageService struct{}

func (service *LanguageService) Create(name string, code string) (language model.Language, err error) {
	tmp := model.Language{
		Name: name,
		Code: code,
	}
	DbEngin.Create(&tmp)
	//res := DbEngin.Where("name = ? AND code = ?", name, code).Find(&tmp)
	return tmp, nil
}

func (service *LanguageService) Find(input map[string]string) (language model.Language, err error) {
	tmp := model.Language{}
	if _, ok := input["id"]; ok {
		DbEngin = DbEngin.Where("id = ?", input["id"])
	}
	DbEngin.Find(&tmp)
	return tmp, nil
}
