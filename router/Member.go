package router

import (
	"github.com/gin-gonic/gin"
	"mall/controller"
	"mall/middleware"
)

func LoadMember(c *gin.Engine) {
	group := c.Group("/member")
	group.Use(middleware.JWT())
	{
		group.GET("/list/:current/:pageSize", controller.MemberList)
	}
}
