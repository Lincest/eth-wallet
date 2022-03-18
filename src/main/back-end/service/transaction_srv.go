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
	"log"
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
		Network:     network,
	}
	if err := db.Create(&newTransaction).Error; err != nil {
		return "", err
	}
	return txHash, nil
}

// GetAndUpdateTransactionByHash 检查并更新交易状态
func (srv *transactionService) GetAndUpdateTransactionByHash(transactionHash string, network string) (*model.Transaction, error) {
	log.Print("find transaction: ", transactionHash)
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
	// 如果在pending就不查收据直接返回
	if isPending {
		transaction.IsPending = isPending
		transaction.GasUsed = strconv.FormatUint(tx.Gas(), 10)
		transaction.Cost = tx.Cost().String()
		if err := db.Save(&transaction).Error; err != nil {
			return transaction, err
		}
		return transaction, nil
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

// AccelerateTransaction 根据transaction的id和新的gasPrice加速(更新)transaction
// 原理: 更改gas price, 并用原来的nonce值重新将交易上链
func (srv *transactionService) AccelerateTransaction(id uint, newGasPrice string, uid uint) error {
	transaction := model.Transaction{}
	if err := db.First(&transaction, id).Error; err != nil {
		return err
	}
	if transaction.UID != uid {
		return fmt.Errorf("没有权限修改")
	}
	client, err := ethclient.Dial(transaction.Network)
	if err != nil {
		return err
	}
	defer client.Close()
	// 查询交易
	_, isPending, err := client.TransactionByHash(context.Background(), common.HexToHash(transaction.Hash))
	if !isPending {
		return fmt.Errorf("交易已经完成, 不可修改")
	}
	nonceInt, err := strconv.Atoi(transaction.Nonce)
	if err != nil {
		return err
	}
	nonce := uint64(nonceInt)
	gasLimitInt, err := strconv.Atoi(transaction.GasLimit)
	if err != nil {
		return err
	}
	gasLimit := uint64(gasLimitInt)
	transVal := big.NewInt(0)
	transVal.SetString(transaction.Value, 10)
	gasPrice := big.NewInt(0)
	gasPrice.SetString(newGasPrice, 10)
	newTx := types.NewTransaction(nonce, transaction.ToAddress, transVal, gasLimit, gasPrice, nil)
	fmt.Printf("cost: %v \n", utils.Wallet.Wei2Eth(newTx.Cost()))
	chainID, err := client.ChainID(context.Background())
	if err != nil {
		return err
	}
	account := model.Account{Address: transaction.FromAddress}
	if err := db.Where(&account).First(&account).Error; err != nil {
		return err
	}
	privKey, err := crypto.HexToECDSA(account.PrivateKeyHex)
	if err != nil {
		return err
	}
	signedTx, err := types.SignTx(newTx, types.NewEIP155Signer(chainID), privKey)
	if err != nil {
		return err
	}
	// 8 - 利用SendTransaction将已签名的事务广播到整个网络
	if err := client.SendTransaction(context.Background(), signedTx); err != nil {
		return err
	}
	fmt.Printf("Transaction Hex: (%#v) 已经被广播", signedTx.Hash().Hex())
	transaction.Hash = signedTx.Hash().Hex() // 更新后哈希值和gasPrice是改变量
	transaction.GasPrice = newGasPrice
	if err := db.Save(&transaction).Error; err != nil {
		return err
	}
	return nil
}
