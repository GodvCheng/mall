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

	//管理员用户操作
	group := r.Group("/user")
	{
		//登录
		group.POST("/login", controller.UserLogin)
		//需要权限
		auth := group.Group("/")
		auth.Use(middleware.JWTAdmin())
		{
			//注册
			group.POST("/register", controller.UserRegister)
			//更新
			group.PUT("/update/:id", controller.UserUpdate)
			//头像上传
			group.POST("/upload/", controller.UploadAvatar)
			//发送邮箱
			group.POST("/sendEmail", controller.SendEmail)
			//验证邮箱
			group.POST("/valEmail", controller.ValidEmail)
		}
	}
}
