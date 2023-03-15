package controller

import (
	"github.com/gin-gonic/gin"
	"mall/result"
	"mall/service"
	"strconv"
)

var MemberService = service.NewMemberService()

func MemberList(c *gin.Context) {
	current, _ := strconv.Atoi(c.Param("current"))
	pageSize, _ := strconv.Atoi(c.Param("pageSize"))
	memberList, total, err := MemberService.MemberList(current, pageSize)
	if err != nil {
		result.Fail(c, Response{
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
