package model

import "github.com/jinzhu/gorm"

type Role struct {
	gorm.Model
	Name        string `json:"name" form:"name"`
	Description string `json:"description"  form:"description"`
	Key         string `json:"key" form:"key"`
}
