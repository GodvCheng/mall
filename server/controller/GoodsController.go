package controller

import (
	"github.com/gin-gonic/gin"
	"mall/model"
	"mall/result"
	"mall/server/service"
	"strconv"
)

var goodsService = service.NewGoodsService()

type response = result.Response

// CreateGoods 创建商品
func CreateGoods(c *gin.Context) {
	var goods model.Goods
	c.ShouldBind(&goods)
	err := goodsService.CreateGoods(&goods)
	if err != nil {
		result.Fail(c, response{
			Code:    result.ERROR,
			Message: err.Error(),
			Data:    nil,
		})
	} else {
		result.OkWithMsg(c, "商品信息创建成功")
	}
}

// ListGoods 商品列表
func ListGoods(c *gin.Context) {
	current, _ := strconv.Atoi(c.Param("current"))
	pageSize, _ := strconv.Atoi(c.Param("pageSize"))
	var goods model.Goods
	c.ShouldBind(&goods)
	goodsList, total, err := goodsService.ListGoods(&goods, current, pageSize)
	if err != nil {
		result.Fail(c, response{
			Code:    result.ERROR,
			Message: err.Error(),
			Data:    nil,
		})
	} else {
		result.OkWithData(c, gin.H{
			"goodsList": goodsList,
			"total":     total,
		})
	}

}

// DeleteGoods 删除商品
func DeleteGoods(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	err := goodsService.DeleteGoods(id)
	if err != nil {
		result.Fail(c, response{
			Code:    509,
			Message: err.Error(),
			Data:    nil,
		})
	} else {
		result.OkWithMsg(c, "删除商品成功")
	}
}

// UpdateGoods 更新商品
func UpdateGoods(c *gin.Context) {
	var goods model.Goods
	c.ShouldBind(&goods)
	err := goodsService.UpdateGoods(&goods)
	if err != nil {
		result.Fail(c, response{
			Code:    510,
			Message: err.Error(),
			Data:    nil,
		})
	} else {
		result.OkWithMsg(c, "更新商品信息成功")
	}
}
func ShowGoods(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	goods, err := goodsService.ShowGoods(id)
	if err != nil {
		result.Fail(c, response{
			Code:    509,
			Message: err.Error(),
			Data:    nil,
		})
	} else {
		result.OkWithData(c, gin.H{
			"goods": goods,
		})
	}
}

// ListGoodsImg 商品图片
func ListGoodsImg(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	imgList, err := goodsService.ListGoodsImg(id)
	if err != nil {
		result.Fail(c, response{
			Code:    509,
			Message: err.Error(),
			Data:    nil,
		})
	} else {
		result.OkWithData(c, gin.H{
			"imgList": imgList,
		})
	}
}

func DisableGoods(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	err := goodsService.DisableGoods(id)
	if err != nil {
		result.Fail(c, response{
			Code:    509,
			Message: err.Error(),
			Data:    nil,
		})
	} else {
		result.OkWithMsg(c, "商品下架成功")
	}
}

func EnableGoods(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	err := goodsService.EnableGoods(id)
	if err != nil {
		result.Fail(c, response{
			Code:    509,
			Message: err.Error(),
			Data:    nil,
		})
	} else {
		result.OkWithMsg(c, "商品上架成功")
	}
}

// ListCarousels TODO
func ListCarousels(c *gin.Context) {

}

func ListGoodsByCategory2(c *gin.Context) {
	category2Id, _ := strconv.Atoi(c.Param("category2Id"))
	current, _ := strconv.Atoi(c.Param("current"))
	pageSize, _ := strconv.Atoi(c.Param("pageSize"))
	goodsList, err := goodsService.ListGoodsByCategory2(category2Id, current, pageSize)
	if err != nil {
		result.Fail(c, response{
			Code:    509,
			Message: err.Error(),
			Data:    nil,
		})
	} else {
		result.OkWithData(c, gin.H{
			"goodsList": goodsList,
		})
	}
}

// ListCategory2 根据一级分类ID查询二级分类
func ListCategory2(c *gin.Context) {
	category1Id, _ := strconv.Atoi(c.Param("category1Id"))
	listCategory2, err := goodsService.ListCategory2(category1Id)
	if err != nil {
		result.Fail(c, response{
			Code:    509,
			Message: err.Error(),
			Data:    nil,
		})
	} else {
		result.OkWithData(c, gin.H{
			"listCategory2": listCategory2,
		})
	}
}

func Category2Info(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	category2, err := goodsService.Category2Info(id)
	if err != nil {
		result.Fail(c, response{
			Code:    509,
			Message: err.Error(),
			Data:    nil,
		})
	} else {
		result.OkWithData(c, gin.H{
			"category2": category2,
		})
	}
}

func UpdateCategory2(c *gin.Context) {
	var category2 model.Category2
	c.ShouldBind(&category2)
	err := goodsService.UpdateCategory2(&category2)
	if err != nil {
		result.Fail(c, response{
			Code:    509,
			Message: err.Error(),
			Data:    nil,
		})
	} else {
		result.OkWithMsg(c, "更新成功")
	}
}
func CreateCategory2(c *gin.Context) {
	var category2 model.Category2
	c.ShouldBind(&category2)
	err := goodsService.CreateCategory2(&category2)
	if err != nil {
		result.Fail(c, response{
			Code:    509,
			Message: err.Error(),
			Data:    nil,
		})
	} else {
		result.OkWithMsg(c, "创建二级分类成功")
	}
}

func DeleteCategory2(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	err := goodsService.DeleteCategory2(id)
	if err != nil {
		result.Fail(c, response{
			Code:    509,
			Message: err.Error(),
			Data:    nil,
		})
	} else {
		result.OkWithMsg(c, "删除二级分类成功")
	}

}

// ListCategory1 获取一级分类
func ListCategory1(c *gin.Context) {
	listCategory1, err := goodsService.ListCategory1()
	if err != nil {
		result.Fail(c, response{
			Code:    509,
			Message: err.Error(),
			Data:    nil,
		})
	} else {
		result.OkWithData(c, gin.H{
			"listCategory1": listCategory1,
		})
	}
}

func Category1Info(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	category1, err := goodsService.Category1Info(id)
	if err != nil {
		result.Fail(c, response{
			Code:    509,
			Message: err.Error(),
			Data:    nil,
		})
	} else {
		result.OkWithData(c, gin.H{
			"category1": category1,
		})
	}
}

func UpdateCategory1(c *gin.Context) {
	var category1 model.Category1
	c.ShouldBind(&category1)
	err := goodsService.UpdateCategory1(&category1)
	if err != nil {
		result.Fail(c, response{
			Code:    509,
			Message: err.Error(),
			Data:    nil,
		})
	} else {
		result.OkWithMsg(c, "更新成功")
	}
}

func CreateCategory1(c *gin.Context) {
	var category1 model.Category1
	c.ShouldBind(&category1)
	var err = goodsService.CreateCategory1(&category1)
	if err != nil {
		result.Fail(c, response{
			Code:    509,
			Message: err.Error(),
			Data:    nil,
		})
	} else {
		result.OkWithMsg(c, "创建一级分类成功")
	}
}

func DeleteCategory1(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	err := goodsService.DeleteCategory1(id)
	if err != nil {
		result.Fail(c, response{
			Code:    509,
			Message: err.Error(),
			Data:    nil,
		})
	} else {
		result.OkWithMsg(c, "删除一级分类成功")
	}
}

func UploadGoodsImages(c *gin.Context) {

}

func UploadGoodsBanner(c *gin.Context) {

}
