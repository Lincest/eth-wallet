package service

import (
	"back-end/model"
	"back-end/utils"
	"fmt"
	"log"
)

/**
    router
    @author: roccoshi
    @desc: user_service
**/

var User = &userService{}

type userService struct{}

func (srv *userService) AddUserByNameAndPassWord(name string, password string) error {
	// 查询是否存在
	if srv.GetUserByName(name) != nil {
		return fmt.Errorf("用户已存在")
	}
	// 新增
	if err := srv.AddUser(&model.User{Name: name, PassWord: utils.Encrypt.Md5encode(password)}); err != nil {
		return err
	}
	return nil
}

func (srv *userService) AddUser(user *model.User) error {
	result := db.Create(&user)
	if result.Error != nil {
		return result.Error
	}
	log.Printf("add user: rows affected = %v", result.RowsAffected)
	return nil
}

func (srv *userService) GetUserByName(name string) *model.User {
	ret := &model.User{Name: name}
	log.Print("Get user by name")
	if err := db.Where(ret).First(ret).Debug().Error; err != nil {
		return nil
	}
	return ret
}

func (srv *userService) GetUserByNameAndPassWord(name string, password string) *model.User {
	var ret = &model.User{Name: name, PassWord: utils.Encrypt.Md5encode(password)}
	log.Print("Get user by name")
	if err := db.Where(ret).First(ret).Debug().Error; err != nil {
		return nil
	}
	return ret
}
