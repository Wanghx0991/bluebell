package models

//定义注册请求的参数
type ParamSignUp struct {
	UserName string `json:"user_name" binding:"required"`
	Password string	`json:"password" binding:"required"`
	RePassword string `json:"re_password" binding:"required,eqfield=Password"`
}

//定义登录请求的参数
type ParamLogin struct{
	UserName string `json:"user_name" binding:"required"`
	Password string	`json:"password" binding:"required"`
}

