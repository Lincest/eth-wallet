package model

import "github.com/ethereum/go-ethereum/common"

/**
    model
    @author: roccoshi
    @desc: transaction
**/

type Transaction struct {
	Model

	UID         uint           `gorm:"not null" json:"uid" form:"uid"`                            // 用户uid
	Hash        string         `gorm:"size:255;not null;unique" json:"hash" form:"hash"`          // 交易哈希值
	Value       string         `gorm:"size:255;not null" json:"value" form:"value"`               // 转账金额 (wei)
	GasPrice    string         `gorm:"size:255;not null" json:"gas_price" form:"gas_price"`       // gas price (wei)
	GasLimit    string         `gorm:"size:255;not null" json:"gas_limit" form:"gas_limit"`       // gas limit (gas 单位)
	Nonce       string         `gorm:"size:255;not null" json:"nonce" form:"nonce"`               // nonce值
	FromAddress common.Address `gorm:"size:255;not null" json:"from_address" form:"from_address"` // 源地址
	ToAddress   common.Address `gorm:"size:255;not null" json:"to_address" form:"to_address"`     // 目的地址
	GasUsed     string         `gorm:"size:255" json:"gas_used" form:"gas_used"`                  // gas used (实际使用的gas 单位)
	Cost        string         `gorm:"size:255" json:"cost" form:"cost"`                          // 总消费 value + gasused * gasprice (wei)
	Status      bool           `json:"status" form:"status"`                                      // 交易状态 成功 / 失败
	IsPending   bool           `json:"is_pending" form:"is_pending"`                              // 交易是否正在等待
	BlockNumber string         `gorm:"size:255" json:"block_number" form:"block_number"`          // block number
	BlockHash   string         `gorm:"size:255" json:"block_hash" form:"block_hash"`              // block hash
	Network     string         `gorm:"size:255" json:"network" form:"network"`                    // network raw url e.g. http://localhost:7545
}

// TransactionReq create transaction request
type TransactionReq struct {
	UID               uint   `json:"uid" form:"uid"`                                   // 用户uid
	FromAddress       string `json:"from_address" form:"from_address"`                 // 源地址
	FromPrivateKeyHex string `json:"from_private_key_hex" form:"from_private_key_hex"` // 源私钥
	ToAddress         string `json:"to_address" form:"to_address"`                     // 目的地址
	Value             string `json:"value" form:"value"`                               // 转账金额 (wei)
	GasPrice          string `json:"gas_price" form:"gas_price"`                       // gas price (wei)
	GasLimit          string `json:"gas_limit" form:"gas_limit"`                       // gas limit (gas 单位)
}
