package service

import "im-go/model"

type SeckillService model.Seckill

func (service *SeckillService) CreateSeckill() (result model.Seckill) {
	var db = DbEngin
	tmp := model.Seckill{}
	tmp.UserId = service.UserId
	tmp.GoodId = service.GoodId
	db.Create(&tmp)
	return tmp
}
