package controller

import (
	"back-end/model"
	"back-end/service"
	"back-end/utils"
	"github.com/ethereum/go-ethereum/common"
	"github.com/gin-gonic/gin"
	"net/http"
)

/**
    controller
    @author: roccoshi
    @desc: 交易
**/

func NewTransactionAction(c *gin.Context) {
	resp := utils.NewBasicResp()
	defer c.JSON(http.StatusOK, resp)
	session := utils.GetSession(c)
	var req model.TransactionReq
	if err := c.ShouldBind(&req); err != nil {
		resp.Code = model.CodeErr
		resp.Msg = err.Error()
		return
	}
	network, err := service.Wallet.GetNetWorkByID(session.NetworkID)
	if err != nil {
		resp.Code = model.CodeErr
		resp.Msg = err.Error()
		return
	}
	transaction, err := service.Transaction.CreateTransaction(
		session.UID,
		common.HexToAddress(req.FromAddress),
		req.FromPrivateKeyHex,
		common.HexToAddress(req.ToAddress),
		req.Value,
		req.GasLimit,
		req.GasPrice,
		network.Url)
	if err != nil {
		resp.Code = model.CodeErr
		resp.Msg = err.Error()
		return
	}
	resp.Data = transaction
}

func CheckTransactionAction(c *gin.Context) {
	resp := utils.NewBasicResp()
	defer c.JSON(http.StatusOK, resp)
	session := utils.GetSession(c)
	transactionHash := c.Param("transaction-hash")
	network, err := service.Wallet.GetNetWorkByID(session.NetworkID)
	if err != nil {
		resp.Code = model.CodeErr
		resp.Msg = err.Error()
		return
	}
	tx, err := service.Transaction.GetAndUpdateTransactionByHash(transactionHash, network.Url)
	if err != nil {
		resp.Code = model.CodeErr
		resp.Msg = err.Error()
		return
	}
	resp.Data = tx
}
