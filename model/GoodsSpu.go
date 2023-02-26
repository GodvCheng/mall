package model

import "github.com/jinzhu/gorm"

type GoodsSpu struct {
	gorm.Model
	Name     string      `json:"name"form:"name"`     // 商品名称
	Detail   string      `json:"detail"form:"detail"` //详细描述
	GoodsSku []*GoodsSku `json:"goodsSku"gorm:"-"`
}
