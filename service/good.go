package service

import (
	"im-go/model"
)

type GoodService model.Good

func (service *GoodService) FindGood() (result model.Good) {
	var db = DbEngin
	tmp := model.Good{}

	if service.Id > 0 {
		db = db.Where("id = ?", service.Id)
	}

	db.First(&tmp)

	return tmp
}

func (service *GoodService) UpdateGood() (result model.Good) {
	var db = DbEngin
	tmp := model.Good{}
	tmp.Id = service.Id
	tmp.Total = service.Total
	db.Save(&tmp)
	return tmp
}
