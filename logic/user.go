package logic

import (
	"bluebell/models"
	"bluebell/util/jwt"
	"bluebell/util/snowflake"
	"errors"
	"fmt"
	"go.uber.org/zap"
)

func SignUp(p *models.ParamSignUp) error {
	var err error
	// 1. 判断用户存不存在
	if _,res := models.NewUserDaoInstance().CheckRecordExist(p.UserName);res{
		zap.L().Info("The Username has been recorded !")
		err = errors.New("User Is Exist.")
		return err
	}
	// 生成userid
	userid := snowflake.GenID()
	user := &models.User{
		UserId:   userid,
		Username: p.UserName,
		Password: p.Password,
	}
	// 保存进数据库
	_, err = models.NewUserDaoInstance().InsertUser(user)
	if err != nil{
		zap.L().Error(fmt.Sprintf("%v",err))
		return err
	}
	return nil
}

func Login(p *models.ParamLogin) (string , error ){
	var err error
	// 1. 判断用户存不存在
	if entity,res := models.NewUserDaoInstance().CheckRecordExist(p.UserName);res{
		p.Password = snowflake.EncryptPassword(p.Password,models.Secret)
		if entity.Password == p.Password{
			// 生成JWT token
			// 通常放到请求头Header 里, postman的Authorization里有一个token,请求的时候拷贝进去就行
			 return jwt.GentToken(entity.UserId,entity.Username)
		}else {
			err = errors.New("Password is not same")
			return "",err
		}

	}
	zap.L().Error("用户不存在")
	err = errors.New("用户不存在")

	return "",err
}
