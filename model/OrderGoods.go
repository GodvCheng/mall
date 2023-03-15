package model

import "github.com/jinzhu/gorm"

type OrderGoods struct { //订单商品表
	gorm.Model
	OrderInfoId uint    `json:"orderInfoId"form:"order_info_id"` //订单id
	GoodsId     uint    `json:"goodsId"form:"goods_id"`          //商品id
	Count       uint    `json:"count"form:"count"`               //商品数量
	Price       float64 `json:"price"form:"price"`               //商品价格
	Comment     string  `json:"comment"form:"comment"`           //订单评论
}
