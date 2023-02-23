package model

import "github.com/jinzhu/gorm"

type GoodsSpu struct {
	gorm.Model
	Name     string      `form:"Name"`   // 商品名称
	Detail   string      `form:"Detail"` //详细描述
	GoodsSku []*GoodsSku `gorm:"-"`
}
