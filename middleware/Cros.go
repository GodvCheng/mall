package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

// Cors 中间件 跨域请求处理
func Cors(c *gin.Context) {
	method := c.Request.Method
	// 必须，接受指定域的请求，可以使用*不加以限制，但不安全
	//c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Origin", c.GetHeader("Origin"))
	fmt.Println(c.GetHeader("Origin"))
	// 必须，设置服务器支持的所有跨域请求的方法
	c.Header("Access-Control-Allow-Methods", "POST, GET, PUT, DELETE, OPTIONS")
	// 服务器支持的所有头信息字段，不限于浏览器在"预检"中请求的字段
	c.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Token")
	// 可选，设置XMLHttpRequest的响应对象能拿到的额外字段
	c.Header("Access-Control-Expose-Headers", "Access-Control-Allow-Headers, Token")
	// 可选，是否允许后续请求携带认证信息Cookie，该值只能是true，不需要则不设置
	c.Header("Access-Control-Allow-Credentials", "true")
	// 放行所有OPTIONS方法
	if method == "OPTIONS" {
		c.AbortWithStatus(http.StatusNoContent)
		return
	}
	c.Next()
}
