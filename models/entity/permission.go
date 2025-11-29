/*
 * @Author: fengzhilaoling fengzhilaoling@gmail.com
 * @Date: 2025-11-29 09:09:36
 * @LastEditors: fengzhilaoling
 * @LastEditTime: 2025-11-29 09:28:05
 * @FilePath: \ginManager\models\entity\permission.go
 * @Description: 权限表
 * Copyright (c) 2025 by fengzhilaoling@gmail.com, All Rights Reserved.
 */

package entity

import "time"

// ----------------- 权限 -----------------
type Permission struct {
	ID        uint64 `gorm:"primaryKey;autoIncrement;comment:权限主键"`
	PermCode  string `gorm:"size:100;not null;uniqueIndex;comment:权限编码如system:user:add"`
	PermName  string `gorm:"size:100;not null;comment:权限名称"`
	Sort      int    `gorm:"not null;default:0;comment:排序"`
	Status    uint8  `gorm:"not null;default:1;comment:1启用 0禁用"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
