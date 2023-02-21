package model

type User struct {
	Id       int    `db:"id" form:"id"`
	Username string `db:"username" form:"username"`
	Password string `db:"password" form:"password"`
}
