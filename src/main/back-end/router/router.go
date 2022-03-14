package router

/**
    router
    @author: roccoshi
    @desc: 路由
**/

import (
	"back-end/conf"
	ctl "back-end/controller"
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"log"
)

func InitRoutes() {
	router := gin.New()
	// 中间件
	// 允许跨域: https://github.com/gin-contrib/cors
	router.Use(gin.Logger(), gin.Recovery(), cors.Default())
	// session
	// store := cookie.NewStore([]byte(utils.Rand.String(16))) // use 16 random string as secret of session, 这样会导致每次服务器重启之前的用户session失效
	sessionConf := conf.Config.Session                   // session config
	store := cookie.NewStore([]byte(sessionConf.Secret)) // 可以手动指定一串secret防止重启服务器cookie失效, 但是这种secret不应该出现在源代码中, 这里为了简化就直接把密钥写死
	router.Use(sessions.Sessions(sessionConf.Name, store))
	// routes
	v1 := router.Group("api/v1")
	{
		v1.POST("/register", ctl.RegisterAction)
		v1.POST("/login", ctl.LoginAction)
		v1.GET("/test-rpc", ctl.RpcTestAction)
	}
	// after auth group
	authGroup := v1.Group("auth")
	authGroup.Use(ctl.LoginCheck) // auth中间件检查用户是否登录
	{
		authGroup.GET("/hello-world", ctl.HelloWorldAction)
		authGroup.POST("/logout", ctl.LogoutAction)
		authGroup.POST("/mnemonic", ctl.MnemonicAction)
		authGroup.GET("/network", ctl.GetNetworkAction)
		authGroup.DELETE("/network", ctl.DeleteNetworkAction)
		authGroup.POST("/network", ctl.AddNetworkAction)
		authGroup.PUT("/network", ctl.UpdateNetworkAction)
	}
	err := router.Run(":8765")
	if err != nil {
		log.Fatalf("服务器启动失败, error: %v", err)
	}
}
