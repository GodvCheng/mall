package model

import "github.com/jinzhu/gorm"

type Category3 struct {
	gorm.Model
	Name        string `json:"name"form:"name"`
	Image       string `json:"image"form:"image"`
	Category2Id uint   `json:"category2Id"form:"category2_id"`
}
