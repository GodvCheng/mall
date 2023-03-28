package main

import (
	"mall/api/dao"
	"mall/api/initilize"
)

func main() {
	dao.InitDao()
	defer dao.CloseDb()
	r := initilize.Router()
	r.Run(":9999")
}
