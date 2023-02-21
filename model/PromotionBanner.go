package model

import "github.com/jinzhu/gorm"

type PromotionBanner struct {
	gorm.Model
	Name  string //活动名称
	Url   string //活动链接
	Image string //活动图片
	Index int    //展示顺序
}
