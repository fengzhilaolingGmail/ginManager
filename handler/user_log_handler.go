/*
 * @Author: fengzhilaoling fengzhilaoling@gmail.com
 * @Date: 2025-12-01 10:59:47
 * @LastEditors: fengzhilaoling
 * @LastEditTime: 2025-12-01 12:20:23
 * @FilePath: \ginManager\handler\user_log_handler.go
 * @Description: 文件解释
 * Copyright (c) 2025 by fengzhilaoling@gmail.com, All Rights Reserved.
 */
// handler/oper_log_handler.go
package handler

import (
	"encoding/csv"
	"ginManager/dto"
	"ginManager/service"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type UserLogHandler struct {
	svc *service.UserLogService
}

func NewUserLogHandler() *UserLogHandler {
	return &UserLogHandler{svc: service.NewUserLogService()}
}

// Page 分页列表
func (h *UserLogHandler) Page(c *gin.Context) {
	var req dto.UserLogListReq
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

// Export 导出 CSV（示例）
func (h *UserLogHandler) Export(c *gin.Context) {
	var req dto.UserLogExportReq
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusOK, dto.FailMsg("参数错误", err))
		return
	}
	// 默认导出最近 3 个月，最多 1 万条
	if req.EndTime.IsZero() {
		req.EndTime = time.Now()
	}
	if req.StartTime.IsZero() {
		req.StartTime = req.EndTime.AddDate(0, -3, 0)
	}
	list, _, err := h.svc.Page(c.Request.Context(), &dto.UserLogListReq{
		Module:    req.Module,
		Username:  req.Username,
		Status:    req.Status,
		StartTime: req.StartTime,
		EndTime:   req.EndTime,
		PageReq:   dto.PageReq{Page: 1, Limit: 10000},
	})
	if err != nil {
		c.JSON(http.StatusOK, dto.FailMsg("导出失败", err))
		return
	}

	// 写 CSV（简单示例）
	c.Header("Content-Type", "text/csv;charset=utf-8")
	c.Header("Content-Disposition", "attachment;filename=UserLog_"+time.Now().Format("20060102")+".csv")
	w := csv.NewWriter(c.Writer)
	w.Write([]string{"用户", "模块", "动作", "路径", "IP", "耗时(ms)", "状态", "时间"})
	for _, v := range list {
		statusStr := "成功"
		if v.Status == 0 {
			statusStr = "失败"
		}
		w.Write([]string{
			*v.Username,
			*v.Module,
			*v.Action,
			*v.Path,
			*v.IP,
			strconv.Itoa(*v.DurationMs),
			statusStr,
			v.CreatedAt.Format("2006-01-02 15:04:05"),
		})
	}
	w.Flush()
}
