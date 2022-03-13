package model

/**
    model
    @author: roccoshi
    @desc: 用户model
**/

// User model.
type User struct {
	Model

	Name     string `gorm:"unique" json:"name" form:"name"` // 用户名 (唯一)
	PassWord string `json:"password" form:"password"`       // 密码
	Mnemonic string `json:"mnemonic" form:"mnemonic"`       // 助记词
}
