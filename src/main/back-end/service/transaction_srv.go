package service

import (
	"back-end/model"
	"back-end/utils"
	"context"
	"crypto/ecdsa"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"math/big"
	"strconv"
)

/**
    service
    @author: roccoshi
    @desc: 交易管理
**/

var Transaction = &transactionService{}

type transactionService struct{}

// CreateTransaction 创建新交易
// transferValue 发送金额 wei
// gasPrice wei
// network url e.g. http://localhost:7545
//
// return: 本交易哈希值
func (srv *transactionService) CreateTransaction(uid uint, fromAddress common.Address,
	fromPrivateKeyHex string, toAddress common.Address, transferValue string, gasLimit string, gasPrice string, network string) (string, error) {
	// 1 - 检查地址合法性
	if !utils.Wallet.IsValidAddress(fromAddress) {
		return "", fmt.Errorf("from address 不合法")
	}
	if !utils.Wallet.IsValidAddress(toAddress) {
		return "", fmt.Errorf("to address 不合法")
	}
	// 2 - 连接客户端
	client, err := ethclient.Dial(network)
	if err != nil {
		return "", err
	}
	defer client.Close()
	// 3 - 加载私钥
	privKey, err := crypto.HexToECDSA(fromPrivateKeyHex)
	if err != nil {
		return "", err
	}
	// 4 - 从私钥 -> 公钥 -> fromAddress
	pubKey := privKey.Public()
	publicKeyECDSA, ok := pubKey.(*ecdsa.PublicKey)
	if !ok {
		return "", fmt.Errorf("pubKey types error")
	}
	fromAddressFromPrivKey := crypto.PubkeyToAddress(*publicKeyECDSA)
	if fromAddress.Hex() != fromAddressFromPrivKey.Hex() {
		return "", fmt.Errorf("传入from address和从私钥推出的address不一致, 请检查")
	}
	// 5 - 获取应该用于交易的nonce值
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		return "", err
	}
	// 6 - 转账金额
	transVal := big.NewInt(0)
	transVal.SetString(transferValue, 10)
	// 7 - 设定gas费用
	gasLimitInt, err := strconv.Atoi(gasLimit)
	if err != nil {
		return "", err
	}
	gasLimitUint := uint64(gasLimitInt)
	gasPriceVal := big.NewInt(0)
	gasPriceVal.SetString(gasPrice, 10)
	// 8 - 生成未签名的Transaction
	tx := types.NewTransaction(nonce, toAddress, transVal, gasLimitUint, gasPriceVal, nil)
	// 9 - 使用私钥对Transaction进行签名
	// 首先获取chainID
	chainID, err := client.ChainID(context.Background())
	if err != nil {
		return "", err
	}
	// 对transaction进行签名, 用EIP155签名
	// signer: 验证和处理签名
	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privKey)
	if err != nil {
		return "", err
	}
	// 8 - 利用SendTransaction将已签名的事务广播到整个网络
	if err := client.SendTransaction(context.Background(), signedTx); err != nil {
		return "", err
	}
	fmt.Printf("Transaction Hex: (%#v) 已经被广播", signedTx.Hash().Hex())
	// 9 - 入库
	txHash := signedTx.Hash().Hex()
	newTransaction := model.Transaction{
		UID:         uid,
		Hash:        txHash,
		Value:       transferValue,
		GasLimit:    gasLimit,
		GasPrice:    gasPrice,
		Nonce:       strconv.FormatUint(nonce, 10),
		FromAddress: fromAddress,
		ToAddress:   toAddress,
	}
	if err := db.Create(&newTransaction).Error; err != nil {
		return "", err
	}
	return txHash, nil
}

// GetAndUpdateTransactionByHash 检查并更新交易状态
func (srv *transactionService) GetAndUpdateTransactionByHash(transactionHash string, network string) (*model.Transaction, error) {
	// 数据库查找该交易
	transaction := &model.Transaction{Hash: transactionHash}
	if err := db.Where(transaction).First(transaction).Error; err != nil {
		return transaction, err
	}
	// 连接
	client, err := ethclient.Dial(network)
	if err != nil {
		return transaction, err
	}
	defer client.Close()
	// 查询交易
	tx, isPending, err := client.TransactionByHash(context.Background(), common.HexToHash(transactionHash))
	if err != nil {
		return transaction, err
	}
	receipt, err := client.TransactionReceipt(context.Background(), tx.Hash())
	if err != nil {
		return transaction, err
	}
	transaction.IsPending = isPending
	transaction.GasUsed = strconv.FormatUint(tx.Gas(), 10)
	transaction.Cost = tx.Cost().String()
	if !isPending {
		transaction.Status = receipt.Status == uint64(1)
		transaction.BlockHash = receipt.BlockHash.String()
		transaction.BlockNumber = receipt.BlockNumber.String()
	}
	if err := db.Save(&transaction).Error; err != nil {
		return transaction, err
	}
	return transaction, nil
	// below is for debug... ============================================================
	//fmt.Printf("transaction %#v, ispending: %v\n", tx, isPending)
	//fmt.Println("是否正在等待: ", isPending)
	//fmt.Println("交易hash: ", tx.Hash().Hex())
	//fmt.Println("转账金额", tx.Value().String())
	//fmt.Println("Gas Limit: ", tx.Gas())
	//fmt.Println("Gas Price: ", tx.GasPrice().Uint64())
	//fmt.Println("Nonce值: ", tx.Nonce())
	//fmt.Println("Data: ", tx.Data())
	//fmt.Println("总花费(value + gaslimit * gasprice)", tx.Cost().Uint64())
	//fmt.Println("To地址: ", tx.To().Hex())
	//chainID, _ := client.ChainID(context.Background())
	//if msg, err := tx.AsMessage(types.NewEIP155Signer(chainID), nil); err == nil {
	//	fmt.Println("from地址: ", msg.From().Hex())
	//}
	//fmt.Println("收据gas used: ", receipt.GasUsed)
	//fmt.Println("收据status: ", receipt.Status)
	//fmt.Println("收据block number: ", receipt.BlockNumber)
	//fmt.Println("收据block hash: ", receipt.BlockHash)
	// above is for debug... ============================================================
}
