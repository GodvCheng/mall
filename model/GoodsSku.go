package model

import (
	"github.com/jinzhu/gorm"
	"time"
)

type GoodsSku struct {
	gorm.Model
	GoodsSpuId           int                `form:"GoodsSpuId"`  //商品spu
	GoodsTypeId          int                `form:"GoodsTypeId"` //商品所属种类
	Name                 string             `form:"Name"`        //商品名称
	Desc                 string             `form:"Desc"`        //商品介绍
	Price                float64            `form:"Price"`       //商品价格
	Unite                string             `form:"Unite"`       //商品单位
	Image                string             `form:"Image"`       //商品图片
	Sales                int                `form:"Sales"`       //商品销量
	Stock                int                `form:"Stock"`       //商品库存
	Status               int                `form:"Status"`      //商品状态
	Time                 time.Time          `form:"Time"`        //添加时间
	GoodsImage           []*GoodsImage      `gorm:"-"`
	IndexGoodsBanner     []*GoodsBanner     `gorm:"-"`
	IndexTypeGoodsBanner []*TypeGoodsBanner `gorm:"-"`
	OrderGoods           []*OrderGoods      `gorm:"-"`
}
