package model

import (
	"github.com/jinzhu/gorm"
	"time"
)

type OrderInfo struct { //订单商品表
	gorm.Model
	OrderNum     string        `gorm:"unique"` //订单编号
	UserId       int           //用户id
	AddressId    int           //地址ID
	PayMethod    int           //付款方式
	TotalCount   int           //商品数量
	TotalPrice   int           //商品总价
	TransitPrice int           //运费
	OrderStatus  int           //订单状态
	TradeNo      string        `gorm:"unique"` //支付编号
	Time         time.Time     //评论时间
	OrderGoods   []*OrderGoods `gorm:"-"`
}
