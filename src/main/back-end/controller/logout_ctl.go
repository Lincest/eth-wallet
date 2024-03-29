package controller

import (
	"back-end/model"
	"back-end/utils"
	"fmt"
	"github.com/ethereum/go-ethereum/log"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
)

/**
    controller
    @author: roccoshi
    @desc: 退出登录
**/

func LogoutAction(c *gin.Context) {
	resp := utils.NewBasicResp()
	defer c.JSON(http.StatusOK, resp)
	session := sessions.Default(c)
	session.Options(sessions.Options{
		Path:   "/",
		MaxAge: -1,
	})
	session.Clear()
	if err := session.Save(); err != nil {
		log.Error("save session failed")
		resp.Code = model.CodeErr
		resp.Msg = fmt.Sprintf("登出失败: %s", err)
	}
}
