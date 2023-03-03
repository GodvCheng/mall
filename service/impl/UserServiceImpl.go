package impl

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
	"mall/dao"
	"mall/model"
	"mall/model/dto"
	"mall/util"
	"reflect"
)

var UserDao = dao.NewUserDao()

// UserService 管理用户服务
type UserService struct {
}

func (u *UserService) AdminGetUserInfo(id int) (user *model.User, err error) {
	user = UserDao.AdminGetUserInfo(id)
	if reflect.DeepEqual(user, model.User{}) {
		return nil, errors.New("管理员查询用户信息失败")
	}
	return
}

func (u *UserService) GetRoles(token string) (roles []*model.Role, err error) {
	claims, err1 := util.ParseToken(token)
	if err1 != nil {
		err = err1
		return nil, err
	} else if claims.Authority == 0 {
		return nil, errors.New("权限不足")
	}
	roles = UserDao.GetRoles()
	if len(roles) == 0 {
		return nil, errors.New("角色列表为空")
	}
	return roles, nil
}

func (u *UserService) EnableUser(id int) error {
	n := UserDao.EnableUser(id)
	if n == 0 {
		return errors.New("用户启用失败")
	}
	return nil
}

func (u *UserService) UserList() (userList []*dto.UserDto, err error) {
	userList = UserDao.UserList()
	if len(userList) == 0 {
		return nil, errors.New("查询用户列表失败")
	}
	return userList, nil
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
func (u *UserService) GetUserInfo(token string) (userDto dto.UserDto, err error) {
	claims, err1 := util.ParseToken(token)
	if err1 != nil {
		err = err1
		return dto.UserDto{}, err
	}
	username := claims.Username
	user := UserDao.GetUserInfo(username)
	//将结构体A的值赋值给结构体B 需要传入结构体指针
	util.Copy(&userDto, &user)
	userDto.Roles = make([]string, 1)
	userDto.Roles[0] = user.Role
	return userDto, nil
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
