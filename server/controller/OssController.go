package controller

import (
	"github.com/gin-gonic/gin"
	"mall/result"
	"mall/server/service"
)

var ossService = service.NewOssService()

func UploadAvatar(c *gin.Context) {
	fileHeader, _ := c.FormFile("file")
	url, err := ossService.UploadAvatar(fileHeader)
	if err != nil {
		result.Fail(c, response{
			Code:    509,
			Message: err.Error(),
			Data:    nil,
		})
	} else {
		result.OkWithData(c, gin.H{
			"url": url,
		})
	}
}

func UploadImage(c *gin.Context) {
	fileHeader, _ := c.FormFile("file")
	url, err := ossService.UploadImage(fileHeader)
	if err != nil {
		result.Fail(c, response{
			Code:    509,
			Message: err.Error(),
			Data:    nil,
		})
	} else {
		result.OkWithData(c, gin.H{
			"url": url,
		})
	}
}

func UploadBanner(c *gin.Context) {
	file, _ := c.FormFile("file")
	url, err := ossService.UploadImage(file)
	if err != nil {
		result.Fail(c, result.Response{
			Code:    509,
			Message: err.Error(),
			Data:    nil,
		})
	} else {
		result.OkWithData(c, gin.H{
			"url": url,
		})
	}
}
