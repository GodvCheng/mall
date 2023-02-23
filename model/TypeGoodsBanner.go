package model

import "github.com/jinzhu/gorm"

type TypeGoodsBanner struct {
	gorm.Model
	GoodsTypeId int `form:"GoodsTypeId"` //商品类型
	GoodsSkuId  int `form:"GoodsSkuId"`  //商品sku
	DisplayType int `form:"DisplayType"` //展示类型 0代表文字，1代表图片
	Index       int `form:"Index"`       //展示顺序
}
