/*
 * @Author: fengzhilaoling fengzhilaoling@gmail.com
 * @Date: 2025-12-01 12:44:49
 * @LastEditors: fengzhilaoling
 * @LastEditTime: 2025-12-01 15:46:58
 * @FilePath: \back-end\handler\role_handler.go
 * @Description: 文件解释
 * Copyright (c) 2025 by fengzhilaoling@gmail.com, All Rights Reserved.
 */
package handler

import (
	"ginManager/dto"
	"ginManager/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type RoleHandler struct {
	svc *service.RoleService
}

func NewRoleHandler() *RoleHandler {
	return &RoleHandler{svc: service.NewRoleService()}
}

// Page 分页
func (h *RoleHandler) Page(c *gin.Context) {
	var req dto.RoleListReq
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusOK, dto.FailMsg("参数错误", err))
		return
	}

	list, total, err := h.svc.Page(c.Request.Context(), &req)
	if err != nil {
		c.JSON(http.StatusOK, dto.FailMsg("查询失败", err))
		return
	}
	c.JSON(http.StatusOK, dto.SuccessPage(list, total))
}

// Create 新增
func (h *RoleHandler) Create(c *gin.Context) {
	var req dto.RoleAddReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusOK, dto.FailMsg("参数错误", err))
		return
	}
	if err := h.svc.Create(c.Request.Context(), &req); err != nil {
		c.JSON(http.StatusOK, dto.FailMsg(err.Error(), err))
		return
	}
	c.JSON(http.StatusOK, dto.Success(nil))
}

// Update 编辑
func (h *RoleHandler) Update(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	var req dto.RoleAddReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusOK, dto.FailMsg("参数错误", err))
		return
	}
	if err := h.svc.Update(c.Request.Context(), id, &req); err != nil {
		c.JSON(http.StatusOK, dto.FailMsg(err.Error(), err))
		return
	}
	c.JSON(http.StatusOK, dto.Success(nil))
}

// Delete 删除
func (h *RoleHandler) Delete(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	if err := h.svc.Delete(c.Request.Context(), id); err != nil {
		c.JSON(http.StatusOK, dto.FailMsg(err.Error(), err))
		return
	}
	c.JSON(http.StatusOK, dto.Success(nil))
}

// Get 单条 + 已有权限 IDs
func (h *RoleHandler) Get(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	role, permIDs, err := h.svc.Get(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusOK, dto.FailMsg("查询失败", err))
		return
	}
	c.JSON(http.StatusOK, dto.Success(gin.H{
		"role":     role,
		"perm_ids": permIDs,
	}))
}
