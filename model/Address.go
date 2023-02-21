package model

import (
	"github.com/jinzhu/gorm"
)

type Address struct {
	gorm.Model
	Receiver  string       //收货人
	Addr      string       //收货地址
	Phone     string       //联系方式
	IsDefault bool         //是否默认地址
	UserId    int          //用户id
	OrderInfo []*OrderInfo `gorm:"-"`
}
