package service

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
	"mall/dao"
	"mall/model"
	"mall/model/dto"
	"mall/util"
)

var UserDao = dao.NewUserDao()

// UserService 管理用户服务
type UserService struct {
}

func (u *UserService) UserList() (userList []*dto.UserDto, err error) {
	userList = UserDao.UserList()
	if len(userList) == 0 {
		return nil, errors.New("查询用户列表失败")
	}
	return userList, nil
}

type UService interface {
	ManagerRegister(user *model.User) error
	UserLogin(username, password string) (string, error)
	UpdateUser(user *model.User) error
	GetUserInfo(token string) (*model.User, error)
	DisableUser(id int) error
	UserList() ([]*dto.UserDto, error)
}

// NewUService 创建一个 UService接口
func NewUService() UService {
	return &UserService{}
}

func (u *UserService) DisableUser(id int) (err error) {
	n := UserDao.DisableUser(id)
	if n == 0 {
		err = errors.New("用户禁用失败")
	}
	return
}
func (u *UserService) ManagerRegister(user *model.User) (err error) {
	n := UserDao.ExistUser(user.Username)
	if n != 0 {
		return errors.New("该用户已存在，不需要重复注册")
	} else if n := UserDao.ExistEmail(user.Email); n != 0 {
		return errors.New("该邮箱已注册")
	} else if n := UserDao.ExistPhone(user.Phone); n != 0 {
		return errors.New("该手机号已注册")
	}
	//密码加密
	b, err1 := bcrypt.GenerateFromPassword([]byte(user.Password), 12)
	if err1 != nil {
		err = err1
		return
	}
	user.Password = string(b)
	num := UserDao.ManagerRegister(user)
	if num == 0 {
		err = errors.New("注册失败")
	}
	return
}
func (u *UserService) GetUserInfo(token string) (user *model.User, err error) {
	claims, err1 := util.ParseToken(token)
	if err1 != nil {
		err = err1
		return nil, err
	}
	username := claims.Username
	user = UserDao.GetUserInfo(username)
	return user, nil
}

func (u *UserService) UserLogin(username, password string) (token string, err error) {
	//判断用户是否存在
	n := UserDao.ExistUser(username)
	if n == 0 {
		return "", errors.New("该用户不存在，请先注册")
	} else {
		//获取用户信息
		user := UserDao.GetUserInfo(username)
		if user.Status == 0 {
			return "", errors.New("该用户已被禁用，请联系管理员")
		} else {
			err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
			if err != nil {
				return "", errors.New("用户名或密码错误")
			} else {
				//登录成功返回token
				//管理员authority为1，普通用户为0
				token, _ = util.GenerateToken(user.ID, user.Username, user.Authority)
				return token, nil
			}
		}
	}
}

func (u *UserService) UpdateUser(user *model.User) (err error) {
	n := UserDao.UpdateUser(user)
	if n == 0 {
		err = errors.New("更新失败")
	}
	return
}
