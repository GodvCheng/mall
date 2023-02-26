package api

import (
	"mall/dao/api"
)

var ApiUserDao = api.NewApiUserDao()

func NewApiUService() ApiUService {
	return &ApiUserService{}
}

type ApiUService interface {
}
type ApiUserService struct {
}
