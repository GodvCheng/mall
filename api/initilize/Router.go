package initilize

import (
	"github.com/gin-gonic/gin"
	"mall/api/middleware"
	"mall/api/router"
)

func Router() *gin.Engine {
	r := gin.Default()
	r.Use(middleware.Cors())
	router.LoadMember(r)
	router.LoadGoods(r)
	router.LoadCart(r)
	return r
}
