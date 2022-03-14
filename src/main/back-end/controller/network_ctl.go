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

func AddNetworkAction(c *gin.Context) {
	resp := utils.NewBasicResp()
	defer c.JSON(http.StatusOK, resp)
	session := utils.GetSession(c)
	var req model.Network
	// 字段校验
	if err := c.ShouldBind(&req); err != nil {
		resp.Code = model.CodeErr
		resp.Msg = err.Error()
		return
	}
	req.UID = session.UID
	// url和chain id合法性校验
	if err := service.Wallet.IsValidUrlWithChainId(req.Url, req.ChainId); err != nil {
		resp.Code = model.CodeErr
		resp.Msg = fmt.Sprintf("url不合法, 原因: [%s]", err.Error())
		return
	}
	// 存
	if err := service.Wallet.AddOrUpdateNetWork(req); err != nil {
		resp.Code = model.CodeErr
		resp.Msg = err.Error()
		return
	}
}

func UpdateNetworkAction(c *gin.Context) {
	resp := utils.NewBasicResp()
	defer c.JSON(http.StatusOK, resp)
	var req model.Network
	session := utils.GetSession(c)
	if err := c.ShouldBind(&req); err != nil {
		resp.Code = model.CodeErr
		resp.Msg = err.Error()
		return
	}
	if req.UID != session.UID {
		resp.Code = model.CodeErr
		resp.Msg = "没有权限删除"
		return
	}
	if err := service.Wallet.UpdateNetWork(req); err != nil {
		resp.Code = model.CodeErr
		resp.Msg = err.Error()
		return
	}
}

func DeleteNetworkAction(c *gin.Context) {
	resp := utils.NewBasicResp()
	defer c.JSON(http.StatusOK, resp)
	session := utils.GetSession(c)
	var req model.Network
	if err := c.ShouldBind(&req); err != nil {
		resp.Code = model.CodeErr
		resp.Msg = err.Error()
		return
	}
	if req.UID != session.UID {
		resp.Code = model.CodeErr
		resp.Msg = "没有权限删除"
		return
	}
	if err := service.Wallet.DeleteNetWork(req); err != nil {
		resp.Code = model.CodeErr
		resp.Msg = err.Error()
	}
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

func GetCurrentNetworkAction(c *gin.Context) {
	resp := utils.NewBasicResp()
	defer c.JSON(http.StatusOK, resp)
	session := utils.GetSession(c)
	network, err := service.Wallet.GetNetWorkByID(session.NetworkID)
	if err != nil {
		resp.Code = model.CodeErr
		resp.Msg = err.Error()
		return
	}
	resp.Data = network
}

func SetCurrentNetworkAction(c *gin.Context) {
	resp := utils.NewBasicResp()
	defer c.JSON(http.StatusOK, resp)
	session := utils.GetSession(c)
	var req model.Network
	if err := c.ShouldBind(&req); err != nil {
		resp.Code = model.CodeErr
		resp.Msg = err.Error()
		return
	}
	session.NetworkID = req.ID
	if err := session.Save(c); nil != err {
		c.Status(http.StatusInternalServerError)
		return
	}
}