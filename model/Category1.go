package model

import "github.com/jinzhu/gorm"

// Category1 一级分类
type Category1 struct {
	gorm.Model
	Name          string       `json:"name"form:"name"`
	Category2List []*Category2 `json:"children"gorm:"-"`
}
