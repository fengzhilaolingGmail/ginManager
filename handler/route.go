/*
 * @Author: fengzhilaoling fengzhilaoling@gmail.com
 * @Date: 2025-11-29 13:43:26
 * @LastEditors: fengzhilaoling
 * @LastEditTime: 2025-11-29 13:43:33
 * @FilePath: \ginManager\handler\route.go
 * @Description: 文件解释
 * Copyright (c) 2025 by fengzhilaoling@gmail.com, All Rights Reserved.
 */
package handler

import (
	"github.com/gin-gonic/gin"
)

// RegisterRoutes 注册所有业务路由
// RegisterRoutes 注册路由函数
// 参数 r 是一个 gin 路由组指针，用于注册相关路由
// RegisterRoutes 注册路由函数
// 参数:
//
//	r - gin路由组指针，用于注册路由
func RegisterRoutes(r *gin.RouterGroup) {
	// 注册登录路由，POST请求到/auth/login路径，调用auth.Login方法处理
	// 创建认证处理器实例
	auth := NewAuthHandler()
	r.POST("/auth/login", auth.Login)

	// 后续
	// user := NewUserHandler()
	// r.GET("/user/info", user.Info)
}
