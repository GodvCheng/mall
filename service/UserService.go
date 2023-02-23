package service

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
	"mall/dao"
	"mall/model"
	"mall/util"
	"reflect"
)

var UserDao = dao.NewUserDao()

// UserService 管理用户服务
type UserService struct {
}

type UService interface {
	ManagerRegister(user *model.User) error
	UserLogin(username, password string) (string, error)
	UpdateUser(user *model.User) error
	GetUserInfo(token string) (*model.User, error)
}

// NewUService 创建一个 UService接口
func NewUService() UService {
	return &UserService{}
}
func (u *UserService) ManagerRegister(user *model.User) (err error) {
	//密码加密
	b, err1 := bcrypt.GenerateFromPassword([]byte(user.Password), 12)
	if err1 != nil {
		err = err1
	}
	user.Password = string(b)
	UserDao.ManagerRegister(user)
	return
}
func (u *UserService) GetUserInfo(token string) (user *model.User, err error) {
	claims, err1 := util.ParseToken(token)
	if err1 != nil {
		err = err1
	}
	username := claims.Username
	user = UserDao.GetUserInfo(username)
	return
}

func (u *UserService) UserLogin(username, password string) (token string, err error) {
	//判断用户是否存在
	existUser := UserDao.ExistUser(username)
	if reflect.DeepEqual(*existUser, model.User{}) {
		return "", errors.New("该用户不存在，请先注册")
	} else {
		// 密码解密 登录成功返回token
		err := bcrypt.CompareHashAndPassword([]byte(existUser.Password), []byte(password))
		if err == nil {
			//管理员authority为1，普通用户为0
			token, _ := util.GenerateToken(existUser.ID, existUser.Username, existUser.Authority)
			return token, nil
		}
	}
	return "", errors.New("token is nil")
}

func (u *UserService) UpdateUser(user *model.User) error {
	UserDao.UpdateUser(user)
	return nil
}
