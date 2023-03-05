package result

var MsgFlags = map[int]string{
	SUCCESS:                 "ok",
	UpdatePasswordSuccess:   "修改密码成功",
	ERROR:                   "fail",
	InvalidParams:           "请求参数错误",
	ErrorExistUser:          "已存在该用户名",
	ErrorNotExistUser:       "该用户不存在",
	ErrorNotCompare:         "账号密码错误",
	ErrorNotComparePassword: "两次密码输入不一致",
	ErrorFailEncryption:     "加密失败",
	ErrorNotExistProduct:    "该商品不存在",
	ErrorNotExistAddress:    "该收获地址不存在",
	ErrorUserNotFound:       "用户不存在",
	ErrorUserNotLogin:       "用户未登录",
	ErrorProductExistCart:   "商品已经在购物车了，数量+1",
	ErrorProductMoreCart:    "超过最大上限",

	ErrorAuthCheckTokenFail:        "Token鉴权失败",
	ErrorAuthCheckTokenTimeout:     "Token已超时",
	ErrorAuthToken:                 "Token生成失败",
	ErrorAuthIsNull:                "token is null",
	ErrorAuthInsufficientAuthority: "权限不足",
	ErrorSendEmail:                 "发送邮件失败",
	ErrorCallApi:                   "调用接口失败",
	ErrorUnmarshalJson:             "解码JSON失败",

	ErrorUploadFile:    "上传失败",
	ErrorAdminFindUser: "管理员查询用户失败",

	ErrorDatabase: "数据库操作出错,请重试",

	ErrorOss: "OSS配置错误",
}

// GetMsg 获取状态码对应信息
func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}
	return MsgFlags[ERROR]
}
