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

	UID            uint           `json:"uid" form:"uid"`                         // 用户ID
	DerivationPath string         `json:"derivation_path" form:"derivation_path"` // 衍生路径 e.g. [m/44'/60'/0'/0/1]
	Address        common.Address `json:"address" form:"address"`                 // 地址
	PrivateKeyHex  string         `json:"private_key_hex" form:"private_key_hex"` // 私钥
}
