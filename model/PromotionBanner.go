package model

import "github.com/jinzhu/gorm"

type PromotionBanner struct {
	gorm.Model
	Name  string `form:"Name"`  //活动名称
	Url   string `form:"Url"`   //活动链接
	Image string `form:"Image"` //活动图片
	Index int    `form:"Index"` //展示顺序
}
