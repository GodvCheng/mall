package dao

import (
	"mall/model"
)

func NewApiUserDao() ApiUDao {
	return &ApiUserDao{}
}

type ApiUDao interface {
	ApiUserRegister(user *model.User)
}
type ApiUserDao struct {
}

func (a ApiUserDao) ApiUserRegister(user *model.User) {
	Db.Create(user)
	Db.Model(user).Update("authority", 0)
}
