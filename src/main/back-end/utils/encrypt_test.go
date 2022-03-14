package utils

import "testing"

/**
    utils
    @author: roccoshi
    @desc: test
**/

func TestIEncrypt_Md5encode(t *testing.T) {
	t.Logf(Encrypt.Md5encode("helloworld"))
	if Encrypt.Md5encode("helloworld") != "fc5e038d38a57032085441e7fe7010b0" {
		t.Fail()
	}
}

func TestIEncrypt_Rc4encode(t *testing.T) {
	origin := "hello world this is eth-wallet"
	password := "123"
	secret, _ := Encrypt.Rc4encode(origin, password)
	secret1, _ := Encrypt.Rc4encode(secret, password)
	t.Logf("加密后: %s", secret)
	t.Logf("解密后: %s", secret1)
	if secret1 != origin {
		t.Fail()
	}
}
