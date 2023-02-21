package router

import (
	"github.com/gin-gonic/gin"
	"mall/controller"
	"net/http"
)

func LoadUser(r *gin.Engine) {
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "success"})
	})

	//用户操作
	group := r.Group("/user")
	{
		group.POST("/register", controller.UserRegister)
		group.POST("/login", controller.UserLogin)
	}
}
