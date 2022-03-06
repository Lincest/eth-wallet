package model

/**
    utils
    @author: roccoshi
    @desc: 基本的返回值
**/

// Result Code
const (
	CodeOK  = 0
	CodeErr = -1
)

// BasicResp 基本返回值
type BasicResp struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}
