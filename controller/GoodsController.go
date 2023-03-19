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
	var goods model.Goods
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
	var goods model.Goods
	c.ShouldBind(&goods)
	fmt.Println(goods)
	goodsList, total, err := GoodsService.ListGoods(&goods, current, pageSize)
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
		result.OkWithMsg(c, "删除商品成功")
	}
}

// UpdateGoods 更新商品
func UpdateGoods(c *gin.Context) {
	var goodsSku model.Goods
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
func ShowGoods(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	goods, err := GoodsService.ShowGoods(id)
	if err != nil {
		result.Fail(c, Response{
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
	imgList, err := GoodsService.ListGoodsImg(id)
	if err != nil {
		result.Fail(c, Response{
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

// ListCarousels TODO
func ListCarousels(c *gin.Context) {

}

func ListGoodsByCategory3(c *gin.Context) {
	strconv.Atoi(c.Param("category3Id"))
	strconv.Atoi(c.Param("current"))
	strconv.Atoi(c.Param("pageSize"))
}

// ListCategory3 根据二级分类ID查询三级分类
func ListCategory3(c *gin.Context) {
	category2Id, _ := strconv.Atoi(c.Param("category2Id"))
	listCategory3, err := GoodsService.ListCategory3(category2Id)
	if err != nil {
		result.Fail(c, Response{
			Code:    509,
			Message: err.Error(),
			Data:    nil,
		})
	} else {
		result.OkWithData(c, gin.H{
			"listCategory3": listCategory3,
		})
	}
}

// ListCategory2 根据一级分类ID查询二级分类
func ListCategory2(c *gin.Context) {
	category1Id, _ := strconv.Atoi(c.Param("category1Id"))
	listCategory2, err := GoodsService.ListCategory2(category1Id)
	if err != nil {
		result.Fail(c, Response{
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

// ListCategory1 获取一级分类
func ListCategory1(c *gin.Context) {
	current, _ := strconv.Atoi(c.Param("current"))
	pageSize, _ := strconv.Atoi(c.Param("pageSize"))
	listCategory1, total, err := GoodsService.ListCategory1(current, pageSize)
	if err != nil {
		result.Fail(c, Response{
			Code:    509,
			Message: err.Error(),
			Data:    nil,
		})
	} else {
		result.OkWithData(c, gin.H{
			"listCategory1": listCategory1,
			"total":         total,
		})
	}
}

func Category3Info(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	category3, err := GoodsService.Category3Info(id)
	if err != nil {
		result.Fail(c, Response{
			Code:    509,
			Message: err.Error(),
			Data:    nil,
		})
	} else {
		result.OkWithData(c, gin.H{
			"category3": category3,
		})
	}
}

func Category2Info(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	category2, err := GoodsService.Category2Info(id)
	if err != nil {
		result.Fail(c, Response{
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

func Category1Info(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	category1, err := GoodsService.Category1Info(id)
	if err != nil {
		result.Fail(c, Response{
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

// UpdateCategory3 更新三级分类
func UpdateCategory3(c *gin.Context) {
	var category3 model.Category3
	c.ShouldBind(&category3)
	err := GoodsService.UpdateCategory3(&category3)
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

func UpdateCategory2(c *gin.Context) {
	var category2 model.Category2
	c.ShouldBind(&category2)
	err := GoodsService.UpdateCategory2(&category2)
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

func UpdateCategory1(c *gin.Context) {
	var category1 model.Category1
	c.ShouldBind(&category1)
	err := GoodsService.UpdateCategory1(&category1)
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

func CreateCategory3(c *gin.Context) {
	var category3 model.Category3
	c.ShouldBind(&category3)
	var err = GoodsService.CreateCategory3(&category3)
	if err != nil {
		result.Fail(c, Response{
			Code:    509,
			Message: err.Error(),
			Data:    nil,
		})
	} else {
		result.OkWithMsg(c, "创建三级分类成功")
	}
}

func CreateCategory2(c *gin.Context) {
	var category2 model.Category2
	c.ShouldBind(&category2)
	err := GoodsService.CreateCategory2(&category2)
	if err != nil {
		result.Fail(c, Response{
			Code:    509,
			Message: err.Error(),
			Data:    nil,
		})
	} else {
		result.OkWithMsg(c, "创建二级分类成功")
	}
}

func CreateCategory1(c *gin.Context) {
	var category1 model.Category1
	c.ShouldBind(&category1)
	var err = GoodsService.CreateCategory1(&category1)
	if err != nil {
		result.Fail(c, Response{
			Code:    509,
			Message: err.Error(),
			Data:    nil,
		})
	} else {
		result.OkWithMsg(c, "创建一级分类成功")
	}
}

func DeleteCategory3(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	err := GoodsService.DeleteCategory3(id)
	if err != nil {
		result.Fail(c, Response{
			Code:    509,
			Message: err.Error(),
			Data:    nil,
		})
	} else {
		result.OkWithMsg(c, "删除三级分类成功")
	}
}
func DeleteCategory2(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	err := GoodsService.DeleteCategory2(id)
	if err != nil {
		result.Fail(c, Response{
			Code:    509,
			Message: err.Error(),
			Data:    nil,
		})
	} else {
		result.OkWithMsg(c, "删除二级分类成功")
	}

}
func DeleteCategory1(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	err := GoodsService.DeleteCategory1(id)
	if err != nil {
		result.Fail(c, Response{
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
