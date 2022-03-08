package utils

import (
	"github.com/ethereum/go-ethereum/common"
	"testing"
)

/**
    utils
    @author: roccoshi
    @desc: //TODO
**/

func TestIWallet_IsValidAddress(t *testing.T) {
	validAddr := "0x5e0055403ABF9F373c39a5714b8Fc7B2A9d02aAe"
	if Wallet.IsValidAddress(validAddr) != true {
		t.Fail()
	}
	validAddr1 := common.HexToAddress(validAddr)
	if Wallet.IsValidAddress(validAddr1) != true {
		t.Fail()
	}
	inValidAddr := "0x5e0055403ABF9F373c39a5714b8Fc7B2A9d02Ae"
	if Wallet.IsValidAddress(inValidAddr) == true {
		t.Fail()
	}
}

func TestIWallet_IsZeroAddress(t *testing.T) {
	zeroAddr := "0x0"
	if Wallet.IsZeroAddress(zeroAddr) != true {
		t.Fail()
	}
	inValidAddr := "0x5e0055403ABF9F373c39a5714b8Fc7B2A9d02aAe"
	if Wallet.IsZeroAddress(inValidAddr) == true {
		t.Fail()
	}
}
