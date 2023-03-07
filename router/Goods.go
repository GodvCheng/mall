package router

import (
	"github.com/gin-gonic/gin"
	"mall/controller"
	"mall/middleware"
)

func LoadProduct(r *gin.Engine) {

	group := r.Group("/goods")
	{
		//上传商品图片
		group.POST("/upload", controller.UploadImage)
		mAuth := group.Group("/")
		//普通权限
		mAuth.Use(middleware.JWT())
		{
			//创建商品
			mAuth.POST("/create", controller.CreateGoods)
			//更新商品
			mAuth.PUT("/update", controller.UpdateGoods)
			//查询所有商品
			mAuth.GET("/list/:current/:pageSize", controller.ListGoods)
			//展示商品详情
			mAuth.GET("/:id", controller.ShowGoods)
			//商品图片
			mAuth.GET("/imgs/:id", controller.ListGoodsImg)
			//商品分类
			mAuth.GET("/categories", controller.ListCategories)
			//轮播图
			mAuth.GET("/carousels", controller.ListCarousels)
			//下架商品
			mAuth.PUT("/disable/:id", controller.DisableGoods)
			//上架商品
			mAuth.PUT("/enable/:id", controller.EnableGoods)
			//商品种类信息
			mAuth.GET("/goodsType/:id", controller.GoodsTypeInfo)
			//更新商品种类信息
			mAuth.PUT("/goodsType/update", controller.UpdateGoodsType)
		}
	}
}
