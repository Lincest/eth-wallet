package utils

import (
	"github.com/ethereum/go-ethereum/common"
	"reflect"
	"regexp"
)

/**
    wallet
    @author: roccoshi
    @desc: 一系列以太坊的检查相关工具
**/

// IsValidAddress 检查一个以太坊地址是否合法
// 输入类型可以是string或者common.Address类型
func (*IWallet) IsValidAddress(address interface{}) bool {
	re := regexp.MustCompile("^0x[0-9a-fA-F]{40}$")
	switch v := address.(type) {
	case string:
		return re.MatchString(v)
	case common.Address:
		return re.MatchString(v.Hex())
	default:
		return false
	}
}

// IsZeroAddress 检查是否为0地址
func (*IWallet) IsZeroAddress(address interface{}) bool {
	var addr common.Address
	switch v := address.(type) {
	case string:
		addr = common.HexToAddress(v)
	case common.Address:
		addr = v
	default:
		return false
	}

	zeroAddressBytes := common.FromHex("0x0000000000000000000000000000000000000000")
	addressBytes := addr.Bytes()
	return reflect.DeepEqual(addressBytes, zeroAddressBytes)
}
