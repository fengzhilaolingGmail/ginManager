/*
 * @Author: fengzhilaoling fengzhilaoling@gmail.com
 * @Date: 2025-12-01 12:44:49
 * @LastEditors: fengzhilaoling
 * @LastEditTime: 2025-12-07 16:31:02
 * @FilePath: \back-end\dto\group.go
 * @Description: 文件解释
 * Copyright (c) 2025 by fengzhilaoling@gmail.com, All Rights Reserved.
 */
// dto/group.go
package dto

type GroupListReq struct {
	PageReq
	GroupName string `form:"group_name"`        // 模糊
	Deleted   *uint8 `form:"deleted,omitempty"` // nil 不筛选 0 未删 1 已删
}

type GroupAddReq struct {
	GroupCode   string `json:"group_code" binding:"required,max=50"`
	GroupName   string `json:"group_name" binding:"required,max=50"`
	Sort        int    `json:"sort"`
	Status      uint8  `json:"status" binding:"oneof=0 1"`
	Description string `json:"description"`
}
