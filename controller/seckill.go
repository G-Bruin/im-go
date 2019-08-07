package controller

import (
	"fmt"
	"im-go/model"
	"im-go/service"
	"net/http"
	"strconv"
	"time"
)

func Apply(w http.ResponseWriter, request *http.Request) {
	request.ParseForm()

	postData := map[string]string{
		"user_id": request.PostForm.Get("user_id"),
		"good_id": request.PostForm.Get("good_id"),
	}

	ch := make(chan map[string]string, 10)
	go sendData(ch, postData)
	go getData(ch)
	time.Sleep(1e9 * 2)
	//data := new(interface{})
	//helper.ResponseOk(w, 200, data, "获取数据成功")
}

func getData(ch chan map[string]string) {
	for v := range ch {
		good_id, _ := strconv.Atoi(v["good_id"])
		tx := service.DbEngin.Begin()
		goodTmp := model.Good{}
		goodDb := service.DbEngin.Where("id = ?", good_id)
		goodDb.First(&goodTmp)
		if goodTmp.Total <= 0 {
			tx.Rollback()
			fmt.Println("商品已经抢完")
			return
		}

		if goodTmp.Id == 0 {
			tx.Rollback()
			fmt.Println("数据不存在")
			return
		}

		//fmt.Println(goodTmp)

		//插入数据

		goodTmp.Total = goodTmp.Total - 1
		goodDb.Save(&goodTmp)

		seckillTmp := model.Seckill{}
		seckillTmp.UserId, _ = strconv.Atoi(v["user_id"])
		seckillTmp.GoodId = good_id
		service.DbEngin.Save(&seckillTmp)
		if seckillTmp.Id <= 0 {
			tx.Rollback()
			fmt.Println("插入数据失败")
			return
		}

		tx.Commit()
		//fmt.Println("抢到了，兄弟")
		time.Sleep(1e9 * 3)
	}
}

func sendData(ch chan map[string]string, postData map[string]string) {
	ch <- postData
}
