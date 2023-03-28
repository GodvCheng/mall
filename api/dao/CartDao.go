package dao

import (
	"mall/model/dto"
)

type CartDao interface {
	ShowCartInfo(goodsId string) *dto.CartDto
}

func NewCartDao() CartDao {
	return &CartDaoImpl{}
}

type CartDaoImpl struct {
}

func (c CartDaoImpl) ShowCartInfo(goodsId string) (cartInfo *dto.CartDto) {
	cartInfo = new(dto.CartDto)
	Db.Table("goods").Where("id = ?", goodsId).Find(cartInfo)
	return
}
