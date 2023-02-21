package model

import "github.com/jinzhu/gorm"

type GoodsSpu struct {
	gorm.Model
	Name     string      // 商品名称
	Detail   string      //详细描述
	GoodsSku []*GoodsSku `gorm:"-"`
}
