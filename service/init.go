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
	var config = setting.DatabaseSetting
	dsName := config.Username + ":" + config.Password + "@(" + config.Host + ":" + config.Port + ")/" +
		config.Database + "?charset=" + config.Charset +
		"&parseTime=True&loc=Local"
	db, err := gorm.Open(config.Driver, dsName)
	if err != nil {
		log.Fatal(err)
	}
	DbEngin = db
	db.AutoMigrate(&model.Email{}, &model.Address{}, &model.CreditCard{}, &model.User{}, &model.Language{})
	db.LogMode(true)
	fmt.Println("init database ok")
}

func CloseDB() {
	DbEngin.Close()
}
