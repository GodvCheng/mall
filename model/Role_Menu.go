package model

import "github.com/jinzhu/gorm"

type RoleMenu struct {
	gorm.Model
	RoleId int `json:"roleId" form:"role_id"`
	MenuId int `json:"menuId" form:"menu_id"`
}
