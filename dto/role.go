/*
 * @Author: fengzhilaoling fengzhilaoling@gmail.com
 * @Date: 2025-12-01 09:23:59
 * @LastEditors: fengzhilaoling
 * @LastEditTime: 2025-12-01 09:24:11
 * @FilePath: \ginManager\dto\role.go
 * @Description: 文件解释
 * Copyright (c) 2025 by fengzhilaoling@gmail.com, All Rights Reserved.
 */
package dto

type RoleListReq struct {
	PageReq
	RoleName string `form:"role_name"` // 模糊
}

type RoleAddReq struct {
	RoleCode string   `json:"role_code" binding:"required,max=50"`
	RoleName string   `json:"role_name" binding:"required,max=50"`
	Sort     int      `json:"sort"`
	Status   uint8    `json:"status" binding:"oneof=0 1"`
	PermIDs  []uint64 `json:"perm_ids"` // 权限 ID 数组
}
