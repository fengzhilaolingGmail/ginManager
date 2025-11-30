/*
 * @Author: fengzhilaoling fengzhilaoling@gmail.com
 * @Date: 2025-11-29 09:08:01
 * @LastEditors: fengzhilaoling
 * @LastEditTime: 2025-11-29 09:12:18
 * @FilePath: \ginManager\models\entity\userGroup.go
 * @Description: 用户组表
 * Copyright (c) 2025 by fengzhilaoling@gmail.com, All Rights Reserved.
 */
package entity

import (
	"time"

	"gorm.io/gorm"
)

// ----------------- 用户组 -----------------
type UserGroup struct {
	ID        uint64 `gorm:"primaryKey;autoIncrement;comment:组主键"`
	GroupCode string `gorm:"size:50;not null;uniqueIndex;comment:组编码"`
	GroupName string `gorm:"size:50;not null;comment:组名称"`
	Sort      int    `gorm:"not null;default:0;comment:排序"`
	Status    uint8  `gorm:"not null;default:1;comment:1启用 0禁用"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"` //软删除
}
