package models

import (
	"bluebell/settings"
	"fmt"
	"testing"
)

func TestUserDAO_InsertUser(t *testing.T) {
	// 1 加载配置文件 -》
	if err := settings.Init();err != nil{
		fmt.Printf("init setting failed err = %s ",err)
	}
	if err := Init(settings.Conf.MySQLConfig);err != nil{
		fmt.Printf("Mysql.Init() failed err = %s ",err)
	}
	user := &User{
		UserId:   1,
		Username: "wanghaoxin",
		Password: "whx123456",
	}

	_,err := NewUserDaoInstance().InsertUser(user)
	if err != nil{
		fmt.Printf("\n%s\n",err)
	}
}

func TestUserDAO_QueryUser(t *testing.T) {
	tt := new(User)
	fmt.Println(tt)
	// 1 加载配置文件 -》
	if err := settings.Init();err != nil{
		fmt.Printf("init setting failed err = %s ",err)
	}
	if err := Init(settings.Conf.MySQLConfig);err != nil{
		fmt.Printf("Mysql.Init() failed err = %s ",err)
	}

	val,res := NewUserDaoInstance().CheckRecordExist("wanghaoxin")
	if res {
		fmt.Printf("the record has been recoreded")
		fmt.Println(val)
	}
}
