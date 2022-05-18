package main

import (
	"back-end/conf"
	"back-end/router"
	"back-end/service"
)

/**
    back_end
    @author: roccoshi
    @desc: 入口
**/

func main() {
	// 加载配置文件config.yml
	conf.LoadConfig()
	// 连接数据库
	service.ConnectDB()
	// 初始化路由, 启动服务
	router.InitRoutes()
}