/*
 * @Author: fengzhilaoling fengzhilaoling@gmail.com
 * @Date: 2025-11-30 09:47:31
 * @LastEditors: fengzhilaoling
 * @LastEditTime: 2025-11-30 09:51:54
 * @FilePath: \ginManager\handler\group_handler.go
 * @Description: 用户组处理器
 * Copyright (c) 2025 by fengzhilaoling@gmail.com, All Rights Reserved.
 */
// handler/group_handler.go
package handler

import (
	"ginManager/dto"
	"ginManager/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type GroupHandler struct {
	svc *service.GroupService
}

func NewGroupHandler() *GroupHandler {
	return &GroupHandler{svc: service.NewGroupService()}
}

// Page 分页
func (h *GroupHandler) Page(c *gin.Context) {
	var req dto.GroupListReq
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
func (h *GroupHandler) Create(c *gin.Context) {
	var req dto.GroupAddReq
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
func (h *GroupHandler) Update(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	var req dto.GroupAddReq
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
func (h *GroupHandler) Delete(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	if err := h.svc.Delete(c.Request.Context(), id); err != nil {
		c.JSON(http.StatusOK, dto.FailMsg(err.Error(), err))
		return
	}
	c.JSON(http.StatusOK, dto.Success(nil))
}

// UpdateStatus 切换组状态
func (h *GroupHandler) UpdateStatus(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	status, _ := strconv.ParseUint(c.Param("status"), 10, 8)
	if err := h.svc.UpdateStatus(c.Request.Context(), id, uint8(status)); err != nil {
		c.JSON(http.StatusOK, dto.FailMsg(err.Error(), err))
		return
	}
	c.JSON(http.StatusOK, dto.Success(nil))
}

// Get 单条
func (h *GroupHandler) Get(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	g, err := h.svc.GetByID(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusOK, dto.FailMsg("查询失败", err))
		return
	}
	c.JSON(http.StatusOK, dto.Success(g))
}
