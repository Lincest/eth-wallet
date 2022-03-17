package model

import (
	"github.com/ethereum/go-ethereum/common"
)

/**
    model
    @author: roccoshi
    @desc: 单条address
**/

type Account struct {
	Model

	UID            uint           `json:"uid" form:"uid"`                                                // 用户ID
	DerivationPath string         `gorm:"size:255" json:"derivation_path" form:"derivation_path"`        // 衍生路径 e.g. [m/44'/60'/0'/0/1]
	Address        common.Address `gorm:"size:255;unique" json:"address" form:"address"`                 // 地址
	PrivateKeyHex  string         `gorm:"size:255;unique" json:"private_key_hex" form:"private_key_hex"` // 私钥
}

type AccountResp struct {
	ID             uint           `json:"ID" form:"ID"`                           // Model.ID
	UID            uint           `json:"uid" form:"uid"`                         // 用户ID
	DerivationPath string         `json:"derivation_path" form:"derivation_path"` // 衍生路径 e.g. [m/44'/60'/0'/0/1]
	Address        common.Address `json:"address" form:"address"`                 // 地址
	PrivateKeyHex  string         `json:"private_key_hex" form:"private_key_hex"` // 私钥
	Balance        string         `json:"balance" form:"balance"`                 // 用户余额 (单位: Wei)
}

func Account2AccountResp(account *Account) AccountResp {
	return AccountResp{
		ID:             account.ID,
		UID:            account.UID,
		DerivationPath: account.DerivationPath,
		Address:        account.Address,
		PrivateKeyHex:  account.PrivateKeyHex,
		Balance:        "",
	}
}
