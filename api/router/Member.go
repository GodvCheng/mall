package router

import "github.com/gin-gonic/gin"

func LoadMember(r *gin.Engine) {
	group := r.Group("/api/member")
	{
		group.POST("/login")
		group.POST("/register")
		group.POST("/forgetPass")
	}

	{
		group.POST("/getInfo")
		group.POST("/address")
		group.GET("/address")
		group.PUT("/address")
		group.DELETE("/address")
		group.GET("/getProfile")
		group.PUT("/editProfile")
		group.PUT("/editPass")
	}
}
