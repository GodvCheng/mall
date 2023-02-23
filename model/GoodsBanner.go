package model

import (
	"github.com/jinzhu/gorm"
)

type GoodsBanner struct {
	gorm.Model
	GoodsSkuId int    `form:"GoodsSkuId"` //商品sku
	Image      string `form:"Image"`      //商品图片
	Index      int    `form:"Index"`      //展示顺序
}
