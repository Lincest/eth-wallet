package service

import (
	"reflect"
	"testing"
)

/**
    service
    @author: roccoshi
    @desc: test
**/

func TestMnemonicService_NewFromMnemonic(t *testing.T) {
	// 对于相同助记词应该生成相同的钱包
	wallet1, _ := Mnemonic.NewFromMnemonic("tag volcano eight thank tide danger coast health above argue embrace heavy")
	wallet2, _ := Mnemonic.NewFromMnemonic("tag volcano eight thank tide danger coast health above argue embrace heavy")
	t.Logf("equal wallet1 and wallet2: %v", reflect.DeepEqual(wallet1, wallet2))
}

func TestMnemonicService_UpdateMnemonicByUser(t *testing.T) {
	testValidStr := "fantasy certain scale receive three start reform trade tape battle elder away"
	testInvalidStr := "invalid string"
	if err := Mnemonic.UpdateMnemonicByName(testValidStr, "Mike"); err != nil {
		t.Fail()
	}
	if err := Mnemonic.UpdateMnemonicByName(testInvalidStr, "Mike"); err == nil {
		t.Fail()
	}
	mnemonic, err := Mnemonic.GetMnemonicByName("Mike")
	if err != nil || mnemonic != testValidStr {
		t.Fail()
	}
	t.Logf("mnemonic of mike is [%s]", mnemonic)
}
