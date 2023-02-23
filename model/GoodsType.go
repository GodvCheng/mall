package model

import "github.com/jinzhu/gorm"

type GoodsType struct {
	gorm.Model
	Name                 string         `form:"Name"`  // 类型名称
	Image                string         `form:"Image"` //类型图片
	Logo                 string         `form:"Logo"`  //logo
	GoodsSku             []*GoodsSku    `gorm:"-"`
	IndexTypeGoodsBanner []*GoodsBanner `gorm:"-"`
}
