package router

import (
	"github.com/gin-gonic/gin"
	"mall/server/controller"
	"mall/server/middleware"
)

func LoadMember(c *gin.Engine) {
	group := c.Group("/member")
	group.Use(middleware.JWT())
	{
		group.GET("/list/:current/:pageSize", controller.MemberList)
	}
}
