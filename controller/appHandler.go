package controller

import (
	"fmt"
	"log"
	"net/http"
	"web/common"
	"web/helper"
	"html/template"
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
		fmt.Println("HandleFunc     " + v.Name())
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

func HelloWorld(w http.ResponseWriter, request *http.Request) {
	map2 := make(map[string]string, 100)
	map2["key"] = "dddd"
	helper.ResponseOk(w, 200, map2, "asdasda")
	//fmt.Println(DbEngin)
	//data := returnData{Title: "sdada"}
	//fmt.Println(data)
	//common.RenderTemplate(w, "hello", &data)
}

func ServerSetup(port string) {
	log.Println("start setup server:"+port)
	http.HandleFunc("/helloworld", common.LogPanics(HelloWorld))

	RegisterView()

	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal("ListenAndServe err: ", err)
	}
}
