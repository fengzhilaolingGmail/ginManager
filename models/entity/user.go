/*
 * @Author: fengzhilaoling fengzhilaoling@gmail.com
 * @Date: 2025-11-29 09:07:54
 * @LastEditors: fengzhilaoling
 * @LastEditTime: 2025-11-29 09:29:02
 * @FilePath: \ginManager\models\entity\user.go
 * @Description: 用户表
 * Copyright (c) 2025 by fengzhilaoling@gmail.com, All Rights Reserved.
 */

package entity

import "time"

// ----------------- 用户 -----------------
type User struct {
	ID           uint64  `gorm:"primaryKey;autoIncrement;comment:用户主键"`
	Username     string  `gorm:"size:50;not null;uniqueIndex;comment:登录账号"`
	PasswordHash string  `gorm:"size:60;not null;comment:bcrypt哈希"`
	Nickname     string  `gorm:"size:50;not null;comment:显示名"`
	Email        *string `gorm:"size:100;uniqueIndex;comment:邮箱"`
	Phone        *string `gorm:"size:20;uniqueIndex;comment:手机"`
	Status       uint8   `gorm:"not null;default:1;comment:1启用 0禁用"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
