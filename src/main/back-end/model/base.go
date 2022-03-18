package model

import (
	"encoding/json"
	"fmt"
	"gorm.io/gorm"
	"reflect"
)

/**
    model
    @author: roccoshi
    @desc: model before all
**/

// Models 所有的model, 用于gorm的auto migrate, 需要手动填写在这里
var Models = []interface{}{
	&User{},
	&Network{},
	&Wallet{},
	&Account{},
	&Transaction{},
}

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

// CommonScan for Scan()
func CommonScan(data interface{}, value interface{}) error {
	if value == nil {
		return nil
	}

	switch value.(type) {
	case []byte:
		return json.Unmarshal(value.([]byte), data)
	case string:
		return json.Unmarshal([]byte(value.(string)), data)
	default:
		return fmt.Errorf("val type is valid, is %+v", value)
	}
}

// CommonValue for Value()
func CommonValue(data interface{}) (interface{}, error) {
	vi := reflect.ValueOf(data)
	// 判断是否为 0 值
	if vi.IsZero() {
		return nil, nil
	}
	return json.Marshal(data)
}
