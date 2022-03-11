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
