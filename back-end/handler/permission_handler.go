/*
 * @Author: fengzhilaoling fengzhilaoling@gmail.com
 * @Date: 2025-12-01 09:25:38
 * @LastEditors: fengzhilaoling
 * @LastEditTime: 2025-12-01 15:43:22
 * @FilePath: \back-end\handler\permission_handler.go
 * @Description: 文件解释
 * Copyright (c) 2025 by fengzhilaoling@gmail.com, All Rights Reserved.
 */
package handler

import (
	"ginManager/dto"
	"ginManager/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type PermissionHandler struct {
	svc *service.PermissionService
}

func NewPermissionHandler() *PermissionHandler {
	return &PermissionHandler{svc: service.NewPermissionService()}
}

// List 全部权限
func (h *PermissionHandler) List(c *gin.Context) {
	list, err := h.svc.List(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusOK, dto.FailMsg("查询失败", err))
		return
	}
	c.JSON(http.StatusOK, dto.Success(list))
}
