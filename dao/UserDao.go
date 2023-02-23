package dao

import (
	_ "github.com/go-sql-driver/mysql"
	"mall/model"
)

type UDao interface {
	ManagerRegister(user *model.User)
	UpdateUser(user *model.User)
	ExistUser(username string) *model.User
	ExistPhone(phone string)
	ExistEmail(email string)
	GetUserInfo(username string) *model.User
}

type UserDao struct {
}

func NewUserDao() UDao {
	return &UserDao{}
}
func (u *UserDao) GetUserInfo(username string) *model.User {
	var user model.User
	Db.Where("username = ?", username).First(&user)
	return &user
}

func (u *UserDao) ExistUser(username string) *model.User {
	var user model.User
	Db.Where("username = ?", username).Find(&user)
	return &user
}

func (u *UserDao) ExistPhone(phone string) {

}

func (u *UserDao) ExistEmail(email string) {

}

func (u *UserDao) UpdateUser(user *model.User) {
	//更新
	Db.Model(user).Updates(*user)
}

// ManagerRegister 管理员注册
func (u *UserDao) ManagerRegister(user *model.User) {
	Db.Create(user)
	Db.Model(user).Updates(map[string]interface{}{"authority": 1, "status": 1})
}
