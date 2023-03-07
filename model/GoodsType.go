package model

import "github.com/jinzhu/gorm"

type GoodsType struct {
	gorm.Model
	Name                 string         `json:"name"form:"name"`   // 类型名称
	Image                string         `json:"image"form:"image"` //类型图片
	Logo                 string         `json:"logo"form:"logo"`   //logo
	GoodsSpu             []*GoodsSpu    `json:"goodsSpu"gorm:"-"`
	IndexTypeGoodsBanner []*GoodsBanner `json:"indexTypeGoodsBanner"gorm:"-"`
}
