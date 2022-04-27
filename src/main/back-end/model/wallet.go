package model

/**
    model
    @author: roccoshi
    @desc: hd wallet
**/

// Wallet model.
type Wallet struct {
	Model

	UID                uint   `json:"uid" form:"uid"`                     // 用户ID
	BaseDerivationPath string `json:"base_derivation_path" form:"base_derivation_path"` // 衍生路径base  e.g. [m/44'/60'/0'/0]
	LastAccountIndex   uint   `json:"last_account_index" form:"last_account_index"`     // 衍生路径[base_derivation_path/account_index]的最后一个account_index
}
