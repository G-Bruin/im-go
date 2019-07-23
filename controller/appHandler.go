package controller

import (
	"fmt"
	"log"
	"net/http"
	"web/common"
)

type returnData struct {
	Title string
}

/* handle a simple get request */
func HelloWorld(w http.ResponseWriter, request *http.Request) {
	data := returnData{Title: "sdada"}
	fmt.Println(data)
	common.RenderTemplate(w, "hello", &data)
}

func ServerSetup(port string) {
	log.Println("start setup server:")
	http.HandleFunc("/", common.LogPanics(HelloWorld))
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal("ListenAndServe err: ", err)
	}
}
