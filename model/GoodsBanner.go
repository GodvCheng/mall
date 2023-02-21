package model

import "github.com/jinzhu/gorm"

type GoodsBanner struct {
	gorm.Model
	GoodsSkuId int    //商品sku
	Image      string //商品图片
	Index      int    //展示顺序
}
