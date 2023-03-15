package impl

import (
	"errors"
	"mall/dao"
	"mall/model/dto"
)

var MemberDao = dao.NewMemberDao()

type MemberServiceImpl struct {
}

func (m MemberServiceImpl) MemberList(current, pageSize int) (memberList []*dto.MemberDto, total int, err error) {
	memberList, total = MemberDao.MemberList(current, pageSize)
	if len(memberList) == 0 {
		err = errors.New("会员列表为空")
	}
	return
}
