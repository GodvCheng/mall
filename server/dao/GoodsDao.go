package dao

import (
	"mall/model"
	"mall/model/dto"
	"strings"
)

type GoodsDao interface {
	CreateGoods(goods *model.Goods) int
	UpdateGoods(goods *model.Goods) int
	ListGoods(goods *model.Goods, current, pageSize int) ([]*dto.GoodsDto, int)
	ShowGoods(id int) *dto.GoodsDto
	ListGoodsImg(id int) []string
	DeleteGoods(id int) int
	DisableGoods(id int) int
	EnableGoods(id int) int
	CreateGoodsDetail(detail *model.GoodsDetail) int
	EditGoodsDetail(detail *model.GoodsDetail) int
	DeleteGoodsDetail(id int) int
	ListGoodsByCategory2(category2Id, current, pageSize int) []*model.Goods

	CreateCategory2(category2 *model.Category2) int
	UpdateCategory2(category2 *model.Category2) int
	ListCategory2(category1Id int) []*model.Category2
	Category2Info(id int) *model.Category2
	DeleteCategory2(id int) int

	CreateCategory1(category1 *model.Category1) int
	UpdateCategory1(category1 *model.Category1) int
	ListCategory1() []*model.Category1
	Category1Info(id int) *model.Category1
	DeleteCategory1(id int) int
}

func NewGoodsDao() GoodsDao {
	return &GoodsDaoImpl{}
}

type GoodsDaoImpl struct {
}

func (g *GoodsDaoImpl) ListGoodsByCategory2(category2Id, current, pageSize int) (goodsList []*model.Goods) {
	Db.Where("category2_id = ?", category2Id).Scopes(Paginate(current, pageSize)).Find(&goodsList)
	return
}

func (g *GoodsDaoImpl) EditGoodsDetail(detail *model.GoodsDetail) int {
	n := Db.Model(detail).Updates(*detail).RowsAffected
	return int(n)
}

func (g *GoodsDaoImpl) DeleteGoodsDetail(id int) int {
	n := Db.Delete(&model.GoodsDetail{}, id).RowsAffected
	return int(n)
}

func (g *GoodsDaoImpl) CreateGoodsDetail(detail *model.GoodsDetail) int {
	n := Db.Create(detail).RowsAffected
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

func (g *GoodsDaoImpl) ListCategory1() (listCategory1 []*model.Category1) {
	Db.Table("category1").Find(&listCategory1)
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

// ListGoods TODO
func (g *GoodsDaoImpl) ListGoods(goods *model.Goods, current, pageSize int) (goodsList []*dto.GoodsDto, total int) {
	switch {
	case strings.TrimSpace(goods.Name) != "" && goods.Status != 0:
		Db.Table("goods").Where("name like ? and status = ?", "%"+goods.Name+"%", goods.Status).Scopes(Paginate(current, pageSize)).Find(&goodsList)
		Db.Table("goods").Where("name like ? and status = ?", "%"+goods.Name+"%", goods.Status).Count(&total)
	case strings.TrimSpace(goods.Name) != "":
		Db.Table("goods").Where("name like ? ", "%"+goods.Name+"%").Scopes(Paginate(current, pageSize)).Find(&goodsList)
		Db.Table("goods").Where("name like ? ", "%"+goods.Name+"%").Count(&total)
	case goods.Status != 0:
		Db.Table("goods").Where("status = ?", goods.Status).Scopes(Paginate(current, pageSize)).Find(&goodsList)
		Db.Table("goods").Where("status = ?", goods.Status).Count(&total)
	default:
		Db.Table("goods").Scopes(Paginate(current, pageSize)).Find(&goodsList)
		Db.Table("goods").Count(&total)
	}
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
	n := Db.Model(&model.Goods{}).Where("id = ?", id).Update("status", 2).RowsAffected
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
