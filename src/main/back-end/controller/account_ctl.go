package controller

import (
	"back-end/model"
	"back-end/service"
	"back-end/utils"
	"bytes"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
)

/**
    controller
    @author: roccoshi
    @desc: 账户管理
**/

func AddAccountAction(c *gin.Context) {
	resp := utils.NewBasicResp()
	defer c.JSON(http.StatusOK, resp)
	session := utils.GetSession(c)
	var req struct {
		PrivateKeyHex string `json:"private_key_hex" form:"private_key_hex"`
	}
	body, err := ioutil.ReadAll(c.Request.Body)
	c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(body)) // 复原Body, 否则缓冲区已经读完无法使用Bind
	if err != nil {
		resp.Code = model.CodeErr
		resp.Msg = err.Error()
	}
	if len(body) == 0 {
		if err := service.Wallet.AddNewAccount(session.UID); err != nil {
			resp.Code = model.CodeErr
			resp.Msg = err.Error()
		}
	} else {
		if err := c.Bind(&req); err != nil {
			resp.Code = model.CodeErr
			resp.Msg = err.Error()
			return
		}
		if err := service.Wallet.AddNewAccountByUIDAndPrivateKey(session.UID, req.PrivateKeyHex); err != nil {
			resp.Code = model.CodeErr
			resp.Msg = err.Error()
		}
	}
}

func GetAccountAction(c *gin.Context) {
	resp := utils.NewBasicResp()
	defer c.JSON(http.StatusOK, resp)
	session := utils.GetSession(c)
	res, err := service.Wallet.GetAllAccountsByUID(session.UID)
	if err != nil {
		resp.Code = model.CodeErr
		resp.Msg = err.Error()
		return
	}
	resp.Data = res
}
