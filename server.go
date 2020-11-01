package main

import (
	"bluebell/dao/mysql"
	"bluebell/dao/redis"
	"bluebell/logger"
	"bluebell/routes"
	"bluebell/settings"
	"fmt"
	"go.uber.org/zap"
)

// Go Web 开发较为通用的脚手架模板

func main() {
	// 1 加载配置文件 -》
	if err := settings.Init();err != nil{
		fmt.Printf("init setting failed err = %s ",err)
	}
	// 2 初始化日志
	if err := logger.Init();err != nil{
		fmt.Printf("logger.Init() failed err = %s ",err)
	}
	zap.L().Info("logger init success")
	// 3 初始化mysql连接
	if err := mysql.Init();err != nil{
		fmt.Printf("Mysql.Init() failed err = %s ",err)
	}
	defer mysql.Close()
	zap.L().Info("Mysql init Success")
	// 4 初始化redis连接
	if err := redis.Init();err != nil{
		fmt.Printf("redis.Init() failed err = %s ",err)
	}
	zap.L().Info("redis init Success")
	defer redis.Close()
	fmt.Printf("all success")
	// 5 注册路由
	r := routes.SetupRouter()
	// 6 启动服务 （优雅关机）
	r.Run()
}
