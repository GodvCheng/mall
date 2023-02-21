package model

import "github.com/jinzhu/gorm"

type GoodsImage struct {
	gorm.Model
	Image      string //商品图片
	GoodsSkuId int    //商品sku
}
