package mysql

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)
var db *gorm.DB
func Init() (err error)  {
	db, err = gorm.Open("mysql", "root:whx1994927@/sql_bluebell?charset=utf8&parseTime=True&loc=Local")
 	if err != nil{
		return err
	}
	return nil
}
func Close()  {
	_ = db.Close()
	return
}

