package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	result2 "mall/result"
	"mall/server/util"
	"time"
)

//JWT token验证中间件
func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		var code = result2.SUCCESS
		var data interface{}
		token := c.GetHeader("token")
		if token == "" {
			code = result2.ErrorAuthIsNull
		} else {
			parseToken, err := util.ParseToken(token)
			if err != nil || parseToken.Username == "" {
				fmt.Errorf("token解析失败%v", err)
			}
			username := parseToken.Username
			//从redis中取出token
			redisToken := util.Rdb.Get(util.Ctx, "token:"+username)
			if redisToken.Val() != token {
				code = result2.ErrorUserNotLogin
			}
			claims, err := util.ParseToken(token)
			if err != nil { //解析出错
				code = result2.ErrorAuthCheckTokenFail
				fmt.Println(err)
			} else if time.Now().Unix() > claims.ExpiresAt { //超时
				code = result2.ErrorAuthCheckTokenTimeout
			}
		}
		if code != result2.SUCCESS {
			c.JSON(result2.ERROR, gin.H{
				"code":    code,
				"message": result2.GetMsg(code),
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
		var code = result2.SUCCESS
		var data interface{}
		token := c.GetHeader("token")
		if token == "" {
			code = result2.ErrorAuthIsNull
		} else {
			parseToken, _ := util.ParseToken(token)
			username := parseToken.Username
			//从redis中取出token
			redisToken := util.Rdb.Get(util.Ctx, "token:"+username)
			if redisToken.Val() != token {
				code = result2.ErrorUserNotLogin
			}
			claims, err := util.ParseToken(token)
			if err != nil {
				code = result2.ErrorAuthCheckTokenFail
			} else if time.Now().Unix() > claims.ExpiresAt {
				code = result2.ErrorAuthCheckTokenTimeout //过期
			} else if claims.Authority == 0 { //权限不足
				code = result2.ErrorAuthInsufficientAuthority
			}
		}
		if code != result2.SUCCESS {
			c.JSON(result2.ERROR, gin.H{
				"code":    code,
				"message": result2.GetMsg(code),
				"data":    data,
			})
			c.Abort()
			return
		}
		c.Set("token", token)
		c.Next() //放行
	}
}
