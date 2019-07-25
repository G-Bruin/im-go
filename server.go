package main

import (
	"fmt"
	"html/template"
	"im-go/common"
	"im-go/controller"
	"im-go/service"
	"log"
	"net/http"
)

func RegisterView() {
	//一次解析出全部模板
	tpl, err := template.ParseGlob("view/*")
	if nil != err {
		log.Fatal(err)
	}
	//通过for循环做好映射
	for _, v := range tpl.Templates() {
		tplname := v.Name()
		fmt.Println("HandleFunc    " + v.Name())
		http.HandleFunc(tplname, func(w http.ResponseWriter,
			request *http.Request) {
			fmt.Println("parse     " + v.Name() + "==" + tplname)
			err := tpl.ExecuteTemplate(w, tplname, nil)
			if err != nil {
				log.Fatal(err.Error())
			}
		})
	}
}

func main() {

	log.Println("start setup server:9002")
	//defer 关闭 db
	defer service.CloseDB()

	http.HandleFunc("/helloworld", common.LogPanics(controller.HelloWorld))
	http.HandleFunc("/language", common.LogPanics(controller.Create))
	////加载所有模板页面
	////RegisterView()
	if err := http.ListenAndServe(":9002", nil); err != nil {
		log.Fatal("ListenAndServe err: ", err)
	}

}
