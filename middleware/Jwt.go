package middleware

import (
	"github.com/gin-gonic/gin"
	"mall/common"
	"mall/util"
	"time"
)

//JWT token验证中间件
func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		var code = 200
		var data interface{}
		token := c.GetHeader("Authorization")
		if token == "" { //token为空
			code = 404
		} else {
			claims, err := util.ParseToken(token)
			if err != nil { //解析出错
				code = common.ErrorAuthCheckTokenFail
			} else if time.Now().Unix() > claims.ExpiresAt { //超时
				code = common.ErrorAuthCheckTokenTimeout
			}
		}
		if code != common.SUCCESS {
			c.JSON(200, gin.H{
				"status": code,
				"msg":    common.GetMsg(code),
				"data":   data,
			})
			c.Abort()
			return
		}
		c.Next()
	}
}

//JWTAdmin token验证中间件
func JWTAdmin() gin.HandlerFunc {
	return func(c *gin.Context) {
		var code int
		var data interface{}
		token := c.GetHeader("Authorization")
		if token == "" { //token为空
			code = common.InvalidParams
		} else {
			claims, err := util.ParseToken(token)
			if err != nil {
				code = common.ErrorAuthCheckTokenFail
			} else if time.Now().Unix() > claims.ExpiresAt {
				code = common.ErrorAuthCheckTokenTimeout
			} else if claims.Authority == 0 {
				code = common.ErrorAuthInsufficientAuthority
			}
		}
		if code != common.SUCCESS {
			c.JSON(200, gin.H{
				"status": code,
				"msg":    common.GetMsg(code),
				"data":   data,
			})
			c.Abort()
			return
		}
		c.Next()
	}
}
