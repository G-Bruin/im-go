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

func (service *LanguageService) Find(input map[string]interface{}) (language model.Language, err error) {
	var db = DbEngin
	tmp := model.Language{}
	if _, ok := input["id"]; ok {
		db = db.Where("id = ?", input["id"])
	}
	db.Find(&tmp)
	return tmp, nil
}
