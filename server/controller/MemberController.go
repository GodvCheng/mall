package controller

import (
	"github.com/gin-gonic/gin"
	"mall/result"
	"mall/server/service"
	"strconv"
)

var memberService = service.NewMemberService()

func MemberList(c *gin.Context) {
	current, _ := strconv.Atoi(c.Param("current"))
	pageSize, _ := strconv.Atoi(c.Param("pageSize"))
	memberList, total, err := memberService.MemberList(current, pageSize)
	if err != nil {
		result.Fail(c, response{
			Code:    509,
			Message: err.Error(),
			Data:    nil,
		})
	} else {
		result.OkWithData(c, gin.H{
			"memberList": memberList,
			"total":      total,
		})
	}
}
