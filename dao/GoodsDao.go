package dao

import (
	"mall/model"
	"mall/model/dto"
)

type GoodsDao interface {
	CreateGoods(goods *model.Goods) int
	UpdateGoods(goods *model.Goods) int
	ListGoods(current, pageSize int) ([]*dto.GoodsDto, int)
	ShowGoods(id int) *dto.GoodsDto
	ListGoodsImg(id int) []string
	DeleteGoods(id int) int
	DisableGoods(id int) int
	EnableGoods(id int) int

	CreateCategory3(category3 *model.Category3) int
	UpdateCategory3(Category3 *model.Category3) int
	ListCategory3(category2Id int) []*model.Category3
	Category3Info(id int) *model.Category3
	DeleteCategory3(id int) int

	CreateCategory2(category2 *model.Category2) int
	UpdateCategory2(category2 *model.Category2) int
	ListCategory2(category1Id int) []*model.Category2
	Category2Info(id int) *model.Category2
	DeleteCategory2(id int) int

	CreateCategory1(category1 *model.Category1) int
	UpdateCategory1(category1 *model.Category1) int
	ListCategory1(current, pageSize int) ([]*model.Category1, int)
	Category1Info(id int) *model.Category1
	DeleteCategory1(id int) int
}

func NewGoodsDao() GoodsDao {
	return &GoodsDaoImpl{}
}

type GoodsDaoImpl struct {
}

func (g *GoodsDaoImpl) DeleteCategory3(id int) int {
	n := Db.Delete(&model.Category3{}, id).RowsAffected
	return int(n)
}

func (g *GoodsDaoImpl) DeleteCategory2(id int) int {
	n := Db.Delete(&model.Category2{}, id).RowsAffected
	return int(n)
}

func (g *GoodsDaoImpl) DeleteCategory1(id int) int {
	n := Db.Delete(&model.Category1{}, id).RowsAffected
	return int(n)
}

func (g *GoodsDaoImpl) CreateCategory1(category1 *model.Category1) int {
	n := Db.Create(category1).RowsAffected
	return int(n)
}

func (g *GoodsDaoImpl) UpdateCategory1(category1 *model.Category1) int {
	n := Db.Model(category1).Updates(*category1).RowsAffected
	return int(n)
}

func (g *GoodsDaoImpl) ListCategory1(current, pageSize int) (listCategory1 []*model.Category1, total int) {
	Db.Scopes(Paginate(current, pageSize)).Find(&listCategory1)
	Db.Table("category1").Count(&total)
	return
}

func (g *GoodsDaoImpl) Category1Info(id int) *model.Category1 {
	var category1 model.Category1
	Db.Where("id = ? ", id).Find(&category1)
	return &category1
}

func (g *GoodsDaoImpl) ListCategory2(category1Id int) (category2List []*model.Category2) {
	Db.Where("category1_id = ?", category1Id).Find(&category2List)
	return
}

func (g *GoodsDaoImpl) Category2Info(id int) *model.Category2 {
	var category2 model.Category2
	Db.Where("id = ?", id).Find(&category2)
	return &category2
}

func (g *GoodsDaoImpl) UpdateCategory2(category2 *model.Category2) int {
	num := Db.Model(category2).Updates(*category2).RowsAffected
	return int(num)
}

func (g *GoodsDaoImpl) CreateCategory2(category2 *model.Category2) int {
	n := Db.Create(category2).RowsAffected
	return int(n)
}

func (g *GoodsDaoImpl) CreateCategory3(category3 *model.Category3) int {
	n := Db.Create(category3).RowsAffected
	return int(n)
}

func (g *GoodsDaoImpl) UpdateCategory3(category3 *model.Category3) int {
	num := Db.Model(category3).Updates(*category3).RowsAffected
	return int(num)
}

func (g *GoodsDaoImpl) Category3Info(id int) *model.Category3 {
	var category3 model.Category3
	Db.Where("id = ?", id).Find(&category3)
	return &category3
}

func (g *GoodsDaoImpl) ListCategory3(category2Id int) (listCategory3 []*model.Category3) {
	Db.Where("category2_id = ?", category2Id).Find(&listCategory3)
	return
}

func (g *GoodsDaoImpl) ListGoods(current, pageSize int) (goodsList []*dto.GoodsDto, total int) {
	Db.Table("goods").Scopes(Paginate(current, pageSize)).Find(&goodsList)
	Db.Table("goods").Count(&total)
	return
}
func (g *GoodsDaoImpl) ShowGoods(id int) *dto.GoodsDto {
	var goods dto.GoodsDto
	Db.Table("goods").Where("id = ?", id).Find(&goods)
	return &goods
}

func (g *GoodsDaoImpl) EnableGoods(id int) int {
	n := Db.Model(&model.Goods{}).Where("id = ?", id).Update("status", 1).RowsAffected
	return int(n)
}

func (g *GoodsDaoImpl) DisableGoods(id int) int {
	n := Db.Model(&model.Goods{}).Where("id = ?", id).Update("status", 0).RowsAffected
	return int(n)
}

func (g *GoodsDaoImpl) DeleteGoods(id int) int {
	num := Db.Delete(&model.Goods{}, id).RowsAffected
	return int(num)
}

func (g *GoodsDaoImpl) ListGoodsImg(id int) (imgList []string) {
	Db.Raw(`SELECT i.image FROM goods s join goods_image i on s.id = i.goods_id where s.id = ?`, id).Pluck("image", &imgList)
	return
}

func (g *GoodsDaoImpl) CreateGoods(goods *model.Goods) int {
	num := Db.Create(goods).RowsAffected
	return int(num)
}

func (g *GoodsDaoImpl) UpdateGoods(goods *model.Goods) int {
	num := Db.Model(goods).Update(*goods).RowsAffected
	return int(num)
}
