/*
 * @Author: fengzhilaoling fengzhilaoling@gmail.com
 * @Date: 2025-11-29 10:40:27
 * @LastEditors: fengzhilaoling
 * @LastEditTime: 2025-11-29 10:41:41
 * @FilePath: \ginManager\router\api\base.go
 * @Description: 通用路由
 * Copyright (c) 2025 by fengzhilaoling@gmail.com, All Rights Reserved.
 */

package api

import "github.com/gin-gonic/gin"

// RegisterBaseRoutes 注册通用路由
func RegisterBaseRoutes(r *gin.RouterGroup) {
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"msg": "pong"})
	})
}
