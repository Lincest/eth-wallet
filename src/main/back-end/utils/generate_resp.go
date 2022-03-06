package utils

import . "back-end/model"

/**
    utils
    @author: roccoshi
    @desc: 基本返回值生成
**/

// NewBasicResp 生成一个基本返回值
func NewBasicResp() *BasicResp {
	return &BasicResp{
		Code: CodeOK,
		Msg:  "success",
		Data: nil,
	}
}

// NewErrBasicResp 生成一个基本错误返回值
func NewErrBasicResp() *BasicResp {
	return &BasicResp{
		Code: CodeErr,
		Msg:  "error",
		Data: nil,
	}
}
