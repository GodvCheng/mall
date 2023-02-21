package controller

import (
	"github.com/gin-gonic/gin"
	"mall/common"
	"mall/service"
)

var UService = service.NewUService()

func UserRegister(c *gin.Context) {
	UService.UserRegister(c)
}
func UserLogin(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	var user = UService.UserLogin(username, password)
	common.NewDataResult(user)
}

func UserUpdate(c *gin.Context) {

}

func UploadAvatar(c *gin.Context) {

}

func SendEmail(c *gin.Context) {

}

func ValidEmail(c *gin.Context) {

}
