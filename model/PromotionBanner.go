package model

import "github.com/jinzhu/gorm"

type PromotionBanner struct {
	gorm.Model
	Name  string `json:"name"form:"name"`   //活动名称
	Url   string `json:"url"form:"url"`     //活动链接
	Image string `json:"image"form:"image"` //活动图片
	Sort  uint   `json:"sort"form:"sort"`   //展示顺序
}
