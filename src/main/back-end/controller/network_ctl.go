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
    @desc: 网络
**/

func GetNetworkAction(c *gin.Context) {
	resp := utils.NewBasicResp()
	defer c.JSON(http.StatusOK, resp)
	session := utils.GetSession(c)
	res, err := service.Wallet.GetAllNetWorkByUid(session.UID)
	if err != nil {
		resp.Code = model.CodeErr
		resp.Msg = err.Error()
		return
	}
	resp.Data = res
}


func RpcTestAction(c *gin.Context) {
	resp := utils.NewBasicResp()
	defer c.JSON(http.StatusOK, resp)
	url := c.Query("url")
	chainId := c.Query("chain-id")
	if err := service.Wallet.IsValidUrlWithChainId(url, chainId); err != nil {
		resp.Code = model.CodeErr
		resp.Msg = err.Error()
	}
}
