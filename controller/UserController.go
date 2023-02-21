package controller

import (
	"github.com/gin-gonic/gin"
	"mall/model"
	"mall/service"
	"net/http"
)

var UService = service.NewUService()

func UserRegister(c *gin.Context) {
	UService.UserRegister(c)
}
func UserLogin(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	var user = UService.UserLogin(username, password)
	c.JSON(http.StatusOK, &user)
}

func UserUpdate(c *gin.Context) {
	var user model.User
	c.ShouldBind(&user)
	UService.UpdateUser(&user)
}

func UploadAvatar(c *gin.Context) {

}

func SendEmail(c *gin.Context) {

}

func ValidEmail(c *gin.Context) {

}
