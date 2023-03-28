package util

import (
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	uuid "github.com/satori/go.uuid"
	"mall/server/conf"
	"mime/multipart"
	"time"
)

func UploadUserAvatar(fileHeader *multipart.FileHeader) (url string, err error) {
	bucketName := conf.AliConf.BucketName
	client, err1 := oss.New(conf.AliConf.Endpoint, conf.AliConf.AccessKeyId, conf.AliConf.AccessKeySecret)
	if err1 != nil {
		err = err1
	}
	//判断bucket是否存在
	ok, _ := client.IsBucketExist(bucketName)
	if !ok {
		//创建bucket
		err2 := client.CreateBucket(bucketName)
		if err2 != nil {
			err = err2
		}
	}
	bucket, err3 := client.Bucket(bucketName)
	if err3 != nil {
		err = err3
	}

	file, _ := fileHeader.Open()
	defer file.Close()
	//获取文件名
	objectName := fileHeader.Filename
	//获取上传日期并格式化为字符串
	s := time.Now().Format("2006/01/02/")
	//优化存储日期 2022/12/20+UUID+fileName
	objectName = "avatar/" + s + uuid.NewV4().String() + objectName
	err4 := bucket.PutObject(objectName, file)
	if err4 != nil {
		err = err4
	}
	url = "https://" + bucketName + "." + conf.AliConf.Endpoint + "/" + objectName
	return
}

func UploadGoodsImage(fileHeader *multipart.FileHeader) (url string, err error) {
	bucketName := conf.AliConf.BucketName
	client, err1 := oss.New(conf.AliConf.Endpoint, conf.AliConf.AccessKeyId, conf.AliConf.AccessKeySecret)
	if err1 != nil {
		err = err1
	}
	//判断bucket是否存在
	ok, _ := client.IsBucketExist(bucketName)
	if !ok {
		//创建bucket
		err2 := client.CreateBucket(bucketName)
		if err2 != nil {
			err = err2
		}
	}
	bucket, err3 := client.Bucket(bucketName)
	if err3 != nil {
		err = err3
	}
	file, _ := fileHeader.Open()
	defer file.Close()
	//获取文件名
	objectName := fileHeader.Filename
	//获取上传日期并格式化为字符串
	s := time.Now().Format("2006/01/02/")
	//优化存储日期 2022/12/20+UUID+fileName
	objectName = "goods/" + s + uuid.NewV4().String() + objectName
	err4 := bucket.PutObject(objectName, file)
	if err4 != nil {
		err = err4
	}
	url = "https://" + bucketName + "." + conf.AliConf.Endpoint + "/" + objectName
	return
}
