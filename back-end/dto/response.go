/*
 * @Author: fengzhilaoling fengzhilaoling@gmail.com
 * @Date: 2025-11-29 13:48:27
 * @LastEditors: fengzhilaoling
 * @LastEditTime: 2025-11-29 17:46:54
 * @FilePath: \ginManager\dto\response.go
 * @Description: 文件解释
 * Copyright (c) 2025 by fengzhilaoling@gmail.com, All Rights Reserved.
 */
package dto

import (
	"ginManager/logger"

	"go.uber.org/zap"
)

// LayuiJSON 官方 table 要求：code=0 成功，其余失败
type LayuiJSON struct {
	Code  int         `json:"code"`
	Msg   string      `json:"msg"`
	Data  interface{} `json:"data,omitempty"`
	Count int64       `json:"count,omitempty"` // 分页总条数
}

// Success 非分页成功
func Success(data interface{}) LayuiJSON {
	return LayuiJSON{Code: 0, Msg: "success", Data: data}
}

// SuccessPage 分页成功
func SuccessPage(list interface{}, total int64) LayuiJSON {
	return LayuiJSON{Code: 0, Msg: "success", Data: list, Count: total}
}

// Fail 失败
func Fail(code int, msg string, err error) LayuiJSON {
	var errStr string
	if err != nil {
		errStr = err.Error()
	}
	logger.L.Error("Response Fail:", zap.String("err", errStr))
	return LayuiJSON{Code: code, Msg: msg}
}

// FailMsg 快捷失败（code=1）
func FailMsg(msg string, err error) LayuiJSON {
	return Fail(1, msg, err)
}
