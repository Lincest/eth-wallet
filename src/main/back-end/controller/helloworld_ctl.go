package controller

import (
	"back-end/model"
	"back-end/utils"
	"fmt"
	"github.com/gin-gonic/gin"
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
	var user = c.Param("user")
	if user == "" {
		resp.Code = model.CodeErr
		resp.Msg = "user not exit"
	}
	resp.Data = fmt.Sprintf("hello world %s", user)
}
