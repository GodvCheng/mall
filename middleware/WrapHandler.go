package middleware

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"mall/common"
	"mall/util"
	"net/http"
)

type EndpointFunc func(c *gin.Context) (any, error)

//中间件 定义统一处理器
func wrapHandler(handler EndpointFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		// 处理跨域请求
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		Cors(c)
		//从header 获取token,
		var token = c.GetHeader("token")
		//解析token,如果token不存在，或者 解析错误,则返回 没权限
		userId, err := util.TokenHandle(token)
		if err != nil {
			fmt.Errorf(err.Error())
			c.JSON(http.StatusBadRequest, gin.H{"code": "ERROR", "message": "token is invalid"})
			return
		}
		//解析后 userId 放入上下文中,后面的service服务，可以获取该用户信息
		c.Set("userId", userId)
		data, err := handler(c)
		if err != nil {
			var systemErr *common.Error
			if errors.As(err, &systemErr) {
				c.JSON(systemErr.Status, common.NewErrorResult(systemErr.Code, systemErr.Msg))
			} else {
				c.JSON(http.StatusBadRequest, gin.H{"code": "ERROR", "message": err.Error()})
			}
			return
		}
		c.JSON(http.StatusOK, gin.H{"code": "SUCCESS", "data": data})
	}
}
