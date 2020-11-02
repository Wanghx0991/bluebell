package mysql

import (
	"bluebell/settings"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)
var db *gorm.DB
func Init(conf *settings.MySQLConfig) (err error)  {
	arg := fmt.Sprintf("%s:%s@/%s?charset=utf8&parseTime=True&loc=Local",conf.User,conf.Password,conf.Dbname)
	db, err = gorm.Open("mysql", arg)
 	if err != nil{
		return err
	}
	return nil
}
func Close()  {
	_ = db.Close()
	return
}

