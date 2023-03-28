package impl

import (
	"errors"
	"mall/model/dto"
	"mall/server/dao"
)

var memberDao = dao.NewMemberDao()

type MemberServiceImpl struct {
}

func (m MemberServiceImpl) MemberList(current, pageSize int) (memberList []*dto.MemberDto, total int, err error) {
	memberList, total = memberDao.MemberList(current, pageSize)
	if len(memberList) == 0 {
		err = errors.New("会员列表为空")
	}
	return
}
