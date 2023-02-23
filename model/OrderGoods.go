package model

import "github.com/jinzhu/gorm"

type OrderGoods struct { //订单商品表
	gorm.Model
	OrderInfoId int    `form:"OrderInfoId"` //订单id
	GoodsSkuId  int    `form:"GoodsSkuId"`  //商品id
	Count       int    `form:"Count"`       //商品数量
	Price       int    `form:"Price"`       //商品价格
	Comment     string `form:"Comment"`     //订单评论
}
