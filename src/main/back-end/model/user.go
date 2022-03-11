package model

/**
    model
    @author: roccoshi
    @desc: 用户model
**/

// User model.
type User struct {
	Model

	Name     string `gorm:"size:255" json:"name" form:"name"`
	PassWord string `gorm:"size:255" json:"password" form:"password"`
}
