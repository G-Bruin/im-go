package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"net/url"
	//"strconv"
	"strings"
	"time"
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
		req, _ = http.NewRequest(method, urlVal, strings.NewReader(data))
	}
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
	fmt.Println(b)
	fmt.Println(string(b))
}

//将get请求的参数进行转义
func getParseParam(param string) string {
	return url.PathEscape(param)
}

func sendData(ch chan string) {
	//for i := 0; i < 10; i++ {
	for {
		ch <- "http://shopdao.com/v2/api/service/card/join"
		ch <- "http://shopdao.com/v2/api/service/card/join"
		ch <- "http://shopdao.com/v2/api/service/card/join"
		ch <- "http://shopdao.com/v2/api/service/card/join"
		ch <- "http://shopdao.com/v2/api/service/card/join"
		ch <- "http://shopdao.com/v2/api/service/card/join"
		ch <- "http://shopdao.com/v2/api/service/card/join"
		ch <- "http://shopdao.com/v2/api/service/card/join"
		ch <- "http://shopdao.com/v2/api/service/card/join"
		ch <- "http://shopdao.com/v2/api/service/card/join"
		time.Sleep(1e9 * 1)
	}
}

func getData(ch chan string, num chan bool) {

	//var arr1 = [22]int{99110, 99111, 99112, 99113, 99114, 99115, 99116, 99117, 99118, 99119, 99120,
	//	99121, 99122, 99123, 99124, 99125, 99126, 99127, 99128, 99129, 99130, 99131}

	for v := range ch {
		//activity_id := arr1[rand.Intn(20)]
		//fmt.Println("开始抢了")
		//union_id := "oBZTV1RzRl-VKP2bjz1eL4tVj" + GetRandomString(4)
		//httpHandle("POST", v, "activity_id="+strconv.Itoa(activity_id)+"&union_id=oBZTV1RzRl-VKP2bjz1eL4tVj"+union_id+"&tag=1674779&others=1111")
		httpHandle("POST", v, "activity_id=167727&union_id=oBZTV1cHTrm7WAJIRQWtHQKoGIPo")
	}
	//num <- false

}

//测试
func main() {
	//for {
	//	//var queueName = [3]string{"high", "default", "slow"}
	//	fmt.Println(rand.Intn(3))
	//	//fmt.Println(queueName[rand.Intn(2)])
	//	time.Sleep(1e9)
	//}
	numChan := make(chan string)
	num := make(chan bool)
	go sendData(numChan)
	go getData(numChan, num)
	<-num
}

func GetRandomString(l int) string {
	str := "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	bytes := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < l; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}
