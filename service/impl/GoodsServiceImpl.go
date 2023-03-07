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

func (g *GoodsService) UpdateGoodsType(goodsType *model.GoodsType) (err error) {
	n := GoodsDao.UpdateGoodsType(goodsType)
	if n == 0 {
		err = errors.New("更新商品种类失败")
	}
	return
}

func (g *GoodsService) GoodsTypeInfo(id int) (goodsType *model.GoodsType, err error) {
	goodsType = GoodsDao.GoodsTypeInfo(id)
	if reflect.DeepEqual(goodsType, model.GoodsType{}) {
		return nil, errors.New("获取商品种类详情失败")
	}
	return
}

func (g *GoodsService) ListCategories() (listCategories []*model.GoodsType, err error) {
	listCategories = GoodsDao.ListCategories()
	if len(listCategories) == 0 {
		return nil, errors.New("查询商品分类失败")
	}
	return
}

func (g *GoodsService) EnableGoods(id int) (err error) {
	n := GoodsDao.EnableGoods(id)
	if n == 0 {
		err = errors.New("商品上架失败")
	}
	return nil
}

func (g *GoodsService) DisableGoods(id int) (err error) {
	n := GoodsDao.DisableGoods(id)
	if n == 0 {
		err = errors.New("商品下架失败")
	}
	return
}

func (g *GoodsService) DeleteGoods(id int) (err error) {
	num := GoodsDao.DeleteGoods(id)
	if num == 0 {
		err = errors.New("删除商品sku失败")
	}
	return
}

func (g *GoodsService) ListGoodsImg(id int) (imgList []string, err error) {
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

func (g *GoodsService) ListGoods(current, pageSize int) (goodsList []*dto.GoodsSkuDto, total int, err error) {
	goodsList, total = GoodsDao.ListGoods(current, pageSize)
	if len(goodsList) == 0 {
		return nil, 0, errors.New("分页查询商品列表失败")
	}
	return
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
