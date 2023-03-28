package router

import (
	"github.com/gin-gonic/gin"
	"mall/api/controller"
)

func LoadCart(r *gin.Engine) {
	group := r.Group("/api/cart")
	{
		group.POST("/add", controller.AddToCart)
		group.DELETE("/delete/:key/:filed", controller.DelFromCart)
		group.GET("/get/:key", controller.ShowCartInfo)

	}
}
