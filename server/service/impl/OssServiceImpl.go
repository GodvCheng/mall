package impl

import (
	"errors"
	"fmt"
	"mall/server/util"
	"mime/multipart"
)

type OssServiceImpl struct {
}

func (o OssServiceImpl) UploadImage(fileHeader *multipart.FileHeader) (url string, err error) {
	url, err = util.UploadGoodsImage(fileHeader)
	if err != nil {
		fmt.Errorf("oss faild with upload goods image:%v", err)
		return "", err
	}
	return url, err
}

func (o OssServiceImpl) UploadAvatar(fileHeader *multipart.FileHeader) (url string, err error) {
	url, err = util.UploadUserAvatar(fileHeader)
	if err != nil {
		fmt.Errorf("oss faild with upload user avatar:%v", err)
		return "", errors.New("用户上传头像失败")
	}
	return url, err
}
