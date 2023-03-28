package service

import (
	"mall/server/service/impl"
	"mime/multipart"
)

type OssService interface {
	UploadAvatar(fileHeader *multipart.FileHeader) (string, error)
	UploadImage(fileHeader *multipart.FileHeader) (string, error)
}

func NewOssService() OssService {
	return &impl.OssServiceImpl{}
}
