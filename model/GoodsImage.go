package model

import "github.com/jinzhu/gorm"

type GoodsImage struct {
	gorm.Model
	Image   string `json:"image" form:"image"`     //商品图片
	GoodsId uint   `json:"goodsId"form:"goods_id"` //商品sku
	Sort    uint   `json:"sort"form:"sort"`
}
