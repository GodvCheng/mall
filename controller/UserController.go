package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"mall/model"
	"mall/service"
	"net/http"
	"strconv"
)

var UService = service.NewUService()

func UserRegister(c *gin.Context) {
	var user model.User
	c.ShouldBind(&user)
	UService.ManagerRegister(&user)
	//c.Request.Header.Del("Authorization")
}
func UserLogin(c *gin.Context) {
	var user model.User
	c.ShouldBind(&user)
	username := user.Username
	password := user.Password
	token := UService.UserLogin(username, password)
	c.Header("token", token)
	// todo 注意删除这段代码
	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})

}

func GetUserInfo(c *gin.Context) {
	token := c.Request.Header.Get("token")
	fmt.Println("token = ", token)
	//token, _ := c.Cookie("token")
	fmt.Println(token)
	user := UService.GetUserInfo(token)
	c.JSON(http.StatusOK, user)
}

func UserUpdate(c *gin.Context) {
	var user model.User
	c.ShouldBind(&user)
	id, _ := strconv.Atoi(c.Param("id"))
	user.ID = uint(id)
	UService.UpdateUser(&user)
}

func UploadAvatar(c *gin.Context) {

}

func SendEmail(c *gin.Context) {

}

func ValidEmail(c *gin.Context) {

}
