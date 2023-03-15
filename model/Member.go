package model

import "github.com/jinzhu/gorm"

type Member struct {
	gorm.Model
	Name      string       `json:"name"form:"name"`
	Username  string       `json:"username"form:"username"gorm:"unique"`
	Password  string       `json:"password"form:"password"`
	Phone     string       `json:"phone"form:"phone"gorm:"unique"`
	Email     string       `json:"email"form:"email"gorm:"unique"`
	Avatar    string       `json:"avatar"form:"avatar"`
	Age       uint         `json:"age"form:"age"`
	Sex       string       `json:"sex"form:"sex"`
	Status    uint         `json:"status"form:"status"`
	Address   []*Address   `json:"address"gorm:"-"`
	OrderInfo []*OrderInfo `json:"orderInfo"gorm:"-"`
}
