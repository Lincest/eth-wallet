package utils

import (
	"github.com/shopspring/decimal"
	"math/big"
	"testing"
)

/**
    utils
    @author: roccoshi
    @desc: 测试eth和wei互转
**/

func TestIWallet_Eth2Wei(t *testing.T) {
	bigValue := new(big.Int)
	bigValue.SetString("1000000000000000000", 10)
	if Wallet.Eth2Wei(1).String() != bigValue.String() {
		t.Log("1")
		t.Fail()
	}
	if Wallet.Eth2Wei(1.00).String() != bigValue.String() {
		t.Log("2")
		t.Fail()
	}
	if Wallet.Eth2Wei("1").String() != bigValue.String() {
		t.Log("3")
		t.Fail()
	}
	if Wallet.Wei2Eth(bigValue).String() != decimal.NewFromInt(1).String() {
		t.Log("4")
		t.Fail()
	}
}

func TestIWallet_GetNewDerivationPath(t *testing.T) {
	oldPath := "m/44'/60'/0'/0/98"
	derivivationPath, err := Wallet.GetNewDerivationPath(oldPath)
	if err != nil {
		t.Fail()
	}
	t.Logf("new of %s is %s", oldPath, derivivationPath.String())
}
