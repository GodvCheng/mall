package model

import (
	"github.com/jinzhu/gorm"
)

type GoodsBanner struct {
	gorm.Model
	GoodsSkuId int    `json:"goodsSkuId"form:"goods_sku_id"` //商品sku
	Image      string `json:"image"form:"image"`             //商品图片
	Sort       int    `json:"sort"form:"sort"`               //展示顺序
}
