package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
)

func httpHandle(method, urlVal, data string) {
	client := &http.Client{}
	var req *http.Request
	if data == "" {
		urlArr := strings.Split(urlVal, "?")
		if len(urlArr) == 2 {
			urlVal = urlArr[0] + "?" + getParseParam(urlArr[1])
		}
		req, _ = http.NewRequest(method, urlVal, nil)
	} else {
		fmt.Println(data)
		req, _ = http.NewRequest(method, urlVal, strings.NewReader(data))
	}
	fmt.Println(strings.NewReader(data))

	//可以添加多个cookie
	cookie1 := &http.Cookie{Name: "X-Xsrftoken", Value: "111", HttpOnly: true}
	req.AddCookie(cookie1)

	req.Header.Add("X-Xsrftoken", "1111")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8") //设置Content-Type

	resp, err := client.Do(req)

	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	b, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(b))
}

//将get请求的参数进行转义
func getParseParam(param string) string {
	return url.PathEscape(param)
}

func sendData(ch chan string) {
	for {
		ch <- "http://www.baidu.com"
	}
}

func getData(ch chan string, num chan bool) {
	for v := range ch {
		fmt.Println("开始抢了")
		//httpHandle("POST", v, "333=1&3333=3233")
		httpHandle("GET", v, "")
	}
	//num <- false

}

//测试
func main() {
	numChan := make(chan string)
	num := make(chan bool)
	go sendData(numChan)
	go getData(numChan, num)
	<-num
}
