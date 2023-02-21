package model

import (
	"github.com/jinzhu/gorm"
	"time"
)

type GoodsSku struct {
	gorm.Model
	GoodsSpuId           int                //商品spu
	GoodsTypeId          int                //商品所属种类
	Name                 string             //商品名称
	Desc                 string             //商品介绍
	Price                float64            //商品价格
	Unite                string             //商品单位
	Image                string             //商品图片
	Sales                int                //商品销量
	Stock                int                //商品库存
	Status               int                //商品状态
	Time                 time.Time          //添加时间
	GoodsImage           []*GoodsImage      `gorm:"-"`
	IndexGoodsBanner     []*GoodsBanner     `gorm:"-"`
	IndexTypeGoodsBanner []*TypeGoodsBanner `gorm:"-"`
	OrderGoods           []*OrderGoods      `gorm:"-"`
}
