package controller

import (
	"back-end/model"
	"back-end/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

/**
    controller
    @author: roccoshi
    @desc: 检查用户登录状态
**/

// LoginCheck 检查用户登录状态
func LoginCheck(c *gin.Context) {
	session := utils.GetSession(c)
	if 0 == session.UID {
		result := utils.NewBasicResp()
		result.Code = model.CodeErr
		result.Msg = "unauthenticated request"
		c.AbortWithStatusJSON(http.StatusOK, result)
		return
	}
	c.Next()
}
