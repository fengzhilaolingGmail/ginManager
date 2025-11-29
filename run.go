/*
 * @Author: fengzhilaoling fengzhilaoling@gmail.com
 * @Date: 2025-11-29 10:28:07
 * @LastEditors: fengzhilaoling
 * @LastEditTime: 2025-11-29 10:48:46
 * @FilePath: \ginManager\run.go
 * @Description: 程序入口
 * Copyright (c) 2025 by fengzhilaoling@gmail.com, All Rights Reserved.
 */

package main

import (
	"fmt"
	"ginManager/config"
	"ginManager/logger"
	"ginManager/models"
	"ginManager/repository"
	"ginManager/router"

	"go.uber.org/zap"
)

func main() {
	// 1. 读配置
	config.Init("") // 默认读取 ./config.yaml
	// 2. 初始化日志
	logger.Init()
	// 3. 初始化数据库
	db := models.Init()
	// 注入 repo 层 DB 对象
	repository.SetDB(db)

	// 4. 路由
	r := router.NewRouter(logger.L, db)
	addr := fmt.Sprintf(":%d", config.C.Server.Port)
	logger.L.Info("server start", zap.String("addr", addr))
	if err := r.Run(addr); err != nil {
		logger.L.Fatal("server fail", zap.Error(err))
	}
}
