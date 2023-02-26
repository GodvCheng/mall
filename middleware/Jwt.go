package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"mall/result"
	"mall/util"
	"time"
)

//JWT token验证中间件
func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		var code = result.SUCCESS
		var data interface{}
		token := c.GetHeader("token")
		if token == "" {
			code = result.ErrorAuth
		} else {
			claims, err := util.ParseToken(token)
			if err != nil { //解析出错
				code = result.ErrorAuthCheckTokenFail
			} else if time.Now().Unix() > claims.ExpiresAt { //超时
				code = result.ErrorAuthCheckTokenTimeout
			}
		}
		if code != result.SUCCESS {
			c.JSON(result.ERROR, gin.H{
				"code":    code,
				"message": result.GetMsg(code),
				"data":    data,
			})
			c.Abort()
			return
		}
		c.Set("token", token)
		c.Next()
	}
}

//JWTAdmin token验证中间件
func JWTAdmin() gin.HandlerFunc {
	return func(c *gin.Context) {
		var code = result.SUCCESS
		var data interface{}
		token := c.GetHeader("token")
		fmt.Println(token)
		if token == "" {
			code = result.ErrorAuth
		} else {
			claims, err := util.ParseToken(token)
			if err != nil {
				code = result.ErrorAuthCheckTokenFail
			} else if time.Now().Unix() > claims.ExpiresAt {
				code = result.ErrorAuthCheckTokenTimeout //过期
			} else if claims.Authority == 0 { //权限不足
				code = result.ErrorAuthInsufficientAuthority
			}
		}
		if code != result.SUCCESS {
			c.JSON(result.ERROR, gin.H{
				"code":    code,
				"message": result.GetMsg(code),
				"data":    data,
			})
			c.Abort()
			return
		}
		c.Set("token", token)
		c.Next() //放行
	}
}
