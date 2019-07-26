package setting

import (
	"fmt"
	"github.com/go-redis/redis"
	"log"
	"strconv"
)

var RdEngin *redis.Client

func InitRedis() {
	RdEngin = redis.NewClient(&redis.Options{
		Addr:     RedisSetting.Host + ":" + strconv.Itoa(RedisSetting.Port),
		Password: RedisSetting.Password,
		DB:       RedisSetting.Db,
	})

	_, err := RdEngin.Ping().Result()

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("init redis ok")
}
