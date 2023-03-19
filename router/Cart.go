package router

import "github.com/gin-gonic/gin"

func LoadCart(r *gin.Engine) {
	r.Group("/cart")
}
