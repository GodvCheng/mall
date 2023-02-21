package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"mall/dao"
	"mall/router"
)

func main() {
	//初始化数据库连接
	dao.InitDb()
	defer dao.CloseDb()
	r := gin.Default()
	router.LoadUser(r)
	router.LoadProduct(r)
	r.Run(":8888")
	fmt.Println("go web start!")
}
