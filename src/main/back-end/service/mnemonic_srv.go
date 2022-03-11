package service

/**
    wallet
    @author: roccoshi
    @desc: 助记词相关操作
**/

import (
	"github.com/miguelmota/go-ethereum-hdwallet"
)

var Mnemonic = &mnemonicService{}

type mnemonicService struct {
	wallet *hdwallet.Wallet
}

func (srv *mnemonicService) NewFromMnemonic(mnemonic string) (*hdwallet.Wallet, error) {
	wallet, err := hdwallet.NewFromMnemonic(mnemonic)
	if err != nil {
		return nil, err
	}
	srv.wallet = wallet
	return wallet, nil
}
