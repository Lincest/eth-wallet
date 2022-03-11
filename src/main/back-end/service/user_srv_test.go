package service

import "testing"

/**
    service
    @author: roccoshi
    @desc: test
**/

func TestUserService_GetUserByName(t *testing.T) {
	if User.GetUserByName("Mike").Name != "Mike" {
		t.Fail()
	}
}

func TestUserService_GetUserByName2(t *testing.T) {
	if User.GetUserByName("aabbccNilNameNotExitaabbcc") != nil {
		t.Fail()
	}
}

func TestUserService_AddUserByNameAndPassWord(t *testing.T) {
	if User.GetUserByNameAndPassWord("Mike", "123").Name != "Mike" {
		t.Fail()
	}
}
