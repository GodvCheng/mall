package model

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	Username  string       `json:"username"form:"username" gorm:"unique"` //用户名
	Password  string       `json:"password"form:"password"`               //密码
	Name      string       `json:"name" form:"name"`                      //真实姓名
	Sex       string       `json:"sex"form:"sex"`
	Age       int          `json:"age"form:"age"`
	Phone     string       `json:"phone"form:"phone" gorm:"unique" ` //手机号
	Email     string       `json:"email"form:"email" gorm:"unique" ` //邮箱
	Avatar    string       `json:"avatar"form:"avatar"`              //头像
	Status    int          `json:"status"form:"status"`              //状态 0:未激活 1:激活
	Authority int          `json:"authority"form:"authority"`        //权限
	Address   []*Address   `gorm:"-"`
	OrderInfo []*OrderInfo `gorm:"-"`
}
