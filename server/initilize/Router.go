package initilize

import (
	"github.com/gin-gonic/gin"
	"mall/server/middleware"
	"mall/server/router"
)

func Router() *gin.Engine {
	r := gin.Default()
	//处理跨域请求
	r.Use(middleware.Cors())
	router.LoadUser(r)
	router.LoadGoods(r)
	router.LoadMember(r)
	return r
}
