package service

import (
	"mall/model"
	"mall/model/dto"
	"mall/service/impl"
)

type GService interface {
	UpdateGoods(sku *model.GoodsSku) error
	CreateGoods(sku *model.GoodsSku) error
	DeleteGoods(id int) error
	ListGoods() ([]dto.GoodsSkuDto, error)
	SearchGoods(id int) (*model.GoodsSku, error)
	ListGoodsImg(id int) (imgList []*dto.GoodsSkuDto, err error)
	ListGoodsType() ([]*model.GoodsType, error)
	DisableGoods(id int) error
	EnableGoods(id int) error
}

func NewGoodsService() GService {
	return &impl.GoodsService{}
}
