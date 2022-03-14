package utils

import (
	"crypto/md5"
	"crypto/rc4"
	"fmt"
)

/**
    utils
    @author: roccoshi
    @desc: 加密 解密相关算法/工具
**/

// Md5encode Md5 加密 string -> string
func (*IEncrypt) Md5encode(x string) string {
	md5Sum := md5.Sum([]byte(x))
	return fmt.Sprintf("%x", md5Sum)
}

// Rc4encode rc4 加密 <=> 解密
// should I use this to encode mnemonic?
func (*IEncrypt) Rc4encode(origin string, key string) (string, error) {
	originByte := []byte(origin)
	keyByte := []byte(key)
	cipher, err := rc4.NewCipher(keyByte)
	if err != nil {
		return "", err
	}
	cipher.XORKeyStream(originByte, originByte)
	return string(originByte), nil
}
