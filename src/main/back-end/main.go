package main

import (
	"back-end/router"
	"back-end/service"
)

/**
    back_end
    @author: roccoshi
    @desc: 入口
**/

func main() {
	service.ConnectDB() // 连接数据库
	router.InitRoutes() // 初始化路由
}
