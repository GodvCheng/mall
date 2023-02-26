package model

import (
	"github.com/jinzhu/gorm"
)

type Dict struct {
	gorm.Model
	Name  string `json:"name" form:"name"`
	Value string `json:"value" form:"value"`
	Type  string `json:"type" form:"type"`
}
