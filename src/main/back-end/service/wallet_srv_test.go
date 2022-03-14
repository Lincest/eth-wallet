package service

import (
	"testing"
)

/**
    service
    @author: roccoshi
    @desc: test
**/

func TestIsRpcUrlAlive(t *testing.T) {
	err := Wallet.IsValidUrlWithChainId("HTTP://127.0.0.1:7545", "1338")
	if err != nil {
		t.Fail()
		t.Log(err)
	}
}

func TestWalletService_GetAllNetWorkByUid(t *testing.T) {
	res, err := Wallet.GetAllNetWorkByUid(15)
	if err != nil {
		t.Fail()
	}
	t.Logf("%#v", res)
}


func TestWalletService_AddNetWork(t *testing.T) {
	//n1 := &model.Network{Name: "以太坊主网", Url: "https://mainnet.infura.io/v3/9aa3d95b3bc440fa88ea12eaa4456161", ChainId: "1", UID: 0}
	//n2 := &model.Network{Name: "Mike测试网络", Url: "https://mainnet.infura.io/v3/9aa3d95b3bc440fa88ea12eaa4456161", ChainId: "1", UID: 15}
	//_ = Wallet.AddNetWork(*n1)
	//_ = Wallet.AddNetWork(*n2)
}