package utils

import (
	"back-end/conf"
	"fmt"
	"os"
	"testing"
)

/**
    utils
    @author: roccoshi
    @desc: init test for utils
**/

func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	teardown()
	os.Exit(code)
}

func setup() {
	fmt.Println("开始初始化测试环境..")
	conf.LoadConfigForTest()
}

func teardown() {
	fmt.Println("结束测试环境")
}
