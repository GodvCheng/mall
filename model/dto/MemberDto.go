package dto

type MemberDto struct {
	ID       uint   `json:"id"form:"id"`
	Username string `json:"username"form:"username"`
	Phone    string `json:"phone"form:"phone"`
	Email    string `json:"email"form:"email"`
	Age      uint   `json:"age"form:"age"`
	Sex      string `json:"sex"form:"sex"`
}
