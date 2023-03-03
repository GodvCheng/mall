package dao

import (
	"mall/model"
	"mall/model/dto"
)

type GDao interface {
	UpdateGoods(sku *model.GoodsSku) int
	CreateGoods(sku *model.GoodsSku) int
	ListGoods() []dto.GoodsSkuDto
	SearchGoods(id int) *model.GoodsSku
	//ListGoodsImg(id int) []*dto.GoodsSkuDto
	ListGoodsImg(id int) []string
	DeleteGoods(id int) int
	DisableGoods(id int) int
	EnableGoods(id int) int
	ListCategories() []*model.GoodsType
}

func NewGoodsDao() GDao {
	return &GoodsDao{}
}

type GoodsDao struct {
}

func (g *GoodsDao) ListCategories() (listCategories []*model.GoodsType) {
	Db.Find(&listCategories)
	return
}

func (g *GoodsDao) EnableGoods(id int) int {
	n := Db.Model(&model.GoodsSku{}).Where("id = ?", id).Update("status", 1).RowsAffected
	return int(n)
}

func (g *GoodsDao) DisableGoods(id int) int {
	n := Db.Model(&model.GoodsSku{}).Where("id = ?", id).Update("status", 0).RowsAffected
	return int(n)
}

func (g *GoodsDao) DeleteGoods(id int) int {
	num := Db.Delete(&model.GoodsSku{}, id).RowsAffected
	return int(num)
}

func (g *GoodsDao) ListGoodsImg(id int) (imgList []string) {
	Db.Raw(`SELECT i.image FROM goods_sku s join goods_image i on s.id = i.goods_sku_id where s.id = ?`, id).Pluck("image", &imgList)
	return
}

func (g *GoodsDao) SearchGoods(id int) *model.GoodsSku {
	var good model.GoodsSku
	Db.Where("id = ?", id).Find(&good)
	return &good
}

func (g *GoodsDao) ListGoods() (goodsList []dto.GoodsSkuDto) {
	//TODO 两张表连接查询，表中字段名重复，起别名，无法映射到结构体
	Db.Raw(`SELECT s.id,s.name,t.name ,s.price,s.sales,s.desc,s.image,s.stock,s.unite,s.status 
			FROM goods_sku s left JOIN goods_type t ON s.goods_type_id = t.id`).Scan(&goodsList)
	return
}

func (g *GoodsDao) CreateGoods(sku *model.GoodsSku) int {
	num := Db.Create(sku).RowsAffected
	return int(num)
}

func (g *GoodsDao) UpdateGoods(sku *model.GoodsSku) int {
	num := Db.Model(sku).Update(*sku).RowsAffected
	return int(num)
}
