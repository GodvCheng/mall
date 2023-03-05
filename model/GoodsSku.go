package model

import (
	"github.com/jinzhu/gorm"
)

type GoodsSku struct {
	gorm.Model
	GoodsSpuId           int                `json:"goodsSpuId"form:"goods_spu_id"`   //商品spuId
	GoodsTypeId          int                `json:"goodsTypeId"form:"goods_type_id"` //商品所属种类id
	Name                 string             `json:"name"form:"name"`                 //商品名称
	Desc                 string             `json:"desc"form:"desc"`                 //商品介绍
	Price                float64            `json:"price"form:"price"`               //商品价格
	Unite                string             `json:"unite"form:"unite"`               //商品单位
	Image                string             `json:"image"form:"image"`               //商品图片
	Sales                int                `json:"sales"form:"sales"`               //商品销量 新建商品时默认为0
	Stock                int                `json:"stock"form:"stock"`               //商品库存
	Status               int                `json:"status"form:"status"`             //商品状态 上架或下架
	GoodsImage           []*GoodsImage      `json:"goodsImage"gorm:"-"`
	IndexGoodsBanner     []*GoodsBanner     `json:"indexGoodsBanner"gorm:"-"`
	IndexTypeGoodsBanner []*TypeGoodsBanner `json:"indexTypeGoodsBanner"gorm:"-"`
	OrderGoods           []*OrderGoods      `json:"orderGoods"gorm:"-"`
}
