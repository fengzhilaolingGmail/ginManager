/*
 * @Author: fengzhilaoling fengzhilaoling@gmail.com
 * @Date: 2025-11-29 10:37:56
 * @LastEditors: fengzhilaoling
 * @LastEditTime: 2025-12-01 11:28:56
 * @FilePath: \ginManager\router\router.go
 * @Description: 路由配置
 * Copyright (c) 2025 by fengzhilaoling@gmail.com, All Rights Reserved.
 */
package router

import (
	"ginManager/handler"
	"ginManager/middleware"
	"ginManager/router/api"
	"ginManager/service"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

// NewRouter 创建并配置好所有路由
func NewRouter(zapLog *zap.Logger, db *gorm.DB) *gin.Engine {
	r := gin.New()

	// 1. 全局中间件
	middleware.RegisterGlobalMW(r, zapLog)
	// 注册全局操作日志中间件
	logSvc := middleware.UserLogMiddleware(service.NewUserLogService())
	r.Use(logSvc)

	// 2. 业务路由分组
	root := r.Group("/api") // 统一加 /api 前缀
	handler.RegisterRoutes(root)
	{
		api.RegisterBaseRoutes(root) // /api/ping

		// 后续继续挂更多模块
		// api.RegisterUserRoutes(root)
		// api.RegisterMenuRoutes(root)
	}

	return r
}
