package dao

import (
	"mall/model"
)

type GoodsDao interface {
	LoadCategories() []*model.Category1
	ListCategory2(category1Id int) []*model.Category2
	LoadBannerImg(goodsId int) []string
	ListGoodsByCategory2Id(category2Id int) []*model.Goods
	LoadGoodsInfo(id int) (*model.Goods, int)
	LoadCategory2Info(id int) (*model.Category2, int)
}

func NewGoodsDao() GoodsDao {
	return &GoodsDaoImpl{}
}

type GoodsDaoImpl struct {
}

func (g GoodsDaoImpl) LoadCategory2Info(id int) (category2 *model.Category2, n int) {
	category2 = new(model.Category2)
	db := Db.Where("id = ?", id).Find(category2)
	n = int(db.RowsAffected)
	return
}

func (g GoodsDaoImpl) LoadGoodsInfo(id int) (goods *model.Goods, n int) {
	goods = new(model.Goods)
	db := Db.Where("id = ? ", id).Find(goods)
	n = int(db.RowsAffected)
	return
}

func (g GoodsDaoImpl) ListGoodsByCategory2Id(category2Id int) (goodsList []*model.Goods) {
	Db.Where("category2_id = ? and status = ?", category2Id, 1).Find(&goodsList)
	return
}

func (g GoodsDaoImpl) ListCategory2(category1Id int) (category2List []*model.Category2) {
	Db.Where("category1_id = ?", category1Id).Find(&category2List)
	return
}

func (g GoodsDaoImpl) LoadCategories() (categories []*model.Category1) {
	Db.Find(&categories)
	return
}
func (g GoodsDaoImpl) LoadBannerImg(goodsId int) (url []string) {
	Db.Table("goods_banner").Where("goods_id = ?", goodsId).Pluck("image", &url)
	return
}
