package service

import (
	"mall/api/service/impl"
)

type MemberService interface {
}

func NewMemberService() MemberService {
	return &impl.MemberServiceImpl{}
}
