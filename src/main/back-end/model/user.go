package model

/**
    model
    @author: roccoshi
    @desc: 用户model
**/

// User model.
type User struct {
	Model

	Name     string `gorm:"size:32" json:"name"`
	PassWord string `gorm:"size:255" json:"password"`
}
