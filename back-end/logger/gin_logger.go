/*
 * @Author: fengzhilaoling fengzhilaoling@gmail.com
 * @Date: 2025-11-29 10:32:07
 * @LastEditors: fengzhilaoling
 * @LastEditTime: 2025-11-29 10:32:14
 * @FilePath: \ginManager\logger\ginLogger.go
 * @Description: gin 日志中间件
 * Copyright (c) 2025 by fengzhilaoling@gmail.com, All Rights Reserved.
 */
package logger

import (
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// GinLogger 返回一个 Gin 中间件函数，用于记录请求日志
func GinLogger(l *zap.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path
		query := c.Request.URL.RawQuery

		// 继续处理请求
		c.Next()

		// 请求结束后记录
		l.Info("gin access",
			zap.Int("status", c.Writer.Status()),
			zap.String("method", c.Request.Method),
			zap.String("path", path),
			zap.String("query", query),
			zap.String("ip", c.ClientIP()),
			zap.String("user-agent", c.Request.UserAgent()),
			zap.Duration("latency", time.Since(start)),
			zap.String("time", start.Format("2006-01-02 15:04:05")),
		)
	}
}
