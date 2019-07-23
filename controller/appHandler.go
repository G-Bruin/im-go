package controller

import (
	"log"
	"net/http"
	"web/common"
	"web/helper"
)

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
	log.Println("start setup server:")
	http.HandleFunc("/", common.LogPanics(HelloWorld))
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal("ListenAndServe err: ", err)
	}
}
