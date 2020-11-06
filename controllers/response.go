package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

//定义程序中用到的响应的内容
/*
	"code": " 10001//程序中的错误码
	"msg": xx,//提示信息
	"data":{},//数据
*/

type ResponseData struct {
	Code ResCode `json:"code"`
	Message interface{} `json:"message"`
	Data interface{} `json:"data"`
}

func ResponseError(c *gin.Context,code ResCode){
	rd := &ResponseData{
		Code:    code,
		Message: code.getMsg(),
		Data:    nil,
	}
	c.JSON(http.StatusOK,rd)
	return
}
func ResponseErrorWithMsg(c *gin.Context,code ResCode,msg interface{}){
	rd := &ResponseData{
		Code:    code,
		Message: msg,
		Data:    nil,
	}
	c.JSON(http.StatusOK,rd)
	return
}

func ResponseSuccess(c *gin.Context,data interface{}){
	rd := &ResponseData{
		Code:    CodeSuccess,
		Message: CodeSuccess.getMsg(),
		Data:    data,
	}
	c.JSON(http.StatusOK,rd)
	return
}