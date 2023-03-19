package impl

import (
	"errors"
	"mall/dao"
	"mall/model"
	"mall/model/dto"
	"reflect"
)

var GoodsDao = dao.NewGoodsDao()

type GoodsServiceImpl struct {
}

func (g *GoodsServiceImpl) DeleteCategory3(id int) (err error) {
	n := GoodsDao.DeleteCategory3(id)
	if n == 0 {
		err = errors.New("删除三级分类失败")
	}
	return
}

func (g *GoodsServiceImpl) DeleteCategory2(id int) (err error) {
	n := GoodsDao.DeleteCategory2(id)
	if n == 0 {
		err = errors.New("删除二级分类失败")
	}
	return
}

func (g *GoodsServiceImpl) DeleteCategory1(id int) (err error) {
	n := GoodsDao.DeleteCategory1(id)
	if n == 0 {
		err = errors.New("删除一级分类失败")
	}
	return
}

func (g *GoodsServiceImpl) ListCategory1(current, pageSize int) (category1List []*model.Category1, total int, err error) {
	category1List, total = GoodsDao.ListCategory1(current, pageSize)
	for _, category1 := range category1List {
		category2List := GoodsDao.ListCategory2(int(category1.ID))
		for _, category2 := range category2List {
			if category1.ID == category2.Category1Id {
				category1.Category2List = category2List
				////获取三级分类
				//category3List := GoodsDao.ListCategory3(int(category2.ID))
				//for _, category3 := range category3List {
				//	if category3.Category2Id == category2.ID {
				//		category2.Category3List = category3List
				//	}
				//}
			}
		}
	}
	if len(category1List) == 0 {
		return nil, 0, errors.New("商品一级分类为空")
	}
	return
}

func (g *GoodsServiceImpl) ListCategory2(category1Id int) (category2List []*model.Category2, err error) {
	category2List = GoodsDao.ListCategory2(category1Id)
	for _, category2 := range category2List {
		category3List := GoodsDao.ListCategory3(int(category2.ID))
		for _, category3 := range category3List {
			if category3.Category2Id == category2.ID {
				category2.Category3List = category3List
			}
		}
	}
	if len(category2List) == 0 {
		return nil, errors.New("商品二级分类为空")
	}
	return
}

func (g *GoodsServiceImpl) ListCategory3(category2Id int) (listCategories []*model.Category3, err error) {
	listCategories = GoodsDao.ListCategory3(category2Id)
	if len(listCategories) == 0 {
		return nil, errors.New("商品三级分类为空")
	}
	return
}

func (g *GoodsServiceImpl) Category1Info(id int) (category1 *model.Category1, err error) {
	category1 = GoodsDao.Category1Info(id)
	if reflect.DeepEqual(category1, model.Category1{}) {
		return nil, errors.New("获取商品一级分类详情失败")
	}
	return
}

func (g *GoodsServiceImpl) Category2Info(id int) (category2 *model.Category2, err error) {
	category2 = GoodsDao.Category2Info(id)
	if reflect.DeepEqual(category2, model.Category2{}) {
		return nil, errors.New("获取商品二级分类详情失败")
	}
	return
}

func (g *GoodsServiceImpl) Category3Info(id int) (category3 *model.Category3, err error) {
	category3 = GoodsDao.Category3Info(id)
	if reflect.DeepEqual(category3, model.Category3{}) {
		return nil, errors.New("获取商品三级分类详情失败")
	}
	return
}

func (g *GoodsServiceImpl) UpdateCategory1(category1 *model.Category1) (err error) {
	n := GoodsDao.UpdateCategory1(category1)
	if n == 0 {
		err = errors.New("更新一级分类失败")
	}
	return
}

func (g *GoodsServiceImpl) UpdateCategory2(category2 *model.Category2) (err error) {
	n := GoodsDao.UpdateCategory2(category2)
	if n == 0 {
		err = errors.New("更新二级种类失败")
	}
	return
}

func (g *GoodsServiceImpl) UpdateCategory3(category3 *model.Category3) (err error) {
	n := GoodsDao.UpdateCategory3(category3)
	if n == 0 {
		err = errors.New("更新三级种类失败")
	}
	return
}

func (g *GoodsServiceImpl) CreateCategory1(category1 *model.Category1) (err error) {
	n := GoodsDao.CreateCategory1(category1)
	if n == 0 {
		err = errors.New("创建一级分类失败")
	}
	return
}

func (g *GoodsServiceImpl) CreateCategory2(category2 *model.Category2) (err error) {
	var n int = GoodsDao.CreateCategory2(category2)
	if n == 0 {
		err = errors.New("创建二级分类失败")
	}
	return
}

func (g *GoodsServiceImpl) CreateCategory3(category3 *model.Category3) (err error) {
	n := GoodsDao.CreateCategory3(category3)
	if n == 0 {
		err = errors.New("创建三级分类失败")
	}
	return
}

func (g *GoodsServiceImpl) EnableGoods(id int) (err error) {
	n := GoodsDao.EnableGoods(id)
	if n == 0 {
		err = errors.New("商品上架失败")
	}
	return nil
}

func (g *GoodsServiceImpl) DisableGoods(id int) (err error) {
	n := GoodsDao.DisableGoods(id)
	if n == 0 {
		err = errors.New("商品下架失败")
	}
	return
}

func (g *GoodsServiceImpl) DeleteGoods(id int) (err error) {
	num := GoodsDao.DeleteGoods(id)
	if num == 0 {
		err = errors.New("删除商品失败")
	}
	return
}

func (g *GoodsServiceImpl) ListGoodsImg(id int) (imgList []string, err error) {
	imgList = GoodsDao.ListGoodsImg(id)
	if len(imgList) == 0 {
		return nil, errors.New("商品相关图片为空")
	}
	return imgList, nil
}

func (g *GoodsServiceImpl) ShowGoods(id int) (goods *dto.GoodsDto, err error) {
	goods = GoodsDao.ShowGoods(id)
	if reflect.DeepEqual(*goods, dto.GoodsDto{}) {
		err = errors.New("查询商品信息失败")
	}
	return
}

func (g *GoodsServiceImpl) ListGoods(goods *model.Goods, current, pageSize int) (goodsList []*dto.GoodsDto, total int, err error) {
	goodsList, total = GoodsDao.ListGoods(goods, current, pageSize)
	if len(goodsList) == 0 || total == 0 {
		return nil, 0, errors.New("商品列表为空")
	}
	return
}
func (g *GoodsServiceImpl) CreateGoods(sku *model.Goods) (err error) {
	n := GoodsDao.CreateGoods(sku)
	if n == 0 {
		err = errors.New("商品信息创建失败")
	}
	return
}

func (g *GoodsServiceImpl) UpdateGoods(sku *model.Goods) (err error) {
	n := GoodsDao.UpdateGoods(sku)
	if n == 0 {
		err = errors.New("商品信息更新失败")
	}
	return
}
