package model

import "github.com/jinzhu/gorm"

type GoodsImage struct {
	gorm.Model
	Image   string `json:"image" form:"image"`     //商品图片
	GoodsId uint   `json:"goodsId"form:"Goods_id"` //商品sku
}
