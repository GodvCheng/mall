package dao

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"mall/model"
)

type UDao interface {
	UserRegister()
	UserLogin(username, password string) model.User
	UpdateUser()
}

type UserDao struct {
}

func NewUserDao() UDao {
	return &UserDao{}
}

func (u *UserDao) UpdateUser() {

}

func (u *UserDao) UserLogin(username, password string) (user model.User) {
	db, err := sqlx.Open("mysql", "root:141097@(localhost:3306)/jmall")
	if err != nil {
		return
	}
	rows, err := db.Query("select * from user where username = ? and password = ?", username, password)
	if err != nil {
		return
	}
	rows.Scan(user)
	return
}
func (u *UserDao) UserRegister() {

}
