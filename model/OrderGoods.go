package model

import "github.com/jinzhu/gorm"

type OrderGoods struct { //订单商品表
	gorm.Model
	OrderInfoId int    //订单id
	GoodsSkuId  int    //商品id
	Count       int    //商品数量
	Price       int    //商品价格
	Comment     string //订单评论
}
