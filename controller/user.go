package controller

import (
	valid "github.com/asaskevich/govalidator"
	"net/http"
)

type Post struct {
	Title string `valid:"alphanum,required,stringlength(2|7)"`
	//Message  string `valid:"duck,ascii"`
	//Message2 string `valid:"animal(dog)"`
	AuthorIP string `valid:"ipv4"`
	Date     string `valid:"-"`
}

func Register(w http.ResponseWriter, request *http.Request) {
	request.ParseForm()
	//name := request.PostForm.Get("name")
	//password := request.PostForm.Get("password")

	post := &Post{
		Title: "MyEx6",
		//Message:  "duck",
		//Message2: "dog",
		AuthorIP: "123.234.54.3",
	}

	result, err := valid.ValidateStruct(post)
	if err != nil {
		println(" " + err.Error())
	}
	println(result)

}
