package service

import (
	"github.com/gin-gonic/gin"
	"mall/dao"
)

var ProductDao = dao.NewProductDao()

type ProductService struct {
}

type ProService interface {
	UpdateProduct(c *gin.Context)
}

func NewProductService() ProService {
	return &ProductService{}
}
func (p *ProductService) UpdateProduct(c *gin.Context) {

}
