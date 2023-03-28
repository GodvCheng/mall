package model

import (
	"github.com/jinzhu/gorm"
)

type Goods struct {
	gorm.Model
	Name             string         `json:"name"form:"name"` //商品名称
	Desc             string         `json:"desc"form:"desc"` //商品介绍
	GoodsNo          string         `json:"goodsNo"gorm:"goods_no"`
	Price            float64        `json:"price"form:"price"`   //商品价格
	Unite            string         `json:"unite"form:"unite"`   //商品单位
	Image            string         `json:"image"form:"image"`   //商品图片
	Sales            uint           `json:"sales"form:"sales"`   //商品销量 新建商品时默认为0
	Stock            uint           `json:"stock"form:"stock"`   //商品库存
	Status           uint           `json:"status"form:"status"` //商品状态 1 上架或 2 下架
	Rate             uint           `json:"rate"form:"rate"`     //评分
	IsNew            uint           `json:"isNew"form:"is_new"`  //是否为新品 1 是 2 否
	Category2Id      uint           `json:"category2Id"form:"category2_id"`
	Detail           string         `json:"detail" gorm:"-"`
	GoodsImage       []*GoodsImage  `json:"goodsImage"gorm:"-"`
	IndexGoodsBanner []*GoodsBanner `json:"indexGoodsBanner"gorm:"-"`
	OrderGoods       []*OrderGoods  `json:"orderGoods"gorm:"-"`
}
