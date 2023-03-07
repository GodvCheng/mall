package model

import "github.com/jinzhu/gorm"

type TypeGoodsBanner struct {
	gorm.Model
	GoodsTypeId uint `json:"goodsTypeId"form:"goods_type_id"` //商品类型
	GoodsSkuId  uint `json:"goodsSkuId"form:"goods_sku_id"`   //商品sku
	DisplayType uint `json:"displayType"form:"display_type"`  //展示类型 0代表文字，1代表图片
	Sort        uint `json:"sort"form:"sort"`                 //展示顺序
}
