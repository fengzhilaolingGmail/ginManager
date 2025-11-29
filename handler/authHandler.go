/*
 * @Author: fengzhilaoling fengzhilaoling@gmail.com
 * @Date: 2025-11-29 13:42:42
 * @LastEditors: fengzhilaoling
 * @LastEditTime: 2025-11-29 17:25:45
 * @FilePath: \ginManager\handler\authHandler.go
 * @Description: 文件解释
 * Copyright (c) 2025 by fengzhilaoling@gmail.com, All Rights Reserved.
 */
package handler

import (
	"fmt"
	"ginManager/config"
	"ginManager/dto"
	"ginManager/logger"
	"ginManager/service"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type AuthHandler struct {
	svc *service.AuthService
}

func NewAuthHandler() *AuthHandler {
	return &AuthHandler{svc: service.NewAuthService()}
}

// Login 登录 -> Layui 格式
func (h *AuthHandler) Login(c *gin.Context) {
	var req dto.LoginReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusOK, dto.FailMsg("参数校验失败", err))
		return
	}
	fmt.Printf("req: %v\n", req)
	token, user, err := h.svc.Login(c.Request.Context(), &req)
	if err != nil {
		logger.L.Warn("login fail", zap.Error(err))
		c.JSON(http.StatusOK, dto.FailMsg("登录失败!", err))
		return
	}
	c.JSON(http.StatusOK, dto.Success(gin.H{
		"access_token": token,
		"expire":       config.C.JWT.Expire,
		"user": gin.H{
			"id":       user.ID,
			"username": user.Username,
			"nickname": user.Nickname,
		},
	}))
}
