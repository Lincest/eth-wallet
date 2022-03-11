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