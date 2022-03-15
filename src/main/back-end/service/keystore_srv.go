package service

import (
	"back-end/utils"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/crypto"
	"log"
	"strconv"
	"time"
)

/**
    service
    @author: roccoshi
    @desc: 实现keystore的导入导出管理
**/

var KeyStore = &keystoreService{}

type keystoreService struct{}

const (
	baseFileDir = "./keystore_storage"
)

// LoadOneKeyStore 输入keystore文件和密钥, 返回解析后的私钥和地址
// return: private key hex, address, error
func (srv *keystoreService) LoadOneKeyStore(keystoreFile []byte, passphrase string) (string, string, error) {
	privateKeyByte, err := keystore.DecryptKey(keystoreFile, passphrase)
	if err != nil {
		return "", "", err
	}
	privateKeyHex := fmt.Sprintf("%x", crypto.FromECDSA(privateKeyByte.PrivateKey))
	if err != nil {
		return "", "", err
	}
	address, err := utils.Wallet.GetAddressFromPrivateKeyHex(privateKeyHex)
	if err != nil {
		return "", "", err
	}
	return privateKeyHex, address, nil
}

// GenerateOneKeyStoreFile 根据私钥和密钥生成对应的一个keystore文件, 文件夹以当前时间戳命名
func (srv *keystoreService) GenerateOneKeyStoreFile(privateKey string, passphrase string) error {
	nowStampStr := strconv.FormatInt(time.Now().Unix(), 10)
	fileDir := fmt.Sprintf("%s/%s", baseFileDir, nowStampStr)
	ks := keystore.NewKeyStore(fileDir, keystore.LightScryptN, keystore.LightScryptP)
	ecdsaPrivateKey, err := crypto.HexToECDSA(privateKey)
	if err != nil {
		log.Fatal(err)
	}
	_, err = ks.ImportECDSA(ecdsaPrivateKey, passphrase)
	if err != nil {
		return err
	}
	return nil
}

// GenerateKeyStoreFiles 根据一组私钥和密钥生成对应的多个keystore文件, 文件夹以当前时间戳命名
func (srv *keystoreService) GenerateKeyStoreFiles(privateKeyList []string, passphrase string) error {
	nowStampStr := strconv.FormatInt(time.Now().Unix(), 10)
	fileDir := fmt.Sprintf("%s/%s", baseFileDir, nowStampStr)
	ks := keystore.NewKeyStore(fileDir, keystore.LightScryptN, keystore.LightScryptP)
	for _, privateKey := range privateKeyList {
		ecdsaPrivateKey, err := crypto.HexToECDSA(privateKey)
		if err != nil {
			log.Fatal(err)
		}
		_, err = ks.ImportECDSA(ecdsaPrivateKey, passphrase)
		if err != nil {
			return err
		}
	}
	return nil
}
