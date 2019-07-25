package setting

import (
	"fmt"
	"github.com/go-ini/ini"
	"log"
	"os"
)

//db配置文件
type DbSetting struct {
	Driver    string `ini:"driver"`
	Host      string `ini:"host"`
	Port      string `ini:"port"`
	Database  string `ini:"database"`
	Username  string `ini:"username"`
	Password  string `ini:"password"`
	Charset   string `ini:"charset"`
	Collation string `ini:"collation"`
}

var DatabaseSetting = &DbSetting{}

// app 配置文件
type AppSetting struct {
	Protocol string `ini:"protocol"`
	Port     string `ini:"http_port"`
	AppModel string `ini:"app_mode"`
}

var ServerSetting = &AppSetting{}

var cfg *ini.File

func init() {
	var err error
	cfg, err = ini.Load("setting/app.ini")
	if err != nil {
		fmt.Printf("Fail to read file: %v", err)
		os.Exit(1)
	}
	mapTo("database", DatabaseSetting)
	mapTo("server", ServerSetting)
	fmt.Println("init app setting ok")
}

func mapTo(section string, v interface{}) {
	err := cfg.Section(section).MapTo(v)
	if err != nil {
		log.Fatalf("Cfg.MapTo "+section+" err: %v", err)
	}
}
