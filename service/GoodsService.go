package service

import (
	"mall/model"
	"mall/model/dto"
	"mall/service/impl"
)

type GoodsService interface {
	UpdateGoods(sku *model.Goods) error
	CreateGoods(sku *model.Goods) error
	ShowGoods(id int) (*dto.GoodsDto, error)
	DeleteGoods(id int) error
	ListGoods(goods *model.Goods, current, pageSize int) ([]*dto.GoodsDto, int, error)
	ListGoodsImg(id int) (imgList []string, err error)
	DisableGoods(id int) error
	EnableGoods(id int) error

	CreateCategory3(category3 *model.Category3) error
	UpdateCategory3(Category3 *model.Category3) error
	ListCategory3(category2Id int) ([]*model.Category3, error)
	Category3Info(id int) (*model.Category3, error)
	DeleteCategory3(id int) error

	CreateCategory2(category2 *model.Category2) error
	UpdateCategory2(category2 *model.Category2) error
	ListCategory2(category1Id int) ([]*model.Category2, error)
	Category2Info(id int) (*model.Category2, error)
	DeleteCategory2(id int) error

	CreateCategory1(category1 *model.Category1) error
	UpdateCategory1(category1 *model.Category1) error
	ListCategory1(current, pageSize int) ([]*model.Category1, int, error)
	Category1Info(id int) (*model.Category1, error)
	DeleteCategory1(id int) error
}

func NewGoodsService() GoodsService {
	return &impl.GoodsServiceImpl{}
}
