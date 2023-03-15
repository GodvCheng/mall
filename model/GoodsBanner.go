package model

import (
	"github.com/jinzhu/gorm"
)

type GoodsBanner struct {
	gorm.Model
	GoodsId uint   `json:"goodsId"form:"goods_id"` //商品sku
	Image   string `json:"image"form:"image"`      //商品图片
	Sort    uint   `json:"sort"form:"sort"`        //展示顺序
}
