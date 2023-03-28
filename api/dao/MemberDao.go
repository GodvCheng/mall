package dao

type MemberDao interface {
}

func NewMemberDao() MemberDao {
	return &MemberDaoImpl{}
}

type MemberDaoImpl struct {
}
