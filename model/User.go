package model

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	Username  string       `db:"username" form:"username"` //用户名
	Password  string       `db:"password" form:"password"` //密码
	Phone     string       `gorm:"unique"`                 //手机号
	Email     string       `gorm:"unique"`                 //邮箱
	Avatar    string       //头像
	Status    int          //状态 0:未激活 1:激活
	Authority int          //权限
	Address   []*Address   `gorm:"-"`
	OrderInfo []*OrderInfo `gorm:"-"`
}
