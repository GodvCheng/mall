package main

import (
	"mall/server/dao"
	"mall/server/initilize"
)

func main() {
	//初始化数据库连接
	dao.InitDb()
	defer dao.CloseDb()
	r := initilize.Router()
	r.Run(":8888")
}
