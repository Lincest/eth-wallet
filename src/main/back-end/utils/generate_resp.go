package utils

import "back-end/model"

/**
    utils
    @author: roccoshi
    @desc: 基本返回值生成
**/

// NewBasicResp 生成一个基本返回值
func NewBasicResp() *model.BasicResp {
	return &model.BasicResp{
		Code: model.CodeOK,
		Msg:  "success",
		Data: nil,
	}
}

// NewErrBasicResp 生成一个基本错误返回值
func NewErrBasicResp() *model.BasicResp {
	return &model.BasicResp{
		Code: model.CodeErr,
		Msg:  "error",
		Data: nil,
	}
}
