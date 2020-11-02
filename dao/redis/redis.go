package redis

import (
	"bluebell/settings"
	"fmt"
	"github.com/go-redis/redis"
	"strconv"
)

// 声明一个全局的rdb变量
var rdb *redis.Client

// 初始化连接
func Init(conf *settings.RedisConfig) (err error) {
	add := fmt.Sprintf("%s:%s",conf.Host,strconv.Itoa(conf.Port))
	fmt.Printf("\n\n add = %v",add)
	rdb = redis.NewClient(&redis.Options{
		Addr: add,
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	_, err = rdb.Ping().Result()
	return err
}

func Close()  {
	_ = rdb.Close()
	return
}