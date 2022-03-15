package controller

import (
	"back-end/model"
	"back-end/service"
	"back-end/utils"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"os"
)

/**
    controller
    @author: roccoshi
    @desc: keystore 管理
**/

func AddAccountByKeyStoreAction(c *gin.Context) {
	resp := utils.NewBasicResp()
	defer c.JSON(http.StatusOK, resp)
	session := utils.GetSession(c)
	passphrase := c.PostForm("passphrase")
	form, err := c.FormFile("keystore")
	if err != nil {
		resp.Code = model.CodeErr
		resp.Msg = err.Error()
	}
	if err := service.KeyStore.AddOneAccountByKeyStoreFile(form, session.UID, passphrase); err != nil {
		resp.Code = model.CodeErr
		resp.Msg = err.Error()
	}
}

func GetKeyStoreAction(c *gin.Context) {
	resp := utils.NewBasicResp()
	defer c.JSON(http.StatusOK, resp)
	session := utils.GetSession(c)
	zipFilePath, err := service.KeyStore.GetAllKeyStoreFilesByUID(session.UID, session.UPassword)
	if err != nil {
		resp.Code = model.CodeErr
		resp.Msg = err.Error()
		return
	}
	c.Header("Content-Disposition", "attachment; filename="+session.UName+"-keystore.zip")
	c.Header("Content-Type", "application/zip")
	file, err := os.Open(zipFilePath)
	if err != nil {
		resp.Code = model.CodeErr
		resp.Msg = err.Error()
		return
	}
	defer file.Close()
	io.Copy(c.Writer, file)
}
