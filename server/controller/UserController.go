package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"mall/model"
	"mall/result"
	"mall/server/service"
	"mall/server/util"
	"strconv"
)

var userService = service.NewUserService()

func UserRegister(c *gin.Context) {
	var user model.User
	c.ShouldBind(&user)
	err := userService.ManagerRegister(&user)
	if err != nil {
		result.Fail(c, response{
			Code:    result.ErrorFailEncryption,
			Message: err.Error(),
			Data:    nil,
		})
	} else {
		result.OkWithMsg(c, "注册成功")
	}
	//c.Request.Header.Del("Authorization")
}

func UserLogin(c *gin.Context) {
	var user model.User
	c.ShouldBind(&user)
	username := user.Username
	password := user.Password
	token, err := userService.UserLogin(username, password)
	if err != nil || token == "" {
		result.Fail(c, response{
			Code:    result.ErrorAuthToken,
			Message: err.Error(),
			Data:    nil,
		})
	} else {
		result.OkWithData(c, gin.H{
			"token": token,
		})
	}

}

func GetUserInfo(c *gin.Context) {
	//从拦截器中取出token
	token := c.MustGet("token")
	user, err := userService.GetUserInfo(token.(string))
	if err != nil {
		result.Fail(c, response{
			Code:    result.ErrorAuthCheckTokenFail,
			Message: err.Error(),
			Data:    nil,
		})
	} else {
		result.OkWithData(c, gin.H{
			"user": user,
		})
	}
}

func UserUpdate(c *gin.Context) {
	var user model.User
	c.ShouldBind(&user)
	fmt.Println(user)
	err := userService.UpdateUser(&user)
	if err != nil {
		result.Fail(c, response{
			Code:    509,
			Message: err.Error(),
			Data:    nil,
		})
	} else {
		result.OkWithMsg(c, "修改成功")
	}
}
func DisableUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	err := userService.DisableUser(id)
	if err != nil {
		result.Fail(c, response{
			Code:    509,
			Message: err.Error(),
			Data:    nil,
		})
	} else {
		result.OkWithMsg(c, "用户禁用成功")
	}
}
func EnableUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	err := userService.EnableUser(id)
	if err != nil {
		result.Fail(c, response{
			Code:    509,
			Message: err.Error(),
			Data:    nil,
		})
	} else {
		result.OkWithMsg(c, "用户启用成功")
	}
}

func Logout(c *gin.Context) {
	token := c.GetHeader("token")
	if token != "" {
		c.Request.Header.Del("token")
		util.Rdb.Del(util.Ctx, "token")
		result.OkWithMsg(c, "退出成功")
	}
}

func UserList(c *gin.Context) {
	userList, err := userService.UserList()
	if err != nil {
		result.Fail(c, response{
			Code:    509,
			Message: err.Error(),
			Data:    nil,
		})
	} else {
		result.OkWithData(c, gin.H{
			"userList": userList,
		})
	}
}

func GetRoles(c *gin.Context) {
	token := c.GetHeader("token")
	roles, err := userService.GetRoles(token)
	if err != nil {
		result.Fail(c, response{
			Code:    509,
			Message: err.Error(),
			Data:    nil,
		})
	} else {
		result.OkWithData(c, gin.H{
			"roles": roles,
		})
	}
}

func AdminGetUserInfo(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	user, err := userService.AdminGetUserInfo(id)
	if err != nil {
		result.Fail(c, response{
			Code:    509,
			Message: err.Error(),
			Data:    nil,
		})
	} else {
		result.OkWithData(c, gin.H{
			"user": user,
		})
	}
}
func GetProfile(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	user, err := userService.GetProfile(id)
	if err != nil {
		result.Fail(c, response{
			Code:    509,
			Message: err.Error(),
			Data:    nil,
		})
	} else {
		result.OkWithData(c, gin.H{
			"user": user,
		})
	}
}
func SendEmail(c *gin.Context) {

}

func ValidEmail(c *gin.Context) {

}
