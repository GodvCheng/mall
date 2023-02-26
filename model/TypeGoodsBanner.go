package model

import "github.com/jinzhu/gorm"

type TypeGoodsBanner struct {
	gorm.Model
	GoodsTypeId int `json:"goodsTypeId"form:"goods_type_id"` //商品类型
	GoodsSkuId  int `json:"goodsSkuId"form:"goods_sku_id"`   //商品sku
	DisplayType int `json:"displayType"form:"display_type"`  //展示类型 0代表文字，1代表图片
	Sort        int `json:"sort"form:"sort"`                 //展示顺序
}
