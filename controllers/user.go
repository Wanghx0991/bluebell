package controllers

import (
	"bluebell/logic"
	"bluebell/models"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
)

func SinUpHandler(c *gin.Context)  {
	// 1.获取参数和参数校验 -> 本层 Controller进行
		//把想要的数据绑定到结构体里去
	var p  = new(models.ParamSignUp)
	//shouldbindjson只能校验请求格式是否是json,以及是否是指定的string类型,具体string里是否为空等,需要手动校验
	if err := c.ShouldBindJSON(p);err != nil{
		// 请求参数有误
		zap.L().Error("SignUp with invalid param",zap.Error(err))
		ResponseError(c,CodeInvalidParam)
		return
	}
	// 手动对请求参数进行详细的业务规则校验
	if err :=validateParam(p);err!=nil{
		e := fmt.Sprintf("%s",err.Error())
		zap.L().Error(e)
		ResponseErrorWithMsg(c,CodeSuccess,e)
		return
	}
	// better choice: 使用validator库进行参数校验 -> gin框架里 tag 是binding, 其他条件下是validator

	// 2. 业务处理 -> 放在logic层
	if err := logic.SignUp(p); err != nil{
		zap.L().Error(fmt.Sprintf("%v",err))
		ResponseErrorWithMsg(c,CodeSuccess,err)
		return
	}

	// 3. 返回响应
	ResponseSuccess(c,"注册成功")
}

func validateParam(p *models.ParamSignUp)  error {
	var err error
	if p == nil{
		err = errors.New("请求参数为空")
		return err
	}
	if len(p.UserName) == 0 || len(p.Password) == 0 || len(p.RePassword) == 0{
		err = errors.New("请求参数为空")
		return err
	}
	if p.Password != p.RePassword{
		err = errors.New("两次密码不一致")
		return err
	}
	return nil
}

func LogInHandler(c *gin.Context)  {
	// 获取请求参数及参数校验
	var p = new(models.ParamLogin)
	if err := c.ShouldBindJSON(p);err != nil{
		// 请求参数有误
		zap.L().Error("Login with invalid param",zap.Error(err))
		c.JSON(http.StatusOK,gin.H{
			"msg": err.Error(),
		})
		return
	}
	//业务逻辑处理
	var token string
	token ,err := logic.Login(p)
	if  err != nil{

		zap.L().Error(fmt.Sprintf("%v",err))
		c.JSON(http.StatusOK,"登录失败")
		return
	}
	// 返回响应
	c.JSON(http.StatusOK,gin.H{
		"msg":"登陆成功",
		"token":token,
	})

}

