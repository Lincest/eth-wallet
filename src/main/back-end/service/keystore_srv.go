package service

import (
	"back-end/utils"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"io/ioutil"
	"log"
	"mime/multipart"
	"os"
	"path/filepath"
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
	address, err := utils.Wallet.GetAddressHexFromPrivateKeyHex(privateKeyHex)
	if err != nil {
		return "", "", err
	}
	return privateKeyHex, address, nil
}

// GenerateOneKeyStoreFile 根据私钥和密钥生成对应的一个keystore文件, 文件夹以当前时间戳命名
// 返回文件夹路径
func (srv *keystoreService) GenerateOneKeyStoreFile(privateKey string, passphrase string) (string, error) {
	nowStampStr := strconv.FormatInt(time.Now().Unix(), 10)
	fileDir := fmt.Sprintf("%s/%s", baseFileDir, nowStampStr)
	ks := keystore.NewKeyStore(fileDir, keystore.LightScryptN, keystore.LightScryptP)
	ecdsaPrivateKey, err := crypto.HexToECDSA(privateKey)
	if err != nil {
		return fileDir, err
	}
	_, err = ks.ImportECDSA(ecdsaPrivateKey, passphrase)
	if err != nil {
		return fileDir, err
	}
	return fileDir, nil
}

// GenerateKeyStoreFiles 根据一组私钥和密钥生成对应的多个keystore文件, 文件夹以当前时间戳命名
// 返回文件夹路径
func (srv *keystoreService) GenerateKeyStoreFiles(privateKeyList []string, passphrase string) (string, error) {
	nowStampStr := strconv.FormatInt(time.Now().Unix(), 10)
	fileDir := fmt.Sprintf("%s/%s", baseFileDir, nowStampStr)
	ks := keystore.NewKeyStore(fileDir, keystore.LightScryptN, keystore.LightScryptP)
	for _, privateKey := range privateKeyList {
		ecdsaPrivateKey, err := crypto.HexToECDSA(privateKey)
		if err != nil {
			return fileDir, err
		}
		_, err = ks.ImportECDSA(ecdsaPrivateKey, passphrase)
		if err != nil {
			return fileDir, err
		}
	}
	return fileDir, nil
}

// AddOneAccountByKeyStoreFile 通过keystore文件和密钥添加一个账户
func (srv *keystoreService) AddOneAccountByKeyStoreFile(file *multipart.FileHeader, uid uint, passphrase string) error {
	keystoreFile, err := file.Open()
	if err != nil {
		return err
	}
	keystoreBytes, err := ioutil.ReadAll(keystoreFile)
	if err != nil {
		return err
	}
	privateKey, addr, err := srv.LoadOneKeyStore(keystoreBytes, passphrase)
	if err != nil {
		return err
	}
	address := common.HexToAddress(addr)
	err = Wallet.AddNewAccountByUIDAndAddressAndPrivateKey(uid, address, privateKey)
	if err != nil {
		return err
	}
	return nil
}

// GetAllKeyStoreFilesByUID 导出uid对应用户所有的私钥为keystore文件
func (srv *keystoreService) GetAllKeyStoreFilesByUID(uid uint, passphrase string) (string, error) {
	log.Printf("uid = %d 正在以密码 = %s导出keystore文件\n", uid, passphrase)
	accountList, err := Wallet.GetAllAccountsByUID(uid)
	if err != nil {
		return "", err
	}
	privateKeyList := make([]string, 0)
	for _, account := range accountList {
		privateKeyList = append(privateKeyList, account.PrivateKeyHex)
	}
	fileDir, err := srv.GenerateKeyStoreFiles(privateKeyList, passphrase)
	if err != nil {
		return "", err
	}
	tmpDir := os.TempDir()
	zipFilePath := filepath.Join(tmpDir, "keystore.zip")
	fmt.Printf("zipfilepath = %s", zipFilePath)
	zipFile, err := utils.Zip.Create(zipFilePath)
	if err != nil {
		return "", err
	}
	err = zipFile.AddDirectory(".", fileDir)
	if err != nil {
		return "", err
	}
	if err := zipFile.Close(); err != nil {
		return "", fmt.Errorf("zip failed, %s", err.Error())
	}
	os.RemoveAll(baseFileDir)
	return zipFilePath, nil
}
