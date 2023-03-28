package router

import (
	"github.com/gin-gonic/gin"
	"mall/server/controller"
	"mall/server/middleware"
)

func LoadUser(r *gin.Engine) {

	group := r.Group("/user")
	{
		//管理员用户登录
		group.POST("/login", controller.UserLogin)
		//普通管理员权限
		mAuth := group.Group("/")
		mAuth.Use(middleware.JWT())
		{
			mAuth.POST("/info", controller.GetUserInfo)
			//退出登录
			mAuth.POST("/logout", controller.Logout)
			//用户头像上传 TODO
			mAuth.POST("/upload", controller.UploadAvatar)
			//登录后查看个人信息
			mAuth.GET("/profile/:id", controller.GetProfile)
			//登录后更新个人信息
			mAuth.PUT("/update", controller.UserUpdate)
		}
		adminAuth := group.Group("/")
		//超级管理员权限
		adminAuth.Use(middleware.JWTAdmin())
		{
			//普通管理员注册
			adminAuth.POST("/register", controller.UserRegister)
			//禁用用户
			adminAuth.PUT("/disable/:id", controller.DisableUser)
			//启用用户
			adminAuth.PUT("/enable/:id", controller.EnableUser)
			//获取普通管理员列表
			adminAuth.GET("/userList", controller.UserList)
			//获取角色列表
			adminAuth.GET("/roles", controller.GetRoles)
			//管理员获取普通管理员信息
			adminAuth.GET("/userInfo/:id", controller.AdminGetUserInfo)
		}
	}
}
