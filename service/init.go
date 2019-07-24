package service

import (
	"errors"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"log"
	"web/model"
)

var DbEngin *gorm.DB

func init() {
	driveName := "mysql"
	dsName := "root:root@(127.0.0.1:3306)/chat?charset=utf8&parseTime=True&loc=Local"
	err := errors.New("")
	db, err := gorm.Open(driveName, dsName)
	if err != nil && "" != err.Error() {
		log.Fatal(err.Error())
	}
	defer db.Close()
	DbEngin = db

	db.AutoMigrate(&model.Email{}, &model.Address{}, &model.CreditCard{}, &model.User{}, &model.Language{})
	db.LogMode(true)
	fmt.Println("init data base ok")

}
