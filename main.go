package main

import (
	"github.com/gin-gonic/gin"
	"mall/dao"
	"mall/middleware"
	"mall/router"
)

func main() {
	//初始化数据库连接
	dao.InitDb()
	defer dao.CloseDb()
	r := gin.Default()
	//处理跨域请求
	r.Use(middleware.Cors())
	router.LoadUser(r)
	router.LoadProduct(r)
	router.LoadMember(r)
	r.Run(":8888")
}
