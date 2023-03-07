package service

import (
	"mall/model"
	"mall/model/dto"
	"mall/service/impl"
)

type GoodsService interface {
	UpdateGoods(sku *model.GoodsSku) error
	CreateGoods(sku *model.GoodsSku) error
	DeleteGoods(id int) error
	ListGoods(current, pageSize int) ([]*dto.GoodsSkuDto, int, error)
	SearchGoods(id int) (*model.GoodsSku, error)
	ListGoodsImg(id int) (imgList []string, err error)
	DisableGoods(id int) error
	EnableGoods(id int) error
	ListCategories() ([]*model.GoodsType, error)
	GoodsTypeInfo(id int) (*model.GoodsType, error)
	UpdateGoodsType(goodsType *model.GoodsType) error
}

func NewGoodsService() GoodsService {
	return &impl.GoodsService{}
}
