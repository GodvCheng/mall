package router

import (
	"github.com/gin-gonic/gin"
	"mall/api/controller"
)

func LoadGoods(r *gin.Engine) {
	group := r.Group("/api/goods")

	{
		group.GET("/category", controller.ListCategory)
		group.GET("/banner/:goodsId", controller.GetBannerImg)
		group.GET("/goodsList/:category2Id", controller.GoodListByCategory2)
		group.GET("/:id", controller.GoodsInfo)
		group.GET("/category2Info/:category2Id", controller.Category2Info)
	}

	{
		group.POST("addToCart")
	}
}
