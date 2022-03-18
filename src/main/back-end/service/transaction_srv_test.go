package service

import (
	"github.com/ethereum/go-ethereum/common"
	"testing"
)

/**
    service
    @author: roccoshi
    @desc: 测试交易
**/

func TestTransactionService_CreateTransaction(t *testing.T) {
	fromAddress := common.HexToAddress("0xBAe52719E15Ab84564183b72a0299966727EC162") // index = 10
	fromPrivateKeyHex := "41794b550479390d8fff0e179655c33dada068b575751520575375e008ba624b"
	toAddress := common.HexToAddress("0x342a409070DBE2e2CDA157F77C581ece17B09795") // index = 11
	transFerValue := "10130000000000000000"                                        // 10.13 ETH
	gasPrice := "20000000000"                                                      // 0.00000002 ETH
	gasLimit := "21000"                                                            // gasLimit * gasPrice = 0.00042 ETH
	network := "http://localhost:7545"

	txhash, err := Transaction.CreateTransaction(15, fromAddress, fromPrivateKeyHex, toAddress, transFerValue, gasLimit, gasPrice, network)
	if err != nil {
		t.Error(err)
		return
	}
	t.Logf("交易发出成功, txhash = %s", txhash) // 0x6c0042290ba16db5a23a630d4598133392f26c3c61f956297aea8bb86c7dbd2e
}

func TestInvalidTransactionService_CreateTransaction(t *testing.T) {
	fromAddress := common.HexToAddress("0xBAe52719E15Ab84564183b72a0299966727EC162") // index = 10
	fromPrivateKeyHex := "41794b550479390d8fff0e179655c33dada068b575751520575375e008ba624b"
	toAddress := common.HexToAddress("0x342a409070DBE2e2CDA157F77C581ece17B09795") // index = 11
	transFerValue := "110130000000000000000"                                       // 101.3 ETH (more than account has)
	gasPrice := "20000000000"                                                      // 0.00000002 ETH
	gasLimit := "21000"                                                            // gasLimit * gasPrice = 0.00042 ETH
	network := "http://localhost:7545"

	txhash, err := Transaction.CreateTransaction(15, fromAddress, fromPrivateKeyHex, toAddress, transFerValue, gasLimit, gasPrice, network)
	if err != nil {
		t.Error(err)
		return
	}
	t.Logf("交易发出成功, txhash = %s", txhash) // 0x6c0042290ba16db5a23a630d4598133392f26c3c61f956297aea8bb86c7dbd2e
}

func TestTransactionService_CheckTransaction(t *testing.T) {
	network := "http://localhost:7545"
	tx, err := Transaction.GetAndUpdateTransactionByHash("0xaeec0548216e738cdaa718b5c0e2a31ea89368ea50ba4909a4fe874adf9430d3", network)
	if err != nil {
		t.Error(err)
		return
	}
	t.Logf("%#v", tx)
}
