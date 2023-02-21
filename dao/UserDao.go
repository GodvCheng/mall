package dao

import (
	_ "github.com/go-sql-driver/mysql"
	"mall/model"
)

type UDao interface {
	UserRegister()
	UserLogin(username, password string) model.User
	UpdateUser(user *model.User)
}

type UserDao struct {
}

func NewUserDao() UDao {
	return &UserDao{}
}

func (u *UserDao) UpdateUser(user *model.User) {
	Db.Model(user).Update(user)
}

func (u *UserDao) UserLogin(username, password string) (user model.User) {
	Db.Where("username = ? and password = ?", username, password).First(&user)
	return user
}
func (u *UserDao) UserRegister() {

}
