package result

const (
	SUCCESS               = 200
	UpdatePasswordSuccess = 201
	ERROR                 = 500
	InvalidParams         = 400

	ErrorExistUser          = 10002
	ErrorNotExistUser       = 10003
	ErrorNotCompare         = 10004
	ErrorNotComparePassword = 10005
	ErrorFailEncryption     = 10006
	ErrorNotExistProduct    = 10007
	ErrorNotExistAddress    = 10008
	ErrorUserNotFound       = 10010
	ErrorUserNotLogin       = 10009
	ErrorProductExistCart   = 20007
	ErrorProductMoreCart    = 20008

	ErrorAuthCheckTokenFail        = 30001 //token 错误
	ErrorAuthCheckTokenTimeout     = 30002 //token 过期
	ErrorAuthToken                 = 30003
	ErrorAuthIsNull                = 508
	ErrorAuthInsufficientAuthority = 30005
	ErrorSendEmail                 = 30007
	ErrorCallApi                   = 30008
	ErrorUnmarshalJson             = 30009
	ErrorAdminFindUser             = 30010

	ErrorDatabase = 40001

	ErrorOss        = 50001
	ErrorUploadFile = 50002
)
