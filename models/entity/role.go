/*
 * @Author: fengzhilaoling fengzhilaoling@gmail.com
 * @Date: 2025-11-29 09:08:11
 * @LastEditors: fengzhilaoling
 * @LastEditTime: 2025-11-29 09:28:31
 * @FilePath: \ginManager\models\entity\role.go
 * @Description: 角色表
 * Copyright (c) 2025 by fengzhilaoling@gmail.com, All Rights Reserved.
 */
package entity

import (
	"time"

	"gorm.io/gorm"
)

// ----------------- 角色 -----------------
type Role struct {
	ID        uint64 `gorm:"primaryKey;autoIncrement;comment:角色主键"`
	RoleCode  string `gorm:"size:50;not null;uniqueIndex;comment:角色编码"`
	RoleName  string `gorm:"size:50;not null;comment:角色名称"`
	Sort      int    `gorm:"not null;default:0;comment:排序"`
	Status    uint8  `gorm:"not null;default:1;comment:1启用 0禁用"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"` //软删除
}
