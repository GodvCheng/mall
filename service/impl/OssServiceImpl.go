package impl

import (
	"errors"
	"fmt"
	"mall/util"
	"mime/multipart"
)

type OssServiceImpl struct {
}

func (o OssServiceImpl) UploadImage(fileHeader *multipart.FileHeader) (url string, err error) {
	url, err = util.UploadGoodsImage(fileHeader)
	if err != nil {
		fmt.Errorf("oss faild:%v", err)
		return "", err
	}
	return url, err
}

func (o OssServiceImpl) UploadAvatar(fileHeader *multipart.FileHeader) (url string, err error) {
	url, err = util.UserUploadAvatar(fileHeader)
	if err != nil {
		fmt.Errorf("oss faild:%v", err)
		return "", errors.New("用户上传头像失败")
	}
	return url, err
}
