package service

import (
	"mall/model"
	"mall/model/dto"
	"mall/service/impl"
)

type UService interface {
	ManagerRegister(user *model.User) error
	UserLogin(username, password string) (string, error)
	UpdateUser(user *model.User) error
	GetUserInfo(token string) (dto.UserDto, error)
	DisableUser(id int) error
	EnableUser(id int) error
	UserList() ([]*dto.UserDto, error)
	GetRoles(token string) ([]*model.Role, error)
	AdminGetUserInfo(id int) (*dto.UserDto, error)
	GetProfile(id int) (*model.User, error)
}

// NewUService 创建一个UService接口
func NewUService() UService {
	return &impl.UserService{}
}
