package service

import (
	"back-end/conf"
	"back-end/model"
	"back-end/utils"
	"context"
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	hdwallet "github.com/miguelmota/go-ethereum-hdwallet"
	"gorm.io/gorm"
)

/**
    service
    @author: roccoshi
    @desc: 钱包相关的service
**/

var Wallet = &walletService{}

type walletService struct{}

// IsValidUrlWithChainId 根据url和chain id查询网络是否存在
func (srv *walletService) IsValidUrlWithChainId(rawurl string, chainId string) error {
	client, err := ethclient.Dial(rawurl)
	if err != nil {
		return fmt.Errorf("建立连接失败, 请确认url的正确性")
	}
	iChainId, err := client.ChainID(context.Background())
	if iChainId == nil || err != nil {
		return fmt.Errorf("无法获取chain ID, 请确认url的正确性")
	}
	if iChainId.String() != chainId {
		return fmt.Errorf("chain ID不一致, RPC端点使用的chainID为: %s", iChainId.String())
	}
	defer client.Close()
	return nil
}

func (srv *walletService) GetAllNetWorkByUid(uid uint) ([]model.Network, error) {
	var networks []model.Network
	if err := db.Where("uid = ? or uid = ?", uid, 0).Find(&networks).Error; err != nil {
		return nil, err
	}
	return networks, nil
}

func (srv *walletService) AddOrUpdateNetWork(network model.Network) error {
	// 如果数据库中存在该id, 直接更新
	existNetwork := &model.Network{}
	err := db.First(&existNetwork, network.ID).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}
	//  err == nil说明找到了
	if err == nil {
		if existNetwork.UID != network.UID {
			return fmt.Errorf("没有权限修改")
		}
		fmt.Printf("%d uid -- uid %d", existNetwork.UID, network.UID)
		if err := Wallet.UpdateNetWork(network); err != nil {
			return err
		}
		return nil
	}
	// 如果数据库中不存在, 就新增
	if err := db.Create(&network).Error; err != nil {
		return err
	}
	return nil
}

func (srv *walletService) GetNetWorkByID(id uint) (model.Network, error) {
	network := model.Network{}
	if err := db.First(&network, id).Error; err != nil {
		return network, err
	}
	return network, nil
}

func (srv *walletService) DeleteNetWork(network model.Network) error {
	if err := db.Delete(&network).Error; err != nil {
		return err
	}
	return nil
}

func (srv *walletService) UpdateNetWork(network model.Network) error {
	if err := db.Save(&network).Error; err != nil {
		return err
	}
	return nil
}

// InitWallet 为用户新建一个钱包
func (srv *walletService) InitWallet(uid uint) error {
	// 检查用户是否有助记词
	mnemonic, err := Mnemonic.GetMnemonicByUid(uid)
	if err != nil {
		return err
	}
	if mnemonic == "" {
		return fmt.Errorf("用户未创建助记词")
	}
	// 检查用户是否已经创建钱包
	iwallet := &model.Wallet{UID: uid}
	if err := db.Where(iwallet).First(iwallet).Error; err == nil {
		return fmt.Errorf("用户已经创建钱包, 最新的衍生路径为 = %s", fmt.Sprintf("%s/%d", iwallet.BaseDerivationPath, iwallet.LastAccountIndex))
	}
	// 为用户新建钱包
	basePath := fmt.Sprintf("%s/%d", conf.Config.Wallet.BasePath, 0)
	account, privateKeyHex, err := srv.GenerateNewAccount(mnemonic, basePath)
	if err != nil {
		return err
	}
	newWallet := model.Wallet{UID: uid, BaseDerivationPath: conf.Config.Wallet.BasePath, LastAccountIndex: 0}
	// 开启事务: 1. 新建钱包 2. 新增账户
	tx := db.Begin()
	if err := tx.Create(&newWallet).Error; err != nil {
		tx.Rollback()
		return err
	}
	newAccount := model.Account{UID: uid, DerivationPath: basePath, Address: account.Address, PrivateKeyHex: privateKeyHex}
	if err := tx.Create(&newAccount).Error; err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}

func (srv *walletService) AddNewAccount(uid uint) error {
	wallet := model.Wallet{UID: uid}
	// 没有钱包时, 首先新增钱包
	if err := db.First(&wallet).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			if err = srv.InitWallet(uid); err != nil {
				return err
			}
			return nil
		}
	}
	// 如果有钱包了, 就根据最新衍生路径新建账户
	lastPath := fmt.Sprintf("%s/%d", wallet.BaseDerivationPath, wallet.LastAccountIndex)
	newPath, err := utils.Wallet.GetNewDerivationPath(lastPath)
	if err != nil {
		return err
	}
	mnemonic, err := Mnemonic.GetMnemonicByUid(uid)
	if err != nil {
		return err
	}
	if mnemonic == "" {
		return fmt.Errorf("用户未创建助记词")
	}
	newAccount, privateKeyHex, err := srv.GenerateNewAccount(mnemonic, newPath)
	if err != nil {
		return err
	}
	tx := db.Begin()
	if err := tx.Create(&model.Account{UID: uid, DerivationPath: newPath, Address: newAccount.Address, PrivateKeyHex: privateKeyHex}).Error; err != nil {
		tx.Rollback()
		return err
	}
	wallet.LastAccountIndex += 1
	if err := tx.Save(&wallet).Error; err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}

func (srv *walletService) AddNewAccountByUIDAndAddressAndPrivateKey(uid uint, address common.Address, privateKey string) error {
	account := model.Account{UID: uid, Address: address, PrivateKeyHex: privateKey}
	if err := db.Create(&account).Error; err != nil {
		return err
	}
	return nil
}

func (srv *walletService) AddNewAccountByUIDAndPrivateKey(uid uint, privateKey string) error {
	address, err := utils.Wallet.GetAddressFromPrivateKeyHex(privateKey)
	if err != nil {
		return err
	}
	if err := srv.AddNewAccountByUIDAndAddressAndPrivateKey(uid, address, privateKey); err != nil {
		return err
	}
	return nil
}

func (srv *walletService) GetAllAccountsByUID(uid uint) ([]model.Account, error) {
	var accountsRes []model.Account
	if err := db.Where("uid = ?", uid).Find(&accountsRes).Error; err != nil {
		return accountsRes, err
	}
	return accountsRes, nil
}

// GenerateNewAccount 由衍生路径和助记词创建新账户
// 返回账户和账户私钥
func (srv *walletService) GenerateNewAccount(mnemonic string, path string) (*accounts.Account, string, error) {
	wallet, err := hdwallet.NewFromMnemonic(mnemonic)
	if err != nil {
		return nil, "", err
	}
	derivePath, err := hdwallet.ParseDerivationPath(path)
	if err != nil {
		return nil, "", err
	}
	account, err := wallet.Derive(derivePath, false)
	if err != nil {
		return nil, "", err
	}
	privateKeyHex, err := wallet.PrivateKeyHex(account)
	if err != nil {
		return nil, "", err
	}
	return &account, privateKeyHex, nil
}
