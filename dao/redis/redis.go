package redis

import (
	"fmt"
	"github.com/go-redis/redis"
	"github.com/spf13/viper"
)

// 声明一个全局的rdb变量
var rdb *redis.Client

// 初始化连接
func Init() (err error) {
	add := fmt.Sprintf("%s:%s",viper.GetString("redis.host"),string(viper.GetString("redis.port")))
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