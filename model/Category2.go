package model

import "github.com/jinzhu/gorm"

type Category2 struct {
	gorm.Model
	Name        string `json:"name"form:"name"`
	Desc        string `json:"desc" form:"desc"`
	Image       string `json:"image"form:"image"`
	Category1Id uint   `json:"category1Id"form:"category1_id"`
}
