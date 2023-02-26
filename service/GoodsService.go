package service

import (
	"errors"
	"mall/dao"
	"mall/model"
	"mall/model/dto"
	"reflect"
)

var GoodsDao = dao.NewGoodsDao()

type GoodsService struct {
}

func (g *GoodsService) ListGoodsImg(id int) (imgList []*dto.GoodsSkuDto, err error) {
	imgList = GoodsDao.ListGoodsImg(id)
	if len(imgList) == 0 {
		return nil, errors.New("查询商品相关图片失败")
	}
	return imgList, nil
}

func (g *GoodsService) SearchGoods(id int) (*model.GoodsSku, error) {
	goods := GoodsDao.SearchGoods(id)
	if reflect.DeepEqual(*goods, model.GoodsSpu{}) {
		return nil, errors.New("未找到该商品")
	}
	return goods, nil
}

func (g *GoodsService) ListGoods() (goodsList []dto.GoodsSkuDto, err error) {
	goodsList = GoodsDao.ListGoods()
	if len(goodsList) == 0 {
		err = errors.New("查询商品列表失败")
		return nil, err
	}
	return goodsList, nil
}

type GService interface {
	UpdateGoods(sku *model.GoodsSku) error
	CreateGoods(sku *model.GoodsSku) error
	ListGoods() ([]dto.GoodsSkuDto, error)
	SearchGoods(id int) (*model.GoodsSku, error)
	ListGoodsImg(id int) (imgList []*dto.GoodsSkuDto, err error)
}

func NewGoodsService() GService {
	return &GoodsService{}
}

func (g *GoodsService) CreateGoods(sku *model.GoodsSku) (err error) {
	n := GoodsDao.CreateGoods(sku)
	if n == 0 {
		err = errors.New("商品信息创建失败")
	}
	return
}

func (g *GoodsService) UpdateGoods(sku *model.GoodsSku) (err error) {
	n := GoodsDao.UpdateGoods(sku)
	if n == 0 {
		err = errors.New("商品信息更新失败")
	}
	return
}
