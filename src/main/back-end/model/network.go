package model

/**
    model
    @author: roccoshi
    @desc: 以太坊存储的所有网络
**/

type Network struct {
	Model

	Name    string `gorm:"size:255;not null" json:"name" form:"name"`
	Url     string `gorm:"size:255;not null" json:"url" form:"url"`
	ChainId string `gorm:"size:255;not null" json:"chain_id" form:"chain_id"`
	UID     uint   `gorm:"not null" json:"uid" form:"uid"` // 用户uid
}
