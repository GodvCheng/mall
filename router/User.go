package router

import (
	"github.com/gin-gonic/gin"
	"mall/common"
	"mall/controller"
	"mall/middleware"
	"net/http"
)

func LoadUser(r *gin.Engine) {
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, common.NewDataResult(gin.H{
			"message": "ok",
		}))
	})

	//用户操作
	group := r.Group("/user")
	{
		//登录 todo 判断是用户还是管理员
		group.POST("/login", controller.UserLogin)
		//后台管理员
		mAuth := group.Group("/")
		mAuth.Use(middleware.JWTAdmin())
		{
			//后台人员注册
			mAuth.POST("/register", controller.UserRegister)
			//获取用户信息
			mAuth.GET("/info", controller.GetUserInfo)
			//更新
			mAuth.PUT("/update/:id", controller.UserUpdate)
			//头像上传
			mAuth.POST("/upload", controller.UploadAvatar)
		}
		//前台访客
		uAuth := group.Group("/api")
		uAuth.POST("/register", controller.ApiUserRegister)
		uAuth.Use(middleware.JWT())
		{
			uAuth.POST("/api/update/:id")

		}
	}
}
