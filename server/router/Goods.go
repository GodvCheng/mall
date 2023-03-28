package router

import (
	"github.com/gin-gonic/gin"
	"mall/server/controller"
	"mall/server/middleware"
)

func LoadGoods(r *gin.Engine) {

	group := r.Group("/goods")
	{
		//上传商品图片
		group.POST("/upload", controller.UploadImage)
		group.POST("/uploadPict")
		mAuth := group.Group("/")
		//普通权限
		mAuth.Use(middleware.JWT())
		{
			//创建商品
			mAuth.POST("/create", controller.CreateGoods)
			//更新商品
			mAuth.PUT("/update", controller.UpdateGoods)
			//查询所有商品
			mAuth.POST("/list/:current/:pageSize", controller.ListGoods)
			//根据二级分类Id查询商品
			mAuth.GET("/list/:category2Id/:current/:pageSize", controller.ListGoodsByCategory2)
			// 展示商品详情
			mAuth.GET("/:id", controller.ShowGoods)
			//商品图片
			mAuth.GET("/imgs/:id", controller.ListGoodsImg)
			//TODO 轮播图
			mAuth.GET("/carousels", controller.ListCarousels)
			//下架商品
			mAuth.PUT("/disable/:id", controller.DisableGoods)
			//上架商品
			mAuth.PUT("/enable/:id", controller.EnableGoods)
			//删除商品
			mAuth.DELETE("/delGoods/:id", controller.DeleteGoods)

			//创建二级分类
			mAuth.POST("/category2", controller.CreateCategory2)
			//商品二级分类列表
			mAuth.GET("/category2List/:category1Id", controller.ListCategory2)
			//商品二级分类信息
			mAuth.GET("/category2/:id", controller.Category2Info)
			//更新商品二级分类
			mAuth.PUT("/category2/update", controller.UpdateCategory2)
			//删除二级分类
			mAuth.DELETE("/delCategory2/:id", controller.DeleteCategory2)

			//创建一级分类
			mAuth.POST("/category1", controller.CreateCategory1)
			//商品一级分类列表
			mAuth.GET("/category1List", controller.ListCategory1)
			//商品一级分类信息
			mAuth.GET("/category1/:id", controller.Category1Info)
			//更新商品一级分类
			mAuth.PUT("/category1/update", controller.UpdateCategory1)
			//删除一级分类
			mAuth.DELETE("/delCategory1/:id", controller.DeleteCategory1)
		}
	}
}
