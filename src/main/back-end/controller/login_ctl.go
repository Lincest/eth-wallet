package controller

import (
	"back-end/model"
	"back-end/service"
	"back-end/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

/**
    controller
    @author: roccoshi
    @desc: 处理用户登录
**/

func LoginAction(c *gin.Context) {
	resp := utils.NewBasicResp()
	defer c.JSON(http.StatusOK, resp)
	var iuser model.User
	if c.ShouldBind(&iuser) != nil {
		resp.Code = model.CodeErr
		resp.Msg = "登录出错"
		return
	}
	user := service.User.GetUserByNameAndPassWord(iuser.Name, iuser.PassWord)
	if user == nil {
		resp.Code = model.CodeErr
		resp.Msg = "用户名或密码错误"
		return
	}
	session := &utils.SessionData{UID: user.ID, UName: user.Name, UPassword: iuser.PassWord, NetworkID: 1}
	if err := session.Save(c); nil != err {
		c.Status(http.StatusInternalServerError)
		return
	}
	resp.Msg = "登陆成功"
}
