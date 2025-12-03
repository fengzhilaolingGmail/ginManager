/*
 * @Author: fengzhilaoling fengzhilaoling@gmail.com
 * @Date: 2025-11-30 09:47:58
 * @LastEditors: fengzhilaoling
 * @LastEditTime: 2025-12-01 19:10:01
 * @FilePath: \back-end\handler\menu_handler.go
 * @Description: 菜单处理器
 * Copyright (c) 2025 by fengzhilaoling@gmail.com, All Rights Reserved.
 */
// handler/menu_handler.go
package handler

import (
	"ginManager/dto"
	"ginManager/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type MenuHandler struct {
	svc     *service.MenuService
	permSvc *service.PermissionService // 权限
}

func NewMenuHandler() *MenuHandler {
	return &MenuHandler{
		svc:     service.NewMenuService(),
		permSvc: service.NewPermissionService(),
	}
}

// Tree 树形菜单（无权限也可看）
func (h *MenuHandler) Tree(c *gin.Context) {
	tree, err := h.svc.Tree(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusOK, dto.FailMsg("查询失败", err))
		return
	}
	c.JSON(http.StatusOK, dto.Success(tree))
}

// Create 新增
func (h *MenuHandler) Create(c *gin.Context) {
	var req dto.MenuAddReq
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
func (h *MenuHandler) Update(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	var req dto.MenuAddReq
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
func (h *MenuHandler) Delete(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	if err := h.svc.Delete(c.Request.Context(), id); err != nil {
		c.JSON(http.StatusOK, dto.FailMsg(err.Error(), err))
		return
	}
	c.JSON(http.StatusOK, dto.Success(nil))
}

// Get 单条
func (h *MenuHandler) Get(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	m, err := h.svc.GetByID(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusOK, dto.FailMsg("查询失败", err))
		return
	}
	c.JSON(http.StatusOK, dto.Success(m))
}

// handler/menu_handler.go
func (h *MenuHandler) Side(c *gin.Context) {
	resp, err := h.svc.SideMenu(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusOK, dto.FailMsg("查询失败", err))
		return
	}
	c.JSON(http.StatusOK, resp)
}

func (h *MenuHandler) TreeForRole(c *gin.Context) {
	roleID, _ := strconv.ParseUint(c.Query("roleId"), 10, 64)

	// 1. 所有菜单（权限）
	all, err := h.svc.Tree(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusOK, dto.FailMsg("查询菜单失败", err))
		return
	}

	// 2. 角色已有权限 ID

	checked, err := h.permSvc.GetPermIDsByRole(c.Request.Context(), roleID)
	if err != nil {
		c.JSON(http.StatusOK, dto.FailMsg("获取角色权限失败", err))
		return
	}
	checkedMap := make(map[uint64]bool)
	for _, id := range checked {
		checkedMap[id] = true
	}
	// 3. 转换并返回
	tree := service.BuildPermTree(all, checkedMap)
	c.JSON(http.StatusOK, dto.Success(tree))
}
