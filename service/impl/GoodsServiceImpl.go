package impl

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

func (g *GoodsService) EnableGoods(id int) (err error) {
	n := GoodsDao.EnableGoods(id)
	if n == 0 {
		return errors.New("商品上架失败")
	}
	return nil
}

func (g *GoodsService) DisableGoods(id int) (err error) {
	n := GoodsDao.DisableGoods(id)
	if n == 0 {
		return errors.New("商品下架失败")
	}
	return nil
}

func (g *GoodsService) DeleteGoods(id int) (err error) {
	num := GoodsDao.DeleteGoods(id)
	if num == 0 {
		return errors.New("删除商品sku失败")
	}
	return
}

func (g *GoodsService) ListGoodsType() (list []*model.GoodsType, err error) {
	list = GoodsDao.ListGoodsType()
	if len(list) == 0 {
		return nil, errors.New("查询商品分类列表失败")
	}
	return
}

func (g *GoodsService) ListGoodsImg(id int) (imgList []*dto.GoodsSkuDto, err error) {
	imgList = GoodsDao.ListGoodsImg(id)
	if len(imgList) == 0 {
		return nil, errors.New("查询商品相关图片失败")
	}
	return imgList, nil
}

func (g *GoodsService) SearchGoods(id int) (goods *model.GoodsSku, err error) {
	goods = GoodsDao.SearchGoods(id)
	if reflect.DeepEqual(*goods, model.GoodsSpu{}) {
		return nil, errors.New("未找到该商品")
	}
	return
}

func (g *GoodsService) ListGoods() (goodsList []dto.GoodsSkuDto, err error) {
	goodsList = GoodsDao.ListGoods()
	if len(goodsList) == 0 {
		err = errors.New("查询商品列表失败")
		return nil, err
	}
	return goodsList, nil
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
