package model

import "github.com/jinzhu/gorm"

type Category2 struct {
	gorm.Model
	Name          string       `json:"name"form:"name"`
	Image         string       `json:"image"form:"image"`
	Category1Id   uint         `json:"category1Id"form:"category1_id"`
	Category3List []*Category3 `gorm:"-"`
}
