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

