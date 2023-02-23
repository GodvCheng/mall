package model

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	Username  string       `form:"username" gorm:"unique"` //用户名
	Password  string       `form:"password"`               //密码
	Sex       string       `form:"sex"`
	Age       int          `form:"age"`
	Phone     string       `form:"phone" gorm:"unique" ` //手机号
	Email     string       `form:"email" gorm:"unique" ` //邮箱
	Avatar    string       `form:"avatar"`               //头像
	Status    int          `form:"status"`               //状态 0:未激活 1:激活
	Authority int          `form:"authority"`            //权限
	Address   []*Address   `gorm:"-"`
	OrderInfo []*OrderInfo `gorm:"-"`
}
