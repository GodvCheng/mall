package dao

import (
	_ "github.com/go-sql-driver/mysql"
	"mall/model"
	"mall/model/dto"
)

type UDao interface {
	ManagerRegister(user *model.User) int
	UpdateUser(user *model.User) int
	ExistUser(username string) int
	ExistPhone(phone string) int
	ExistEmail(email string) int
	GetUserInfo(username string) *model.User
	DisableUser(userId int) int
	UserList() []*dto.UserDto
}

type UserDao struct {
}

func (u *UserDao) UserList() (userList []*dto.UserDto) {
	Db.Table("user").Find(&userList)
	return
}

func (u *UserDao) DisableUser(userId int) int {
	num := Db.Model(&model.User{}).Where("id = ?", userId).Update("status", 0).RowsAffected
	return int(num)
}

func NewUserDao() UDao {
	return &UserDao{}
}
func (u *UserDao) GetUserInfo(username string) *model.User {
	var user model.User
	Db.Where("username = ?", username).First(&user)
	return &user
}

func (u *UserDao) ExistUser(username string) int {
	var user model.User
	n := Db.Where("username = ?", username).Find(&user).RowsAffected
	return int(n)
}

func (u *UserDao) ExistPhone(phone string) int {
	var user model.User
	num := Db.Where("phone = ?", phone).Find(&user).RowsAffected
	return int(num)
}

func (u *UserDao) ExistEmail(email string) int {
	var user model.User
	num := Db.Where("email = ?", email).Find(&user).RowsAffected
	return int(num)
}

// UpdateUser 更新
func (u *UserDao) UpdateUser(user *model.User) int {
	num := Db.Model(user).Updates(*user).RowsAffected
	return int(num)
}

// ManagerRegister 普通管理员注册 authority为0
func (u *UserDao) ManagerRegister(user *model.User) int {
	Db.Create(user)
	num := Db.Model(user).Update("authority", 0).RowsAffected
	return int(num)
}
