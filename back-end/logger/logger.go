/*
 * @Author: fengzhilaoling fengzhilaoling@gmail.com
 * @Date: 2025-11-29 09:41:31
 * @LastEditors: fengzhilaoling
 * @LastEditTime: 2025-11-29 10:33:14
 * @FilePath: \ginManager\logger\logger.go
 * @Description: 全局日志实例
 * Copyright (c) 2025 by fengzhilaoling@gmail.com, All Rights Reserved.
 */
package logger

import (
	"ginManager/config"
	"path/filepath"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

var L *zap.Logger

// Init 初始化全局 Logger
func Init() {
	c := config.C.Log
	w := zapcore.AddSync(&lumberjack.Logger{
		Filename:   filepath.Join(c.Path, c.Filename),
		MaxSize:    c.MaxSize,
		MaxBackups: c.MaxBackups,
		MaxAge:     c.MaxAge,
		Compress:   c.Compress,
	})

	level := zapcore.InfoLevel
	_ = level.UnmarshalText([]byte(c.Level))

	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig()),
		w,
		level,
	)
	L = zap.New(core, zap.AddCaller())
}
