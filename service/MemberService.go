package service

import (
	"mall/model/dto"
	"mall/service/impl"
)

type MemberService interface {
	MemberList(current, pageSize int) ([]*dto.MemberDto, int, error)
}

func NewMemberService() MemberService {
	return &impl.MemberServiceImpl{}
}
