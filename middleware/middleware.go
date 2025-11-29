/*
 * @Author: fengzhilaoling fengzhilaoling@gmail.com
 * @Date: 2025-11-29 10:38:35
 * @LastEditors: fengzhilaoling
 * @LastEditTime: 2025-11-29 10:41:52
 * @FilePath: \ginManager\middleware\middleware.go
 * @Description: 中间件注册
 * Copyright (c) 2025 by fengzhilaoling@gmail.com, All Rights Reserved.
 */
package middleware

import (
	"ginManager/logger"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// RegisterGlobalMW 注册全局中间件
func RegisterGlobalMW(r *gin.Engine, zapLog *zap.Logger) {
	r.Use(gin.Recovery())           //  panic 恢复
	r.Use(logger.GinLogger(zapLog)) //  请求日志
	// r.Use(CORS())                   //  后续再扩展
}
