package helper

import (
	"encoding/json"
	"log"
	"net/http"
)

type RetrunData struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
	Rows    interface{} `json:"rows,omitempty"`
	Total   interface{} `json:"total,omitempty"`
}

func ResponseFail(w http.ResponseWriter, code int, msg string) {
	Resp(w, code, nil, msg)
	panic(msg)
}

func ResponseOk(w http.ResponseWriter, code int, data interface{}, msg string) {
	Resp(w, code, data, msg)
}

func ResponseList(w http.ResponseWriter, code int, lists interface{}, total interface{}) {
	//分页数目,
	RespList(w, code, lists, total)
}

func Resp(w http.ResponseWriter, code int, data interface{}, msg string) {
	w.Header().Set("Content-Type", "application/json")
	//设置200状态
	w.WriteHeader(http.StatusOK)
	//输出
	//定义一个结构体
	h := RetrunData{
		Code:    code,
		Message: msg,
		Data:    data,
	}
	//将结构体转化成json字符串
	ret, err := json.Marshal(h)
	if err != nil {
		log.Println(err.Error())
	}
	//输出
	w.Write(ret)
}
func RespList(w http.ResponseWriter, code int, data interface{}, total interface{}) {

	w.Header().Set("Content-Type", "application/json")
	//设置200状态
	w.WriteHeader(http.StatusOK)

	h := RetrunData{
		Code:  code,
		Rows:  data,
		Total: total,
	}
	//将结构体转化成json字符串
	ret, err := json.Marshal(h)
	if err != nil {
		log.Println(err.Error())
	}
	//输出
	w.Write(ret)
}
