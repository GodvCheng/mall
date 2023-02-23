package service

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
	"mall/dao"
	"mall/model"
)

var ApiUserDao = dao.NewApiUserDao()

func NewApiUService() ApiUService {
	return &ApiUserService{}
}

type ApiUService interface {
	UserRegister(user *model.User)
}
type ApiUserService struct {
}

func (u *ApiUserService) UserRegister(user *model.User) {
	//  密码加密
	password, err := bcrypt.GenerateFromPassword([]byte(user.Password), 12)
	if err != nil {
		errors.New("用户注册密码加密失败")
	}
	user.Password = string(password)
	ApiUserDao.ApiUserRegister(user)
}
