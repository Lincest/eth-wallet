package controller

import (
	"back-end/model"
	"back-end/utils"
	"github.com/gin-gonic/gin"
	csrf "github.com/utrack/gin-csrf"
	"net/http"
)

/**
    controller
    @author: roccoshi
    @desc: hello world (you know)
**/

func HelloWorldAction(c *gin.Context) {
	resp := utils.NewBasicResp()
	defer c.JSON(http.StatusOK, resp)
	session := utils.GetSession(c)
	if session.UID == 0 {
		resp.Code = model.CodeErr
		resp.Msg = "you are not login"
	}
	resp.Data = gin.H{"name": session.UName, "token": csrf.GetToken(c)}
}
