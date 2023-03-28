package controller

import (
	"github.com/gin-gonic/gin"
	"mall/api/service"
	"mall/model/dto"
	"mall/result"
)

var cartService = service.NewCartService()

func AddToCart(ctx *gin.Context) {
	var cart dto.CartDto
	ctx.ShouldBind(&cart)
	err := cartService.AddToCart(&cart)
	if err != nil {
		result.Fail(ctx, response{
			Code:    509,
			Message: err.Error(),
			Data:    nil,
		})
	} else {
		result.Ok(ctx)
	}
}

func DelFromCart(ctx *gin.Context) {
	key := ctx.Param("key")
	filed := ctx.Param("filed")
	err := cartService.DelFromCart(key, filed)
	if err != nil {
		result.Fail(ctx, response{
			Code:    509,
			Message: err.Error(),
			Data:    nil,
		})
	} else {
		result.Ok(ctx)
	}
}

func ShowCartInfo(ctx *gin.Context) {
	key := ctx.Param("key")
	cartInfoList, err := cartService.ShowCartInfo(key)
	if err != nil {
		result.Fail(ctx, response{
			Code:    509,
			Message: err.Error(),
			Data:    nil,
		})
	} else {
		result.OkWithData(ctx, gin.H{"cartInfoList": cartInfoList})
	}
}
