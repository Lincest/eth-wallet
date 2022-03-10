package service

import (
	"fmt"
	"os"
	"testing"
)

/**
    service
    @author: roccoshi
    @desc: 为测试类提供上下文环境
**/

func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	teardown()
	os.Exit(code)
}

func setup() {
	fmt.Println("开始初始化测试环境..")
	ConnectDB()
}

func teardown() {
	fmt.Println("结束测试环境")
}
