package setting

import (
	"fmt"
	"github.com/go-redis/redis"
	"log"
	"strconv"
)

var RdEngin *redis.Client

//var Pubsub *redis.PubSub

func InitRedis() {
	RdEngin = redis.NewClient(&redis.Options{
		Addr:     RedisSetting.Host + ":" + strconv.Itoa(RedisSetting.Port),
		Password: RedisSetting.Password,
		DB:       RedisSetting.Db,
	})

	if _, err := RdEngin.Ping().Result(); err != nil {
		log.Fatal("Connect to redis error", err)
		return
	}
	//Pubsub = Redis.Subscribe("ws:msgpool:" + config.Config.Address)
	//if _, err := Pubsub.Receive(); err != nil {
	//	panic(err)
	//}
	fmt.Println("init redis ok")
}
