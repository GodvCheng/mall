package service

import (
	"mall/model"
	"mall/model/dto"
	"mall/server/service/impl"
)

type UserService interface {
	ManagerRegister(user *model.User) error
	UserLogin(username, password string) (string, error)
	GetUserInfo(token string) (dto.UserDto, error)
	UpdateUser(user *model.User) error
	DisableUser(id int) error
	EnableUser(id int) error
	UserList() ([]*dto.UserDto, error)
	GetRoles(token string) ([]*model.Role, error)
	AdminGetUserInfo(id int) (*dto.UserDto, error)
	GetProfile(id int) (*model.User, error)
}

// NewUserService 创建一个UService接口
func NewUserService() UserService {
	return &impl.UserServiceImpl{}
}
