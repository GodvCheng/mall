package service

import (
	"mall/model"
	"mall/model/dto"
	"mall/server/service/impl"
)

type GoodsService interface {
	CreateGoods(goods *model.Goods) error
	DeleteGoods(id int) error
	UpdateGoods(goods *model.Goods) error
	ShowGoods(id int) (*dto.GoodsDto, error)
	ListGoods(goods *model.Goods, current, pageSize int) ([]*dto.GoodsDto, int, error)
	ListGoodsImg(id int) (imgList []string, err error)
	DisableGoods(id int) error
	EnableGoods(id int) error
	ListGoodsByCategory2(category2Id, current, pageSize int) ([]*model.Goods, error)

	AddGoodsDetail(detail *model.GoodsDetail) error
	EditGoodsDetail(detail *model.GoodsDetail) error
	DeleteGoodsDetail(id int) error

	CreateCategory2(category2 *model.Category2) error
	UpdateCategory2(category2 *model.Category2) error
	ListCategory2(category1Id int) ([]*model.Category2, error)
	Category2Info(id int) (*model.Category2, error)
	DeleteCategory2(id int) error

	CreateCategory1(category1 *model.Category1) error
	UpdateCategory1(category1 *model.Category1) error
	ListCategory1() ([]*model.Category1, error)
	Category1Info(id int) (*model.Category1, error)
	DeleteCategory1(id int) error
}

func NewGoodsService() GoodsService {
	return &impl.GoodsServiceImpl{}
}
