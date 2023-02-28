package dto

type UserDto struct {
	ID        int    `json:"id"form:"id"`
	Username  string `json:"username"form:"username"` //用户名
	Name      string `json:"name" form:"name"`        //真实姓名
	Sex       string `json:"sex"form:"sex"`
	Age       int    `json:"age"form:"age"`
	Phone     string `json:"phone"form:"phone"`         //手机号
	Email     string `json:"email"form:"email" `        //邮箱
	Avatar    string `json:"avatar"form:"avatar"`       //头像
	Status    int    `json:"status"form:"status"`       //状态 0:未激活 1:激活
	Authority int    `json:"authority"form:"authority"` //权限
	Roles     string `json:"roles" form:"roles"`
}
