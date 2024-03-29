package controller

import (
	"back-end/model"
	"back-end/service"
	"back-end/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

/**
    controller
    @author: roccoshi
    @desc: 用户注册
**/

func RegisterAction(c *gin.Context) {
	resp := utils.NewBasicResp()
	defer c.JSON(http.StatusOK, resp)
	var iuser model.User
	if c.ShouldBind(&iuser) != nil {
		resp.Code = model.CodeErr
		resp.Msg = "注册出错"
		return
	}
	name := iuser.Name
	password := iuser.PassWord
	if name == "" || password == "" {
		resp.Code = model.CodeErr
		resp.Msg = "用户名和密码不得为空"
		return
	}
	if err := service.User.AddUserByNameAndPassWord(name, password); err != nil {
		resp.Code = model.CodeErr
		resp.Msg = fmt.Sprintf("error: %v", err)
		return
	}
}
