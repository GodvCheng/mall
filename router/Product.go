package router

import (
	"github.com/gin-gonic/gin"
	"mall/controller"
)

func LoadProduct(r *gin.Engine) {

	group := r.Group("/pro")
	{
		//商品操作
		group.GET("products", controller.ListProducts)
		group.GET("product/:id", controller.ShowProduct)
		group.POST("products", controller.SearchProducts)
		group.GET("imgs/:id", controller.ListProductImg)   //商品图片
		group.GET("categories", controller.ListCategories) //商品分类
		group.GET("carousels", controller.ListCarousels)   //轮播图
	}
}
