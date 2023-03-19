package result

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Response struct {
	Code    int                    `json:"code"`
	Message string                 `json:"message"`
	Data    map[string]interface{} `json:"data"`
}

func Result(c *gin.Context, code int, msg string, data map[string]interface{}) {
	c.JSON(http.StatusOK, Response{
		code,
		msg,
		data,
	})
}
func Error(code int, msg string) Response {
	return Response{
		code,
		msg,
		nil,
	}
}

func Ok(c *gin.Context) {
	OkWithData(c, nil)
}

func OkWithData(c *gin.Context, data map[string]interface{}) {
	Result(c, SUCCESS, MsgFlags[SUCCESS], data)
}

func OkWithMsg(c *gin.Context, msg string) {
	Result(c, SUCCESS, msg, nil)
}

func Fail(c *gin.Context, err Response) {
	Result(c, err.Code, err.Message, nil)
}

func FailWithMsg(c *gin.Context, err Response, msg string) {
	Result(c, err.Code, msg, nil)
}
