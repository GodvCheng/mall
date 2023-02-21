package model

import "github.com/jinzhu/gorm"

type GoodsType struct {
	gorm.Model
	Name                 string         // 类型名称
	Image                string         //类型图片
	Logo                 string         //logo
	GoodsSku             []*GoodsSku    `gorm:"-"`
	IndexTypeGoodsBanner []*GoodsBanner `gorm:"-"`
}
