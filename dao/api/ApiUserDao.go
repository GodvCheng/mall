package api

func NewApiUserDao() ApiUDao {
	return &ApiUserDao{}
}

type ApiUDao interface {
}
type ApiUserDao struct {
}
