package router

/**
    router
    @author: roccoshi
    @desc: 路由
**/

import (
	ctl "back-end/controller"
	"back-end/utils"
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func InitRoutes() {
	router := gin.New()
	// 中间件
	// 允许跨域: https://github.com/gin-contrib/cors
	router.Use(gin.Logger(), gin.Recovery(), cors.Default())
	// session
	store := cookie.NewStore([]byte(utils.Rand.String(16))) // use 16 random string as secret of session
	router.Use(sessions.Sessions("eth-wallet-session", store))
	// routes
	v1 := router.Group("api/v1")
	// 检查用户是否登录
	{
		v1.GET("/hello-world/:user", ctl.HelloWorldAction)
		v1.POST("/register", ctl.RegisterAction)
	}
	// after auth group
	authGroup := v1.Group("auth")
	authGroup.Use(ctl.LoginCheck)
	{

	}
	router.Run(":8080")
}
