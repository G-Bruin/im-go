package common

import (
	"fmt"
	"html/template"
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
