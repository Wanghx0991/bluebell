package routes

import (
	"bluebell/logger"
	"github.com/gin-gonic/gin"
	"net/http"
)

func SetupRouter() *gin.Engine {
	r := gin.New()
	r.Use(logger.GinLogger(),logger.GinRecovery(true))
	r.GET("/", func(context *gin.Context) {
		context.JSON(http.StatusOK,gin.H{
			"router_setup":"fine",
		})
	})
	return r
}
