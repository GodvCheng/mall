package dao

import (
	"mall/model"
	"mall/model/dto"
)

type GoodsDao struct {
}

type GDao interface {
	UpdateGoods(sku *model.GoodsSku) int
	CreateGoods(sku *model.GoodsSku) int
	ListGoods() []dto.GoodsSkuDto
	SearchGoods(id int) *model.GoodsSku
	ListGoodsImg(id int) []*dto.GoodsSkuDto
}

func (g *GoodsDao) ListGoodsImg(id int) (imgList []*dto.GoodsSkuDto) {
	//TODO
	Db.Raw(`SELECT i.image FROM goods_sku s join goods_image i on s.id = i.goods_sku_id where s.id = ?`, id).Scan(&imgList)
	return imgList
}

func (g *GoodsDao) SearchGoods(id int) *model.GoodsSku {
	var good model.GoodsSku
	Db.Where("id = ?", id).Find(&good)
	return &good
}

func (g *GoodsDao) ListGoods() (goodsList []dto.GoodsSkuDto) {
	Db.Raw(`SELECT s.id,s.name,t.name goodsTypeName,s.price,s.sales,s.desc,s.image,s.stock,s.unite,s.status " +
		"FROM goods_sku s JOIN goods_type t ON s.goods_type_id = t.id`).Scan(&goodsList)
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

func NewGoodsDao() GDao {
	return &GoodsDao{}
}
