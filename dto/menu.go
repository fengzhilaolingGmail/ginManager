/*
 * @Author: fengzhilaoling fengzhilaoling@gmail.com
 * @Date: 2025-11-30 09:42:30
 * @LastEditors: fengzhilaoling
 * @LastEditTime: 2025-11-30 09:43:28
 * @FilePath: \ginManager\dto\menu.go
 * @Description: 文件解释
 * Copyright (c) 2025 by fengzhilaoling@gmail.com, All Rights Reserved.
 */
// dto/menu.go
package dto

type MenuAddReq struct {
	ParentID  uint64  `json:"parent_id"`
	Title     string  `json:"title" binding:"required,max=50"`
	Name      *string `json:"name"`      // 路由 name
	Path      *string `json:"path"`      // 前端路由
	Component *string `json:"component"` // 前端组件
	Sort      int     `json:"sort"`
	Status    uint8   `json:"status" binding:"oneof=0 1"`
}
