package main

import (
	"fmt"
	"mall/dao"
	"mall/router"
)

func main() {
	//初始化数据库连接
	dao.InitDb()
	defer dao.CloseDb()
	r := router.NewRouter()
	r.Run(":8888")
	fmt.Println("go web start!")
}
