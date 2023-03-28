package dto

type CartDto struct {
	MemberId uint    `json:"memberId"form:"memberId"`
	Id       uint    `json:"id"form:"id"`
	Name     string  `json:"name"form:"name"`
	Image    string  `json:"image"form:"image"`
	Desc     string  `json:"desc"form:"desc"`
	Number   uint    `json:"number"form:"number"`
	Color    string  `json:"color"form:"color"`
	Price    float64 `json:"price"form:"price"`
}
