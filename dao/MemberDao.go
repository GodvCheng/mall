package dao

import "mall/model/dto"

type MemberDao interface {
	MemberList(current, pageSize int) ([]*dto.MemberDto, int)
}

func NewMemberDao() MemberDao {
	return &MemberDaoImpl{}
}

type MemberDaoImpl struct {
}

func (m MemberDaoImpl) MemberList(current, pageSize int) (memberList []*dto.MemberDto, total int) {
	Db.Table("member").Scopes(Paginate(current, pageSize)).Find(&memberList)
	Db.Table("member").Count(&total)
	return
}
