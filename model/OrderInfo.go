package model

import (
	"github.com/jinzhu/gorm"
	"time"
)

type OrderInfo struct { //订单商品表
	gorm.Model
	OrderNum     string        `form:"OrderNum" gorm:"unique"` //订单编号
	UserId       int           `form:"OrderNum"`               //用户id
	AddressId    int           `form:"UserId"`                 //地址ID
	PayMethod    int           `form:"AddressId"`              //付款方式
	TotalCount   int           `form:"PayMethod"`              //商品数量
	TotalPrice   int           `form:"TotalCount"`             //商品总价
	TransitPrice int           `form:"TotalPrice"`             //运费
	OrderStatus  int           `form:"TransitPrice"`           //订单状态
	TradeNo      string        `form:"TradeNo" gorm:"unique"`  //支付编号
	Time         time.Time     `form:"Time"`                   //评论时间
	OrderGoods   []*OrderGoods `gorm:"-"`
}
