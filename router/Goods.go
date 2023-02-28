package router

import (
	"github.com/gin-gonic/gin"
	"mall/controller"
	"mall/middleware"
)

func LoadProduct(r *gin.Engine) {

	group := r.Group("/goods")
	{
		mAuth := group.Group("/")
		//普通权限
		mAuth.Use(middleware.JWT())
		{
			//创建商品
			mAuth.POST("/create", controller.CreateGoods)
			//更新商品
			mAuth.PUT("/update/:id", controller.UpdateGoods)
			//查询所有商品
			mAuth.GET("/list", controller.ListGoods)
			//展示某个商品详情
			mAuth.GET("/:id", controller.ShowGoods)
			//根据条件搜索商品
			mAuth.POST("/goods", controller.SearchGoods)
			mAuth.GET("/imgs/:id", controller.ListGoodsImg)     //商品图片
			mAuth.GET("/categories", controller.ListCategories) //商品分类
			mAuth.GET("/carousels", controller.ListCarousels)   //轮播图
			mAuth.PUT("/disable/:id", controller.DisableGoods)
			mAuth.PUT("/enable/:id", controller.EnableGoods)
		}
	}
}
