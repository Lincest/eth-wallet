package utils

import (
	"crypto/md5"
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
