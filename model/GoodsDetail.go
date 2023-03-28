package model

import "github.com/jinzhu/gorm"

type GoodsDetail struct {
	gorm.Model
	Detail string `form:"detail" json:"detail"`
}
