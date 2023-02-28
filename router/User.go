package router

import (
	"github.com/gin-gonic/gin"
	"mall/controller"
	"mall/middleware"
	"net/http"
)

func LoadUser(r *gin.Engine) {
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "ok",
		})
	})

	group := r.Group("/user")
	{
		//登录
		group.POST("/login", controller.UserLogin)
		//获取个人信息 普通权限
		group.POST("/info", middleware.JWT(), controller.GetUserInfo)
		//退出登录
		group.POST("/logout", middleware.JWT(), controller.Logout)
		//头像上传 TODO
		group.POST("/upload", middleware.JWT(), controller.UploadAvatar)
		adminAuth := group.Group("/")
		//超级管理员权限
		adminAuth.Use(middleware.JWTAdmin())
		{
			//后台人员注册
			adminAuth.POST("/register", controller.UserRegister)
			//禁用用户
			adminAuth.PUT("/disable/:id", controller.DisableUser)
			//启用用户
			adminAuth.PUT("/enable/:id", controller.EnableUser)
			//更新
			adminAuth.PUT("/update/:id", controller.UserUpdate)
			//获取用户列表
			adminAuth.GET("/userList", controller.UserList)
		}
	}
}
