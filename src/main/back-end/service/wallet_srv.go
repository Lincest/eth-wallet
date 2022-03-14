package service

import (
	"back-end/model"
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/ethclient"
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
	//  err == nil说明找到了
	if err == nil {
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
