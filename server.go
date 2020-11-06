package main

import (
	"bluebell/logger"
	"bluebell/models"
	"bluebell/routes"
	"bluebell/settings"
	"bluebell/util/snowflake"
	"fmt"
	"go.uber.org/zap"
)

// Go Web 开发较为通用的脚手架模板


func Init()  {
	// 1 加载配置文件 -》
	if err := settings.Init();err != nil{
		fmt.Printf("init setting failed err = %s ",err)
	}
	// 2 初始化日志
	if err := logger.Init(settings.Conf.LogConfig,settings.Conf.Mode);err != nil{
		fmt.Printf("logger.Init() failed err = %s ",err)
	}
	zap.L().Info("logger init success")
	// 3 初始化mysql连接
	if err := models.Init(settings.Conf.MySQLConfig);err != nil{
		fmt.Printf("Mysql.Init() failed err = %s ",err)
	}
	zap.L().Info("Mysql init Success")
	// 4 初始化redis连接
	if err := models.RedisInit(settings.Conf.RedisConfig);err != nil{
		fmt.Printf("redis.Init() failed err = %s ",err)
	}
	zap.L().Info("redis init Success")
	// 初始化ID生成器- 雪花算法
	fmt.Printf("\nstarttime = %v",settings.Conf.StartTime)
	if err := snowflake.Init(settings.Conf.StartTime,settings.Conf.MachineID); err != nil{
		zap.L().Error("generate id failed!")
		return
	}

	fmt.Printf("all success")
}


func main() {
	Init()
	defer models.Close()
	defer models.RedisClose()
	// 5 注册路由
	r := routes.SetupRouter()
	// 6 启动服务 （优雅关机）
	r.Run()
}
