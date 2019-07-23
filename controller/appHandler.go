package controller

import (
	"fmt"
	"html/template"
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
	renderTemplate(w, "hello", &data)
}

func ServerSetup(port string) {
	log.Println("start setup server:")
	http.HandleFunc("/", common.LogPanics(HelloWorld))
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal("ListenAndServe err: ", err)
	}
}

func renderTemplate(w http.ResponseWriter, tmpl string, p *returnData) {
	t, err := template.ParseFiles("tmpl/" + tmpl + ".gtpl")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	t.Execute(w, p)
}
