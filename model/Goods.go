package model

import (
	"github.com/jinzhu/gorm"
)

type Goods struct {
	gorm.Model
	Name             string         `json:"name"form:"name"`     //商品名称
	Desc             string         `json:"desc"form:"desc"`     //商品介绍
	Price            float64        `json:"price"form:"price"`   //商品价格
	Unite            string         `json:"unite"form:"unite"`   //商品单位
	Image            string         `json:"image"form:"image"`   //商品图片
	Sales            uint           `json:"sales"form:"sales"`   //商品销量 新建商品时默认为0
	Stock            uint           `json:"stock"form:"stock"`   //商品库存
	Status           uint           `json:"status"form:"status"` //商品状态 上架或下架
	Category3Id      uint           `json:"category3Id"form:"category3_id"`
	GoodsImage       []*GoodsImage  `json:"goodsImage"gorm:"-"`
	IndexGoodsBanner []*GoodsBanner `json:"indexGoodsBanner"gorm:"-"`
	OrderGoods       []*OrderGoods  `json:"orderGoods"gorm:"-"`
}
