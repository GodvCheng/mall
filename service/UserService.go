package service

import (
	"errors"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"mall/dao"
	"mall/model"
	"mall/util"
	"reflect"
)

var UserDao = dao.NewUserDao()

// UserService 管理用户服务
type UserService struct {
	NickName string `form:"nick_name" json:"nick_name"`
	UserName string `form:"user_name" json:"user_name"`
	Password string `form:"password" json:"password"`
	Key      string `form:"key" json:"key"` // 前端进行判断
}

type UService interface {
	ManagerRegister(user *model.User)
	UserLogin(username, password string) string
	UpdateUser(user *model.User)
	GetUserInfo(token string) *model.User
}

// NewUService 创建一个 UService接口
func NewUService() UService {
	return &UserService{}
}
func (u *UserService) ManagerRegister(user *model.User) {
	//密码加密
	b, err := bcrypt.GenerateFromPassword([]byte(user.Password), 12)
	if err != nil {
		errors.New("密码加密错误！")
	}
	fmt.Println(string(b))
	user.Password = string(b)
	UserDao.ManagerRegister(user)
}
func (u *UserService) GetUserInfo(token string) *model.User {
	claims, err := util.ParseToken(token)
	if err != nil {
		errors.New("token parse failed")
	}
	username := claims.Username
	user := UserDao.GetUserInfo(username)
	return user
}

func (u *UserService) UserLogin(username, password string) string {
	//todo 密码解密 登录成功返回token
	user := UserDao.UserLogin(username, password)
	if !reflect.DeepEqual(*user, model.User{}) { //判断user是否存在

		//管理员authority为1，普通用户为0
		token, _ := util.GenerateToken(user.ID, user.Username, user.Authority)
		return token
	}
	return "token is nil!"
}

func (u *UserService) UpdateUser(user *model.User) {
	UserDao.UpdateUser(user)
}
