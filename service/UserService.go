package service

import (
	"github.com/gin-gonic/gin"
	"mall/dao"
	"mall/model"
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
	UserRegister(c *gin.Context)
	UserLogin(username, password string) model.User
	UpdateUser(c *gin.Context, uId uint)
}

// NewUService 创建一个 UService接口
func NewUService() UService {
	return &UserService{}
}

func (u *UserService) UserRegister(c *gin.Context) {

}

func (u *UserService) UserLogin(username, password string) (user model.User) {
	user = UserDao.UserLogin(username, password)
	return
}

func (u *UserService) UpdateUser(c *gin.Context, uId uint) {

}
