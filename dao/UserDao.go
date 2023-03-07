package dao

import (
	_ "github.com/go-sql-driver/mysql"
	"mall/model"
	"mall/model/dto"
)

type UserDao interface {
	ManagerRegister(user *model.User) int
	UpdateUser(user *model.User) int
	ExistUser(username string) int
	ExistPhone(phone string) int
	ExistEmail(email string) int
	GetUserInfo(username string) model.User
	DisableUser(userId int) int
	UserList() []*dto.UserDto
	EnableUser(id int) int
	GetRoles() []*model.Role
	AdminGetUserInfo(id int) *dto.UserDto
	GetProfile(id int) *model.User
}

func NewUserDao() UserDao {
	return &UserDaoImpl{}
}

type UserDaoImpl struct {
}

func (u *UserDaoImpl) GetProfile(id int) *model.User {
	var user model.User
	Db.Where("id = ?", id).Find(&user)
	return &user
}

func (u *UserDaoImpl) AdminGetUserInfo(id int) *dto.UserDto {
	var user dto.UserDto
	Db.Table("user").Where("id = ?", id).Find(&user)
	return &user
}

func (u *UserDaoImpl) GetRoles() (roles []*model.Role) {
	Db.Find(&roles)
	return
}

func (u *UserDaoImpl) EnableUser(id int) int {
	num := Db.Model(&model.User{}).Where("id = ?", id).Update("status", 1).RowsAffected
	return int(num)
}

func (u *UserDaoImpl) UserList() (userList []*dto.UserDto) {
	Db.Table("user").Where("authority = ?", 0).Find(&userList)
	return
}

func (u *UserDaoImpl) DisableUser(userId int) int {
	num := Db.Model(&model.User{}).Where("id = ?", userId).Update("status", 0).RowsAffected
	return int(num)
}

func (u *UserDaoImpl) GetUserInfo(username string) model.User {
	var user model.User
	Db.Where("username = ?", username).First(&user)
	return user
}

func (u *UserDaoImpl) ExistUser(username string) int {
	var user model.User
	n := Db.Where("username = ?", username).Find(&user).RowsAffected
	return int(n)
}

func (u *UserDaoImpl) ExistPhone(phone string) int {
	var user model.User
	num := Db.Where("phone = ?", phone).Find(&user).RowsAffected
	return int(num)
}

func (u *UserDaoImpl) ExistEmail(email string) int {
	var user model.User
	num := Db.Where("email = ?", email).Find(&user).RowsAffected
	return int(num)
}

// UpdateUser 更新
func (u *UserDaoImpl) UpdateUser(user *model.User) int {
	num := Db.Save(user).RowsAffected
	return int(num)
}

// ManagerRegister 普通管理员注册 authority为0
func (u *UserDaoImpl) ManagerRegister(user *model.User) int {
	Db.Create(user)
	num := Db.Model(user).Update("authority", 0).RowsAffected
	return int(num)
}
