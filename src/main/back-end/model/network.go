package model

/**
    model
    @author: roccoshi
    @desc: 以太坊存储的所有网络
**/

type Network struct {
	Model `json:"-"`

	Name    string `gorm:"unique" json:"name" form:"name"`
	Url     string `json:"url" form:"url"`
	ChainId string `json:"chain_id" form:"chain_id"`
	UID     uint   `json:"-"` // 用户uid
}
