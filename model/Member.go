package model

import "github.com/jinzhu/gorm"

type Member struct {
	gorm.Model
	Name     string
	Username string
	Password string
	Avatar   string
	Age      int
	Sex      string
}
