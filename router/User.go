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
		//获取个人信息
		//普通权限
		mAuth := group.Group("/")
		mAuth.Use(middleware.JWT())
		{
			mAuth.POST("/info", controller.GetUserInfo)
			//退出登录
			mAuth.POST("/logout", controller.Logout)
			//头像上传 TODO
			mAuth.POST("/upload", controller.UploadAvatar)
			//用户获取个人信息
			mAuth.GET("/profile/:id", controller.GetProfile)
		}
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
			adminAuth.PUT("/update", controller.UserUpdate)
			//获取用户列表
			adminAuth.GET("/userList", controller.UserList)
			//获取角色列表
			adminAuth.GET("/roles", controller.GetRoles)
			//管理员获取用户信息
			adminAuth.GET("/userInfo/:id", controller.AdminGetUserInfo)
		}
	}
}
