package service

import (
	"back-end/model"
	"github.com/ethereum/go-ethereum/common"
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
	n1 := &model.Network{Name: "以太坊主网", Url: "https://mainnet.infura.io/v3/9aa3d95b3bc440fa88ea12eaa4456161", ChainId: "1", UID: 0}
	n1.ID = 1
	//n2 := &model.Network{Name: "Mike测试网络", Url: "https://mainnet.infura.io/v3/9aa3d95b3bc440fa88ea12eaa4456161", ChainId: "1", UID: 15}
	_ = Wallet.AddOrUpdateNetWork(*n1)
	//_ = Wallet.AddOrUpdateNetWork(*n2)
}

func TestNetWorkNotExist(t *testing.T) {
	existNetwork := &model.Network{}
	err := db.Limit(1).Find(&existNetwork, 15).Error
	if err != nil {
		t.Logf("err = %s", err)
		t.Fail()
	}
	t.Logf("%v", existNetwork)
	t.Logf("is nil ? %v", existNetwork.ID == 0)
}

func TestWalletService_GetNetWorkByID(t *testing.T) {
	res, _ := Wallet.GetNetWorkByID(16)
	t.Logf("%v", res)
}

func TestWalletService_InitWallet(t *testing.T) {
	err := Wallet.InitWallet(28)
	if err != nil {
		t.Error(err)
	}
}

func TestWalletService_AddNewAccount(t *testing.T) {
	err := Wallet.AddNewAccount(28)
	if err != nil {
		t.Error(err)
	}
}

func TestWalletService_GetBalanceByAddress(t *testing.T) {
	addr := common.HexToAddress("0xc06EbC7F2d398C156C811A16E89dB3D0ABa61D9e")
	eth, err := Wallet.GetBalanceByAddress(addr, 2)
	if err != nil {
		t.Fail()
	}
	t.Logf("eth of %x is %s", addr, eth)
}
