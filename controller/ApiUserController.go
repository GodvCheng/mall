package controller

import (
	"github.com/gin-gonic/gin"
	"mall/model"
	"mall/service"
)

var ApiUService = service.NewApiUService()

func ApiUserRegister(c *gin.Context) {
	var user model.User
	c.ShouldBind(&user)
	ApiUService.UserRegister(&user)
}
