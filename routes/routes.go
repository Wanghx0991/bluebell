package routes

import (
	"bluebell/controllers"
	"bluebell/logger"
	"bluebell/middleware"
	"github.com/gin-gonic/gin"
	"net/http"
)

func SetupRouter() *gin.Engine {
	r := gin.New()
	r.Use(logger.GinLogger(),logger.GinRecovery(true))
	//注册业务路由
	r.POST("/signup",controllers.SinUpHandler)
	r.POST("/login",controllers.LogInHandler)
	r.GET("/ping", middleware.JWTAuthMiddleware(),func(context *gin.Context) {
		//如果是登录的用户,判断请求头里是否存在 有效的jwt
		context.JSON(http.StatusOK,gin.H{
			"msg": "pong",
		})
	})
	r.NoRoute(func(context *gin.Context) {
		context.JSON(http.StatusOK,gin.H{
			"msg": "404Not found",
		})
	})
	return r
}
