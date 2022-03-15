package utils

import (
	"back-end/conf"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/shopspring/decimal"
	"math/big"
	"strconv"
	"strings"
)

/**
    utils
    @author: roccoshi
    @desc: 一系列以太坊的实用工具
**/

// Eth2Wei 将ETH转化为wei
func (*IWallet) Eth2Wei(iamount interface{}) *big.Int {
	amount := decimal.NewFromFloat(0)
	switch v := iamount.(type) {
	case int:
		amount = decimal.NewFromFloat(float64(v))
	case string:
		amount, _ = decimal.NewFromString(v)
	case float64:
		amount = decimal.NewFromFloat(v)
	case int64:
		amount = decimal.NewFromFloat(float64(v))
	case decimal.Decimal:
		amount = v
	case *decimal.Decimal:
		amount = *v
	}

	mul := decimal.NewFromFloat(float64(10)).Pow(decimal.NewFromFloat(float64(18)))
	result := amount.Mul(mul)

	wei := new(big.Int)
	wei.SetString(result.String(), 10)

	return wei
}

// Wei2Eth 将wei转化为ETH
func (*IWallet) Wei2Eth(ivalue interface{}) decimal.Decimal {
	value := new(big.Int)
	switch v := ivalue.(type) {
	case string:
		value.SetString(v, 10)
	case *big.Int:
		value = v
	}

	mul := decimal.NewFromFloat(float64(10)).Pow(decimal.NewFromFloat(float64(18)))
	num, _ := decimal.NewFromString(value.String())
	result := num.Div(mul)

	return result
}

// GetNewDerivationPath m/44'/60'/0'/0/0 => m/44'/60'/0'/0/1
func (*IWallet) GetNewDerivationPath(lastPath string) (string, error) {
	pathSplits := strings.Split(lastPath, "/")
	lastAccountIndex, err := strconv.Atoi(pathSplits[len(pathSplits)-1])
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%s/%d", conf.Config.Wallet.BasePath, lastAccountIndex+1), nil
}

// GetAddressHexFromPrivateKeyHex 根据私钥获取地址
func (*IWallet) GetAddressHexFromPrivateKeyHex(privateKeyHex string) (string, error) {
	ecdsaPrivatekey, err := crypto.HexToECDSA(privateKeyHex)
	if err != nil {
		return "", err
	}
	address := crypto.PubkeyToAddress(ecdsaPrivatekey.PublicKey)
	return address.Hex(), nil
}

// GetAddressFromPrivateKeyHex 根据私钥获取地址
func (*IWallet) GetAddressFromPrivateKeyHex(privateKeyHex string) (common.Address, error) {
	ecdsaPrivatekey, err := crypto.HexToECDSA(privateKeyHex)
	if err != nil {
		return common.Address{}, err
	}
	address := crypto.PubkeyToAddress(ecdsaPrivatekey.PublicKey)
	return address, nil
}
