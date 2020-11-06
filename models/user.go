package models

import (
	"bluebell/util/snowflake"
	"errors"
	"fmt"
	"sync"
)

type User struct {
	UserId   int64  `gorm:"column:user_id" json:"user_id"`
	Username string `gorm:"column:username" json:"username"`
	Password string `gorm:"column:password" json:"password"`
}
//数据访问对象
type UserDAO struct {

}

var userDao *UserDAO
var userOnce sync.Once

const Secret = "wanghaoxin"
func NewUserDaoInstance() *UserDAO {
	userOnce.Do(func() {
		userDao = &UserDAO{}
	})
	return userDao
}

//插入一条user数据
func (dao *UserDAO)InsertUser(user *User) (*User,error) {
	var err error
	t := new(User)
	t,rest := dao.CheckRecordExist(user.Username)
	if rest {
		fmt.Printf("The recored has been recorded!")
		err = errors.New("The recored has been recorded!")
		return t,err
	}

	// 加密
	user.Password = snowflake.EncryptPassword(user.Password, Secret)
	err = db.Create(user).Error
	if err != nil{
		return nil, err
	}
	return user,err
	//执行SQL语句入库
}

// 指定用户名的用户是否存在
func (dao *UserDAO)CheckRecordExist(name string) (*User,bool) {
	t := new(User)
	res := db.Where("username = ?",name).First(t)
	if res.RowsAffected != 0{
		return t,true
	}
	return nil,false
}
