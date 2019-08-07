package service

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"im-go/model"
	"im-go/setting"
	"log"
)

var DbEngin *gorm.DB

func init() {
	var err error
	var config = setting.DatabaseSetting
	dsName := config.Username + ":" + config.Password + "@(" + config.Host + ":" + config.Port + ")/" +
		config.Database + "?charset=" + config.Charset +
		"&parseTime=True&loc=Local"
	DbEngin, err = gorm.Open(config.Driver, dsName)
	if err != nil {
		log.Fatal(err)
	}
	DbEngin.AutoMigrate(&model.User{}, &model.Friend{}, &model.FriendApply{}, &model.Massage{}, &model.Seckill{}, &model.Good{})
	DbEngin.LogMode(true)
	fmt.Println("init database ok")
}

func CloseDB() {
	DbEngin.Close()
}
