package router

import (
	"github.com/gin-gonic/gin"
	"mall/controller"
	"net/http"
)

// NewRouter 路由配置
func NewRouter() *gin.Engine {
	//生成默认路由
	r := gin.Default()
	
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "SUCCESS",
		})
	})

	//用户操作
	group := r.Group("/user")
	{
		group.POST("/register", controller.UserRegister)
		group.POST("/login", controller.UserLogin)

		//商品操作
		group.GET("products", controller.ListProducts)
		group.GET("product/:id", controller.ShowProduct)
		group.POST("products", controller.SearchProducts)
		group.GET("imgs/:id", controller.ListProductImg)   //商品图片
		group.GET("categories", controller.ListCategories) //商品分类
		group.GET("carousels", controller.ListCarousels)   //轮播图
	}

	return r
}
