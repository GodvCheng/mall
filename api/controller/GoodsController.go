package controller

import (
	"github.com/gin-gonic/gin"
	"mall/api/service"
	"mall/result"
	"strconv"
)

type response = result.Response

var goodsService = service.NewGoodsService()

func ListCategory(ctx *gin.Context) {
	categories, err := goodsService.GetCategories()
	if err != nil {
		result.Fail(ctx, response{
			Code:    509,
			Message: err.Error(),
			Data:    nil,
		})
	} else {
		result.OkWithData(ctx, gin.H{
			"categories": categories,
		})
	}
}
func GetBannerImg(ctx *gin.Context) {
	goodsId, _ := strconv.Atoi(ctx.Param("goodsId"))
	url, err := goodsService.GetBannerImg(goodsId)
	if err != nil {
		result.Fail(ctx, response{
			Code:    509,
			Message: err.Error(),
			Data:    nil,
		})
	} else {
		result.OkWithData(ctx, gin.H{
			"url": url,
		})
	}
}

func GoodListByCategory2(ctx *gin.Context) {
	category2Id, _ := strconv.Atoi(ctx.Param("category2Id"))
	goodsList, err := goodsService.GetGoodsListByCategory2Id(category2Id)
	if err != nil {
		result.Fail(ctx, response{
			Code:    509,
			Message: err.Error(),
			Data:    nil,
		})
	} else {
		result.OkWithData(ctx, gin.H{
			"goodsList": goodsList,
		})
	}
}

func GoodsInfo(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	goods, err := goodsService.GetGoodsInfo(id)
	if err != nil {
		result.Fail(ctx, response{
			Code:    509,
			Message: err.Error(),
			Data:    nil,
		})
	} else {
	}
	result.OkWithData(ctx, gin.H{
		"goods": goods,
	})
}

func Category2Info(ctx *gin.Context) {
	category2Id, _ := strconv.Atoi(ctx.Param("category2Id"))
	category2, err := goodsService.GetCategory2Info(category2Id)
	if err != nil {
		result.Fail(ctx, response{
			Code:    509,
			Message: err.Error(),
			Data:    nil,
		})
	} else {
		result.OkWithData(ctx, gin.H{
			"category2": category2,
		})
	}
}
