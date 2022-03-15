package service

import (
	"io/ioutil"
	"testing"
)

/**
    service
    @author: roccoshi
    @desc: test
**/

func TestKeystoreService_LoadOneKeyStore(t *testing.T) {
	dir := "./keystore_storage/1647339987"
	rd, err := ioutil.ReadDir(dir)
	if err != nil {
		t.Fail()
	}
	for _, file := range rd {
		t.Logf("filename: %s", file.Name())
		jsonBytes, _ := ioutil.ReadFile(dir + "/" + file.Name())
		priv, addr, err := KeyStore.LoadOneKeyStore(jsonBytes, "123")
		if err != nil {
			t.Error(err)
		}
		t.Logf("privatekey = %v, address = %v", priv, addr)
	}
}

func TestKeystoreService_GenerateOneKeyStoreFile(t *testing.T) {
	if _, err := KeyStore.GenerateOneKeyStoreFile("ee6030fcdcb30fbc3459456e7ef820c4c5c78e24fac11a2c886edec0bf1fb46f", "123"); err != nil {
		t.Fail()
	}
}

func TestKeystoreService_GenerateKeyStoreFiles(t *testing.T) {
	privateKeyList := []string{"dd0a7622038ce49d5e42218b9bdfe06073ec0f4cc13ddad794796681a6b798d6", "1e16ec571068634249989a5f26512c40e5776ad3f362f737f7f4844b45bb26a2", "a3bceab329d6943d848f868919300c1e4e3ba4fc945918ee46a3a4cb1657f6fd"}
	if _, err := KeyStore.GenerateKeyStoreFiles(privateKeyList, "123"); err != nil {
		t.Fail()
	}
}
