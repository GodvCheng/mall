package model

import "github.com/jinzhu/gorm"

type GoodsImage struct {
	gorm.Model
	Image      string `form:"Image"`      //商品图片
	GoodsSkuId int    `form:"GoodsSkuId"` //商品sku
}
