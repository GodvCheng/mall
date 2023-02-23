package service

import (
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
	// todo 密码加密
	ApiUserDao.ApiUserRegister(user)
}
