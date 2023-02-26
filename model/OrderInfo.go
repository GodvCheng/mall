package model

import (
	"github.com/jinzhu/gorm"
	"time"
)

type OrderInfo struct { //订单商品表
	gorm.Model
	OrderNo      string        `json:"orderNo"form:"order_no" gorm:"unique"` //订单编号
	UserId       int           `json:"userId"form:"user_id"`                 //用户id
	AddressId    int           `json:"addressId"form:"address_id"`           //地址ID
	PayMethod    int           `json:"payMethod"form:"pay_method"`           //付款方式
	TotalCount   int           `json:"totalCount"form:"total_count"`         //商品数量
	TotalPrice   float64       `json:"totalPrice"form:"total_price"`         //商品总价
	TransitPrice float64       `json:"transitPrice"form:"transit_price"`     //运费
	OrderStatus  int           `json:"orderStatus"form:"order_status"`       //订单状态
	TradeNo      string        `json:"tradeNo"form:"trade_no" gorm:"unique"` //支付编号
	CommentTime  time.Time     `json:"commentTime"form:"comment_time"`       //评论时间
	OrderGoods   []*OrderGoods `json:"orderGoods"gorm:"-"`
}
