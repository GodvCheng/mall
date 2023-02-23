package dao

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"mall/model"
)

type UDao interface {
	ManagerRegister(user *model.User)
	UserLogin(username, password string) *model.User
	UpdateUser(user *model.User)
	ExistUser(username string)
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

func (u *UserDao) ExistUser(username string) {

}

func (u *UserDao) ExistPhone(phone string) {

}

func (u *UserDao) ExistEmail(email string) {

}

func (u *UserDao) UpdateUser(user *model.User) {
	fmt.Println(*user)
	//更新
	Db.Model(user).Updates(*user)
}

func (u *UserDao) UserLogin(username, password string) *model.User {
	// todo 大坑，只能声明为结构体类型，在传值的时候转换为指针
	var user model.User
	Db.Where("username = ? and password = ?", username, password).First(&user)
	return &user
}

// ManagerRegister 管理员注册
func (u *UserDao) ManagerRegister(user *model.User) {
	Db.Create(user)
	Db.Model(user).Update("authority", 1)
}
