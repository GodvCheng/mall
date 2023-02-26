package util

import "github.com/aliyun/aliyun-oss-go-sdk/oss"

func UploadFile(path string) (err error) {
	endpoint := "https://oss-cn-qingdao.aliyuncs.com"
	// 阿里云账号AccessKey拥有所有API的访问权限，风险很高。强烈建议您创建并使用RAM用户进行API访问或日常运维，请登录RAM控制台创建RAM用户。
	accessKeyId := "LTAI5tHU5a9g4yuJzoSNjwKM"
	accessKeySecret := "wyxSSSOmhoxuqcY2CiD3pITJPZ2yjC"
	// yourBucketName填写Bucket名称。
	bucketName := "gin-mall"
	client, err1 := oss.New(endpoint, accessKeyId, accessKeySecret)
	if err1 != nil {
		err = err1
		return
	}
	bucket, err2 := client.Bucket(bucketName)
	if err2 != nil {
		err = err2
		return
	}
	bucket.PutObjectFromFile("", "")
	return nil
}
