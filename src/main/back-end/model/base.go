package model

import "gorm.io/gorm"

/**
    model
    @author: roccoshi
    @desc: model before all
**/

// Model represents meta data of entity.
type Model struct {
	gorm.Model
	/**
	gorm.Model包括:
	ID        uint `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt DeletedAt `gorm:"index"`
	*/
}

// Models 所有的model, 用于gorm的auto migrate, 需要手动填写在这里
var Models = []interface{}{
	&User{},
	&Network{},
}
