package model

import "github.com/jinzhu/gorm"

type TypeGoodsBanner struct {
	gorm.Model
	GoodsTypeId int //商品类型
	GoodsSkuId  int //商品sku
	DisplayType int //展示类型 0代表文字，1代表图片
	Index       int //展示顺序
}
