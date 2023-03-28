package service

import (
	"mall/api/service/impl"
	"mall/model"
)

type GoodsService interface {
	GetCategories() ([]*model.Category1, error)
	GetBannerImg(goodsId int) ([]string, error)
	GetGoodsListByCategory2Id(category2Id int) ([]*model.Goods, error)
	GetGoodsInfo(id int) (*model.Goods, error)
	GetCategory2Info(id int) (*model.Category2, error)
}

func NewGoodsService() GoodsService {
	return &impl.GoodsServiceImpl{}
}
