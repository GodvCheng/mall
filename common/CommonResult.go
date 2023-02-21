package common

import "net/http"

// Result 定义统一结果返回
type Result struct {
	Code    int    `json:"code"`
	Data    any    `json:"data"`
	Message string `json:"message"`
}

func NewOkResult(msg string) *Result {
	return &Result{
		Code:    SUCCESS,
		Data:    nil,
		Message: msg,
	}
}

// NewDataResult 返回数据
func NewDataResult(data any) *Result {
	return &Result{
		Code:    SUCCESS,
		Data:    data,
		Message: "",
	}
}

func NewErrorResult(code int, msg string) *Result {
	return &Result{
		Code:    code,
		Data:    nil,
		Message: msg,
	}
}

type Error struct {
	Status int
	Code   string
	Msg    string
}

func (e *Error) Error() string {
	return e.Msg
}

func BadRequestErr(code string, msg string) error {
	return &Error{
		Status: http.StatusBadRequest,
		Code:   code,
		Msg:    msg,
	}
}

func NotFoundErr(code string, msg string) error {
	return &Error{
		Status: http.StatusNotFound,
		Code:   code,
		Msg:    msg,
	}
}

func UnauthorizedErr() error {
	return &Error{
		Status: http.StatusUnauthorized,
		Code:   "Unauthorized",
		Msg:    "",
	}
}

func InternalServerErr(code string, msg string) error {
	return &Error{
		Status: http.StatusInternalServerError,
		Code:   code,
		Msg:    msg,
	}
}
