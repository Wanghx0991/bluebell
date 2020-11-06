package logic

import (
	"bluebell/models"
	"bluebell/settings"
	"bluebell/util/snowflake"
	"fmt"
	"testing"
)

func TestSignUp(t *testing.T) {
	// 1 加载配置文件 -》
	if err := settings.Init();err != nil{
		fmt.Printf("init setting failed err = %s ",err)
	}
	if err := models.Init(settings.Conf.MySQLConfig);err != nil{
		fmt.Printf("Mysql.Init() failed err = %s ",err)
	}
	P := &models.ParamSignUp{
		UserName:   "lalala",
		Password:   "123123",
		RePassword: "123123",
	}
	snowflake.Init("2020-10-20",1)
	SignUp(P)
}

func TestLogin(t *testing.T) {
	// 1 加载配置文件 -》
	if err := settings.Init();err != nil{
		fmt.Printf("init setting failed err = %s ",err)
	}
	if err := models.Init(settings.Conf.MySQLConfig);err != nil{
		fmt.Printf("Mysql.Init() failed err = %s ",err)
	}
	P := &models.ParamLogin{
		UserName:   "hahaha",
		Password:   "123123",
	}
	var token string
	token ,err := Login(P)
	if err != nil{
		fmt.Sprintf("%v",err)
	}
	fmt.Printf("已登录,token = %v",token)
}