/*
 * @Author: fengzhilaoling fengzhilaoling@gmail.com
 * @Date: 2025-11-29 09:09:47
 * @LastEditors: fengzhilaoling
 * @LastEditTime: 2025-11-29 09:27:53
 * @FilePath: \ginManager\models\entity\menu.go
 * @Description: 菜单表
 * Copyright (c) 2025 by fengzhilaoling@gmail.com, All Rights Reserved.
 */

package entity

import "time"

// ----------------- 菜单 -----------------
type Menu struct {
	ID        uint64  `gorm:"primaryKey;autoIncrement;comment:菜单主键"`
	ParentID  uint64  `gorm:"not null;default:0;comment:父菜单ID，0=根"`
	Path      *string `gorm:"size:200;comment:前端路由"`
	Component *string `gorm:"size:100;comment:前端组件路径"`
	Name      *string `gorm:"size:50;comment:路由name"`
	Title     string  `gorm:"size:50;not null;comment:菜单标题"`
	Icon      *string `gorm:"size:50;comment:图标"`
	Sort      int     `gorm:"not null;default:0;comment:排序"`
	Status    uint8   `gorm:"not null;default:1;comment:1启用 0禁用"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
