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
    @desc: 助记词更新
**/

func MnemonicAction(c *gin.Context) {
	resp := utils.NewBasicResp()
	defer c.JSON(http.StatusOK, resp)
	session := utils.GetSession(c)
	iuser := &model.User{Name: session.UName}
	if err := c.ShouldBind(iuser); err != nil {
		resp.Code = model.CodeErr
		resp.Msg = fmt.Sprintf("参数错误, error: [%s]", err)
	}
	if err := service.Mnemonic.UpdateMnemonicByName(iuser.Mnemonic, iuser.Name); err != nil {
		resp.Code = model.CodeErr
		resp.Msg = fmt.Sprintf("保存助记词错误, error: [%s]", err)
	}
}

func GetMnemonicAction(c *gin.Context) {
	resp := utils.NewBasicResp()
	defer c.JSON(http.StatusOK, resp)
	session := utils.GetSession(c)
	uid := session.UID
	mnemonic, err := service.Mnemonic.GetMnemonicByUid(uid)
	if err != nil {
		resp.Code = model.CodeErr
		resp.Msg = fmt.Sprintf("助记词查询错误, error: [%s]", err)
	}
	resp.Data = mnemonic
}

