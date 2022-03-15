package service

/**
    wallet
    @author: roccoshi
    @desc: 助记词相关操作
**/

import (
	"back-end/model"
	"fmt"
	"github.com/miguelmota/go-ethereum-hdwallet"
	"github.com/tyler-smith/go-bip39"
)

var Mnemonic = &mnemonicService{}

type mnemonicService struct{}

func (srv *mnemonicService) NewFromMnemonic(mnemonic string) (*hdwallet.Wallet, error) {
	wallet, err := hdwallet.NewFromMnemonic(mnemonic)
	if err != nil {
		return nil, err
	}
	return wallet, nil
}

// UpdateMnemonicByName 更新用户助记词
func (srv *mnemonicService) UpdateMnemonicByName(mnemonic string, name string) error {
	// 检查助记词合法性
	if ok := bip39.IsMnemonicValid(mnemonic); !ok {
		return fmt.Errorf("助记词[%s]不合法", mnemonic)
	}
	if err := db.Model(&model.User{}).Where("name = ?", name).Update("mnemonic", mnemonic).Error; err != nil {
		return fmt.Errorf("更新助记词失败: %s", err)
	}
	return nil
}

func (srv *mnemonicService) GetMnemonicByName(name string) (string, error) {
	ret := &model.User{Name: name}
	if err := db.Where(ret).First(ret).Error; err != nil {
		return "", err
	}
	return ret.Mnemonic, nil
}

func (srv *mnemonicService) GetMnemonicByUid(uid uint) (string, error) {
	ret := model.User{}
	if err := db.First(&ret, uid).Error; err != nil {
		return "", err
	}
	return ret.Mnemonic, nil
}
