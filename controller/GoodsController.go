package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"mall/model"
	"mall/result"
	"mall/service"
	"strconv"
)

var GoodsService = service.NewGoodsService()

type Response = result.Response

// CreateGoods 创建商品
func CreateGoods(c *gin.Context) {
	var goods model.GoodsSku
	c.ShouldBind(&goods)
	err := GoodsService.CreateGoods(&goods)
	if err != nil {
		result.Fail(c, Response{
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
	goodsList, total, err := GoodsService.ListGoods(current, pageSize)
	if err != nil {
		result.Fail(c, Response{
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

// ShowGoods 商品详情
func ShowGoods(c *gin.Context) {
	paramId := c.Param("id")
	id, _ := strconv.Atoi(paramId)
	goods, err := GoodsService.SearchGoods(id)
	if err != nil {
		result.Fail(c, Response{
			Code:    509,
			Message: err.Error(),
			Data:    nil,
		})
	} else {
		result.OkWithData(c, goods)
	}
}

// DeleteGoods 删除商品
func DeleteGoods(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	err := GoodsService.DeleteGoods(id)
	if err != nil {
		result.Fail(c, Response{
			Code:    509,
			Message: err.Error(),
			Data:    nil,
		})
	} else {
		result.OkWithMsg(c, "删除商品sku成功")
	}
}

// UpdateGoods 更新商品
func UpdateGoods(c *gin.Context) {
	var goodsSku model.GoodsSku
	c.ShouldBind(&goodsSku)
	err := GoodsService.UpdateGoods(&goodsSku)
	if err != nil {
		result.Fail(c, Response{
			Code:    510,
			Message: err.Error(),
			Data:    nil,
		})
	} else {
		result.OkWithMsg(c, "更新商品信息成功")
	}
}

// ListGoodsImg 商品图片
func ListGoodsImg(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	imgList, err := GoodsService.ListGoodsImg(id)
	if err != nil {
		result.Fail(c, Response{
			Code:    509,
			Message: err.Error(),
			Data:    nil,
		})
	} else {
		result.OkWithData(c, imgList)
	}
}

func DisableGoods(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	err := GoodsService.DisableGoods(id)
	if err != nil {
		result.Fail(c, Response{
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
	err := GoodsService.EnableGoods(id)
	if err != nil {
		result.Fail(c, Response{
			Code:    509,
			Message: err.Error(),
			Data:    nil,
		})
	} else {
		result.OkWithMsg(c, "商品上架成功")
	}
}

// ListCategories 商品种类
func ListCategories(c *gin.Context) {
	listCategories, err := GoodsService.ListCategories()
	if err != nil {
		result.Fail(c, Response{
			Code:    509,
			Message: err.Error(),
			Data:    nil,
		})
	} else {
		result.OkWithData(c, listCategories)
	}
}

func ListCarousels(c *gin.Context) {

}

func GoodsTypeInfo(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	goodsType, err := GoodsService.GoodsTypeInfo(id)
	if err != nil {
		result.Fail(c, Response{
			Code:    509,
			Message: err.Error(),
			Data:    nil,
		})
	} else {
		result.OkWithData(c, goodsType)
	}
}

func UpdateGoodsType(c *gin.Context) {
	var goodsType model.GoodsType
	c.ShouldBind(&goodsType)
	fmt.Println(goodsType)
	err := GoodsService.UpdateGoodsType(&goodsType)
	if err != nil {
		result.Fail(c, Response{
			Code:    509,
			Message: err.Error(),
			Data:    nil,
		})
	} else {
		result.OkWithMsg(c, "更新成功")
	}
}
