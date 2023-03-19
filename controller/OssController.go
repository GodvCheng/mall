package controller

import (
	"github.com/gin-gonic/gin"
	"mall/result"
	"mall/service"
)

var OssService = service.NewOssService()

func UploadAvatar(c *gin.Context) {
	fileHeader, _ := c.FormFile("file")
	url, err := OssService.UploadAvatar(fileHeader)
	if err != nil {
		result.Fail(c, Response{
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
	url, err := OssService.UploadImage(fileHeader)
	if err != nil {
		result.Fail(c, Response{
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
