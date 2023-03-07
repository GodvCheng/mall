package model

import "github.com/jinzhu/gorm"

type GoodsImage struct {
	gorm.Model
	Image      string `json:"image" form:"image"`            //商品图片
	GoodsSkuId uint   `json:"goodsSkuId"form:"Goods_sku_id"` //商品sku
}
