package model

import (
	"github.com/jinzhu/gorm"
)

type Address struct {
	gorm.Model
	Receiver  string       `json:"receiver"form:"receiver"`    //收货人
	Addr      string       `json:"addr"form:"addr"`            //收货地址
	Phone     string       `json:"phone"form:"phone"`          //联系方式
	IsDefault bool         `json:"isDefault"form:"is_default"` //是否默认地址
	UserId    uint         `json:"userId"form:"user_id"`       //用户id
	OrderInfo []*OrderInfo `json:"orderInfo"gorm:"-"`
}
