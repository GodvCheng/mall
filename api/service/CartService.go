package service

import (
	"mall/api/service/impl"
	"mall/model/dto"
)

type CartService interface {
	AddToCart(cart *dto.CartDto) error
	ShowCartInfo(key string) ([]*dto.CartDto, error)
	DelFromCart(key, filed string) error
}

func NewCartService() CartService {
	return &impl.CartServiceImpl{}
}
