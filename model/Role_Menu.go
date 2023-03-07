package model

import "github.com/jinzhu/gorm"

type RoleMenu struct {
	gorm.Model
	RoleId uint `json:"roleId" form:"role_id"`
	MenuId uint `json:"menuId" form:"menu_id"`
}
