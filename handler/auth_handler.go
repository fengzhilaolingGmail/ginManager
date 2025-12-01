/*
 * @Author: fengzhilaoling fengzhilaoling@gmail.com
 * @Date: 2025-11-29 13:42:42
 * @LastEditors: fengzhilaoling
 * @LastEditTime: 2025-12-01 11:46:29
 * @FilePath: \ginManager\handler\auth_handler.go
 * @Description: 文件解释
 * Copyright (c) 2025 by fengzhilaoling@gmail.com, All Rights Reserved.
 */
package handler

import (
	"ginManager/config"
	"ginManager/dto"
	"ginManager/logger"
	"ginManager/models/entity"
	"ginManager/service"
	"net/http"
	"time"

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
	start := time.Now()
	var req dto.LoginReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusOK, dto.FailMsg("参数校验失败", err))
		return
	}
	module := "auth"
	action := "login"
	method := c.Request.Method
	path := c.Request.URL.Path
	ip := c.ClientIP()
	userAgent := c.Request.UserAgent()
	var durationMs int
	log := entity.UserLog{
		Module:     &module,
		Action:     &action,
		Method:     &method,
		Path:       &path,
		IP:         &ip,
		UserAgent:  &userAgent,
		Status:     1,
		DurationMs: &durationMs,
		CreatedAt:  time.Now(),
	}
	log.Username = &req.Username
	token, user, err := h.svc.Login(c.Request.Context(), &req)
	if err != nil {
		logger.L.Warn("login fail", zap.Error(err))
		c.JSON(http.StatusOK, dto.FailMsg("登录失败!", err))
		log.Status = 0
		errStr := err.Error()
		log.ErrorMsg = &errStr
		durationMs = int(time.Since(start).Milliseconds())
		log.DurationMs = &durationMs
		_ = service.NewUserLogService().Create(c.Request.Context(), &log)
		return
	}
	log.UserID = &user.ID
	log.Username = &user.Username
	durationMs = int(time.Since(start).Milliseconds())
	log.DurationMs = &durationMs
	_ = service.NewUserLogService().Create(c.Request.Context(), &log)
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
