# keystore管理	

## 1 - keystore导入

1. 用户可以手动导入一个keystore
2. 系统将根据keystore文件和密码生成对应的私钥和地址存入数据库中

```go
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
```

## 2 - keystore导出

1. 系统查询该用户的所有的privatekey
2. 从session中获取用户的密码作为密钥
3. 通过privatekey slice和密钥生成keystore文件, 放入临时文件夹中
4. 然后通过zip压缩将keystore文件打包
5. api返回`application/zip`格式的文件



