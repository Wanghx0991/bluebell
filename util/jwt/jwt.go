package jwt

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"time"
)

//MyClaims 自定义声明结构体并内嵌jwt.StandardClaims
//jwt包自带的 jwt.StandardClaims只包含了官方字段
//如果我们需要额外记录一个username的字段,需要自定义结构体
// 如果想要保存更多信息,都可以添加到这个结构体中

const TokenExpireDuration = time.Hour * 2
var sercet = []byte("winter iscoming")

type MyClaims struct {
	UserId int64 `json:"user_id"`
	Username string `json:"username"`
	jwt.StandardClaims
}

//生成Token JWT
func GentToken(userid int64, username string) (string, error) {
	//创建一个我们自己的声明
	t := MyClaims{
		UserId:     userid    ,
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(TokenExpireDuration).Unix(),
			Issuer: "bluebell",
		},
	}
	//使用指定的签名方法创建签名对象
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,t)
	// 使用指定的secret签名并获得完整的编码后的字符串token
	return token.SignedString(sercet)
}

//解析Token
func ParseToken(tokenstring string)(*MyClaims,error)  {
	//解析token
	var mc = new(MyClaims)
	token,err := jwt.ParseWithClaims(tokenstring,mc, func(token *jwt.Token) (interface{}, error) {
		return sercet,nil
	})
	if err != nil{
		return nil, err
	}
	// token转换成struct
	if  token.Valid{//校验token
		return mc,nil
	}
	return nil, errors.New("invalid token")
}