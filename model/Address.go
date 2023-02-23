package model

import (
	"github.com/jinzhu/gorm"
)

type Address struct {
	gorm.Model
	Receiver  string       `form:"Receiver"`  //收货人
	Addr      string       `form:"Addr"`      //收货地址
	Phone     string       `form:"Phone"`     //联系方式
	IsDefault bool         `form:"IsDefault"` //是否默认地址
	UserId    int          `form:"UserId"`    //用户id
	OrderInfo []*OrderInfo `gorm:"-"`
}
