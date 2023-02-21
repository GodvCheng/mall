package dao

type ProductDao struct {
}

type ProDao interface {
	UpdateProduct()
}

func (p *ProductDao) UpdateProduct() {

}

func NewProductDao() ProDao {
	return &ProductDao{}
}
