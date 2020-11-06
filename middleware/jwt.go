package middleware

import (
	"bluebell/util/jwt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func JWTAuthMiddleware() func(c *gin.Context) {
	return func(c *gin.Context) {
		//客户端携带token有3中方式: 1.放在请求头 2. 放在请求体 3. 放在url
		//假设token 放在header的 Authorization 中,并使用Bearer开头
		// Authorization: Bearer xxx.xxx.xx
		//具体的实现方式取决于业务情况
		authHeader := c.Request.Header.Get("Authorization")
		if authHeader == ""{
			c.JSON(http.StatusOK,gin.H{
				"code":2003,
				"msg": "请求头为空",
			})
			c.Abort()
			return
		}
		//按空格切分
		parts := strings.SplitN(authHeader," ",2)
		if !(len(parts) ==2 && parts[0] == "Bearer"){
			c.JSON(http.StatusOK,gin.H{
				"code":2004,
				"msg":"请求头格式有误",
			})
			c.Abort()
			return
		}
		// 解析token
		// parts[1] 是获取到的tokenstring,使用之前先用jwt函数解析
		mc, err := jwt.ParseToken(parts[1])
		if err != nil{
			c.JSON(http.StatusOK,gin.H{
				"code":2005,
				"msg":"无效的token",
			})
			c.Abort()
			return
		}
		//将当前请求的额username信息保存到请求的上下文c中
		c.Set("username",mc.Username)
		c.Next() //后续的处理函数可以用c.get("username")来获取当前请求的用户信息
	}
}